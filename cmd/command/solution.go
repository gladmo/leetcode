package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/spf13/cobra"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/store"
)

var solutionCmd = &cobra.Command{
	Use:       "solution",
	Short:     "题解管理",
	Long:      "你测试过的题解都可以被查看，还可以将之前提交的代码检出到问题列表中",
	ValidArgs: []string{"get", "describe", "code", "checkout", "list", "diff"},
}

func init() {
	solutionCmd.AddCommand(
		solutionListCmd,
		solutionGetCmd,
		solutionDescribeCmd,
		solutionCodeCmd,
		solutionCheckoutCmd,
		solutionDiffCmd,
		solutionTidyCmd,
		solutionEmptyCmd,
	)
}

const timeFormat = "2006-01-02"

var solutionListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有你测试过的题解",
	Run: func(cmd *cobra.Command, args []string) {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetRowLine(true)
		table.SetHeader([]string{"问题ID", "标题", "难度", "版本数", "测试次数", "首次提交", "最后提交"})

		list, err := store.ListSolution()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, solution := range list {
			info := leet.GetQuestionInfo(leet.Parse(solution.QuestionID))
			table.Append([]string{
				info.QuestionID,
				info.TranslatedTitle,
				info.Difficulty,
				fmt.Sprint(solution.Version),
				fmt.Sprint(solution.Times),
				solution.FirstTime.Format(timeFormat),
				solution.LastTime.Format(timeFormat),
			})
		}

		table.Render()
	},
}

var solutionGetCmd = &cobra.Command{
	Use:   "get question_id|leetcode_url",
	Short: "获取某一题目的题解列表",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		titleSlug := leet.Parse(strings.TrimSpace(args[0]))
		info := leet.GetQuestionInfo(titleSlug)

		solutions, err := store.GetSolution(info.QuestionID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetRowLine(true)
		// table.SetAutoMergeCells(true)
		table.SetHeader([]string{"序号", "语言", "测试次数", "评价", "消耗", "创建时间", "备注"})

		for idx, solution := range solutions {
			table.Append([]string{
				fmt.Sprint(idx + 1),
				solution.Language,
				fmt.Sprint(solution.Times),
				solution.Evaluation,
				fmt.Sprintf("%v", solution.Consumption-(solution.Consumption%time.Millisecond)),
				solution.CreatedAt.Format(timeFormat),
				solution.Remark,
			})
		}

		table.Render()
	},
}

var solutionCodeCmd = &cobra.Command{
	Use:   "code question_id|leetcode_url [solution_no]",
	Short: "显示你的题解代码",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 || len(args) > 2 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		// 展示最后一个
		var latest bool
		var idx int
		var err error
		if len(args) == 1 {
			latest = true
		} else {
			solutionIndex := strings.TrimSpace(args[1])
			idx, err = strconv.Atoi(solutionIndex)
			if err != nil {
				latest = true
			}
			// 转换成数组索引
			idx--
		}

		titleSlug := leet.Parse(strings.TrimSpace(args[0]))

		info := leet.GetQuestionInfo(titleSlug)

		solutions, err := store.GetSolution(info.QuestionID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if latest || idx > len(solutions) {
			idx = len(solutions) - 1
		}

		solution := solutions[idx]
		fmt.Println(leet.CustomerCode(solution.Code))
	},
}

var solutionCheckoutCmd = &cobra.Command{
	Use:   "checkout question_id|leetcode_url solution_no",
	Short: "将题解恢复到 questions 目录中，可以使用test重新测试",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		titleSlug := leet.Parse(strings.TrimSpace(args[0]))
		solutionIndex := strings.TrimSpace(args[1])
		idx, err := strconv.Atoi(solutionIndex)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		idx--

		info := leet.GetQuestionInfo(titleSlug)
		solutions, err := store.GetSolution(info.QuestionID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if idx < 0 || idx >= len(solutions) {
			idx++
			fmt.Println(fmt.Sprintf("最大题解数为: %d, 当前输入: %d", len(solutions), idx))
			return
		}
		solution := solutions[idx]

		err = ioutil.WriteFile(solution.SourceDir, []byte(solution.Code), 0755)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(fmt.Sprintf(
			"题目: %s, 编号: %s, 代码检出成功。",
			info.TranslatedTitle,
			info.QuestionID,
		))

		codeLine := fmt.Sprintf("%s:%d", solution.SourceDir, leet.CustomerCodeLine(solution.Code))
		fmt.Println(fmt.Sprintf("路径: %s", codeLine))
	},
}

