package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/gladmo/leetcode/leet"
)

var getCmd = &cobra.Command{
	Use:   "get question_id|leetcode_url",
	Short: "get leetcode question from leetcode-cn.com",
	Example: `
leetcode get 795
leetcode get leetcode-cn.com/problems/k-th-symbol-in-grammar
leetcode get https://leetcode-cn.com/problems/k-th-symbol-in-grammar
leetcode get https://leetcode-cn.com/problems/k-th-symbol-in-grammar/solution/
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		override, err := cmd.PersistentFlags().GetBool("override")
		if err != nil {
			fmt.Println(err.Error())
		}
		withDetail, err := cmd.PersistentFlags().GetBool("with-detail")
		if err != nil {
			fmt.Println(err.Error())
		}

		param := leet.Parse(strings.TrimSpace(args[0]))
		if param == "" {
			cmd.Println(fmt.Sprintf("未找到关于 「%s」 的相关题目", args[0]))
			cmd.Help()
			os.Exit(1)
			return
		}

		res, err := leet.Fetch(param)
		if err != nil {
			panic(err)
		}

		res.Download(override)

		leet.InfoPrint(leet.GetQuestionInfo(param), withDetail)
	},
}

func init() {
	getCmd.PersistentFlags().Bool("override", false, "override")
	getCmd.PersistentFlags().Bool("with-detail", false, "with-detail")
}
