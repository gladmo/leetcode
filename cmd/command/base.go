package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/gladmo/leetcode/leet"
)

var baseCmd = &cobra.Command{
	Use:   "base [question_id|leetcode_url]",
	Short: "clear & replace all question use you specified",
	Long:  "恢复所有问题为默认状态，并将 serial 目录的所有问题覆盖到 tag 目录下",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		if len(args) == 1 {
			param := leet.Parse(strings.TrimSpace(args[0]))

			withDetail, err := cmd.PersistentFlags().GetBool("with-detail")
			if err != nil {
				fmt.Println(err.Error())
			}

			info := leet.GetQuestionInfo(param)
			info.Print(withDetail)

			for idx, dir := range info.SaveDir {
				for _, language := range info.Languages {
					err = leet.ToBeClear{
						Dir:      dir,
						Language: language,
					}.Clear()
					if err != nil {
						fmt.Println(err.Error())
					}
				}

				if idx > 0 {
					err = leet.CopyDirectory(info.SaveDir[0], info.SaveDir[idx])
					if err != nil {
						fmt.Println(err.Error())
					}
				}
			}

		} else {
			titles := leet.GetAllQuestionTitleSlug()

			for _, title := range titles {
				info := leet.GetQuestionInfo(title)
				for idx, dir := range info.SaveDir {
					for _, language := range info.Languages {
						err := leet.ToBeClear{
							Dir:      dir,
							Language: language,
						}.Clear()
						if err != nil {
							fmt.Println(err.Error())
						}
					}

					if idx > 0 {
						err := leet.CopyDirectory(info.SaveDir[0], info.SaveDir[idx])
						if err != nil {
							fmt.Println(err.Error())
						}
					}
				}
			}
		}
	},
}

func init() {
	baseCmd.PersistentFlags().Bool("with-detail", false, "with-detail")
}