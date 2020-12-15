package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/gladmo/leetcode/leet"
)

var infoCmd = &cobra.Command{
	Use:   "info question_id|leetcode_url",
	Short: "print leetcode question info",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		param := leet.Parse(strings.TrimSpace(args[0]))

		withDetail, err := cmd.PersistentFlags().GetBool("with-detail")
		if err != nil {
			fmt.Println(err.Error())
		}

		leet.InfoPrint(leet.GetQuestionInfo(param), withDetail)
	},
}

func init() {
	infoCmd.PersistentFlags().Bool("with-detail", false, "with-detail")
}