var solutionDescribeCmd = &cobra.Command{
	Use:   "describe question_id|leetcode_url solution_no",
	Short: "详细描述你的题解内容",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		titleSlug := leet.Parse(strings.TrimSpace(args[0]))
		solutionIndex := strings.TrimSpace(args[1])
		idx, err := strconv.Atoi(solutionIndex)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// 转为数组索引
		idx--

		info := leet.GetQuestionInfo(titleSlug)
		solutions, err := store.GetSolution(info.QuestionID)
		if err != nil {
			fmt.Println(err)
			return
		}

		if idx < 0 || idx >= len(solutions) {
			idx++
			fmt.Println(fmt.Sprintf("最大题解数为: %d, 当前输入: %d", len(solutions), idx))
			return
		}

		solution := solutions[idx]
		code := leet.CustomerCode(solution.Code)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetRowLine(true)
		table.Append([]string{"编号:", info.QuestionID})
		table.Append([]string{"题目:", info.TranslatedTitle})
		table.Append([]string{"难度:", info.Difficulty})
		table.Append([]string{"语言:", solution.Language})
		codeLine := fmt.Sprintf("%s:%d", solution.SourceDir, leet.CustomerCodeLine(solution.Code))
		table.Append([]string{"路径:", codeLine})
		table.Append([]string{"时间:", solution.CreatedAt.Format(timeFormat)})
		table.Render()

		fmt.Println("代码:")
		fmt.Println(code)
	},
}

var solutionDiffCmd = &cobra.Command{
	Use:   "diff question_id|leetcode_url solution_no_1 solution_no_2",
	Short: "比较两次代码的不同之处",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		titleSlug := leet.Parse(strings.TrimSpace(args[0]))
		solutionIndex := strings.TrimSpace(args[1])
		idx, err := strconv.Atoi(solutionIndex)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		idx--

		solutionIndex2 := strings.TrimSpace(args[2])
		idx2, err := strconv.Atoi(solutionIndex2)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		idx2--

		info := leet.GetQuestionInfo(titleSlug)
		solutions, err := store.GetSolution(info.QuestionID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if idx < 0 || idx >= len(solutions) || idx2 < 0 || idx2 >= len(solutions) {
			idx++
			fmt.Println(fmt.Sprintf("最大题解数为: %d, 当前输入: %d %d", len(solutions), idx, idx2))
			return
		}

		solution1 := solutions[idx]
		solution2 := solutions[idx2]

		dmp := diffmatchpatch.New()

		diffs := dmp.DiffMain(
			leet.CustomerCode(solution1.Code),
			leet.CustomerCode(solution2.Code),
			true)

		for _, diff := range diffs {
			switch diff.Type {
			case diffmatchpatch.DiffDelete:
				fmt.Print("---")
			case diffmatchpatch.DiffInsert:
				fmt.Print("+++")
			}

			fmt.Print(diff.Text)
		}
	},
}

var solutionTidyCmd = &cobra.Command{
	Use:   "tidy question_id|leetcode_url solution_no",
	Short: "清空单个问题所有题解",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		titleSlug := leet.Parse(strings.TrimSpace(args[0]))
		info := leet.GetQuestionInfo(titleSlug)

		err := store.TidySolution(info.QuestionID)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("整理完成")
	},
}

var solutionEmptyCmd = &cobra.Command{
	Use:   "empty question_id|leetcode_url solution_no",
	Short: "清空单个问题所有题解",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		titleSlug := leet.Parse(strings.TrimSpace(args[0]))
		info := leet.GetQuestionInfo(titleSlug)

		err := store.RemoveSolution(info.QuestionID)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("清空完成")
	},
}
