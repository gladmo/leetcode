package command

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/store"
)

func init() {
	markQuestionCmd.PersistentFlags().String("remark", "", "[必填]为书签添加备注")
	markSolutionCmd.PersistentFlags().String("remark", "", "[必填]为书签添加备注")

	markCmd.AddCommand(markSolutionCmd, markQuestionCmd, markListCmd)
}

var markCmd = &cobra.Command{
	Use:       "mark solution|question question_id [solution_id]",
	Short:     "问题或者题解的书签",
	Long:      "将问题或者题解记录下来，时常翻阅时常回顾",
	ValidArgs: []string{"solution", "question", "list"},
}

var markSolutionCmd = &cobra.Command{
	Use:   "solution question_id solution_id",
	Short: "记录题解",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		remark, err := cmd.PersistentFlags().GetString("remark")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if len(remark) == 0 {
			fmt.Println("请填使用 --remark 填写备注")
			return
		}

		solutionIndex := strings.TrimSpace(args[1])
		idx, err := strconv.Atoi(solutionIndex)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		idx--

		titleSlug := leet.Parse(strings.TrimSpace(args[0]))
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

		err = store.MarkSolution(
			store.NewMarkSolution(remark, solution),
		)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(fmt.Sprintf("书签添加成功：%s", info.TranslatedTitle))
	},
}

var markQuestionCmd = &cobra.Command{
	Use:   "question question_id",
	Short: "记录题目",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		remark, err := cmd.PersistentFlags().GetString("remark")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if len(remark) == 0 {
			fmt.Println("请填使用 --remark 填写备注")
			return
		}

		titleSlug := leet.Parse(strings.TrimSpace(args[0]))
		info := leet.GetQuestionInfo(titleSlug)

		err = store.MarkQuestion(
			store.NewMarkQuestion(info.QuestionID, info.TranslatedTitle, remark),
		)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(fmt.Sprintf("书签添加成功：%s", info.TranslatedTitle))
	},
}

var markListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出书签",
	Run: func(cmd *cobra.Command, args []string) {
		qs, err := store.ListMarkQuestions()
		if err != nil {
			fmt.Println(err.Error())
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetRowLine(true)
		// table.SetAutoMergeCells(true)
		table.SetHeader([]string{"序号", "题号", "标题", "备注", "创建时间"})

		fmt.Println("问题书签：")
		for idx, question := range qs {
			table.Append([]string{
				fmt.Sprint(idx + 1),
				question.QuestionID,
				question.Title,
				question.Remark,
				question.CreatedAt.Format(timeFormat),
			})
		}

		table.Render()

		ss, err := store.ListMarkSolutions()
		if err != nil {
			fmt.Println(err.Error())
		}

		table2 := tablewriter.NewWriter(os.Stdout)
		table2.SetRowLine(true)
		fmt.Println("题解书签：")
		table2.SetHeader([]string{"序号", "题号", "标题", "代码Hash(前8位)", "备注", "创建时间"})
		for idx, solution := range ss {
			info := leet.GetQuestionInfo(leet.Parse(solution.QuestionID))
			table2.Append([]string{
				fmt.Sprint(idx + 1),
				solution.QuestionID,
				info.TranslatedTitle,
				solution.CodeHash[:8],
				solution.Remark,
				solution.CreatedAt.Format(timeFormat),
			})
		}

		table2.Render()
	},
}
