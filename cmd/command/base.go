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
	Short: "clear & replace all question use you specified (backup all unbanked)",
	Long:  "恢复所有问题为默认状态，并将 serial 目录的所有问题覆盖到 tag 目录下",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		override, err := cmd.PersistentFlags().GetBool("override-backup")
		if err != nil {
			fmt.Println(err.Error())
		}

		if len(args) == 1 {
			param := leet.Parse(strings.TrimSpace(args[0]))

			withDetail, err := cmd.PersistentFlags().GetBool("with-detail")
			if err != nil {
				fmt.Println(err.Error())
			}

			info := leet.GetQuestionInfo(param)
			leet.InfoPrint(info, withDetail)

			for idx, dir := range info.SaveDir {

				if idx > 0 {
					err = leet.CopyDirectory(info.SaveDir[0], info.SaveDir[idx])
					if err != nil {
						fmt.Println(err.Error())
					}
				}

				err = leet.BackupClear{
					Dir: dir,
				}.Backup(override)
				if err != nil {
					fmt.Println(err.Error())
				}

				for _, language := range info.Languages {
					err = leet.BackupClear{
						Dir:      dir,
						Language: language,
					}.Clear()
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

					if idx > 0 {
						err := leet.CopyDirectory(info.SaveDir[0], info.SaveDir[idx])
						if err != nil {
							fmt.Println(err.Error())
						}
					}

					err = leet.BackupClear{
						Dir: dir,
					}.Backup(override)
					if err != nil {
						fmt.Println(err.Error())
					}

					for _, language := range info.Languages {
						err := leet.BackupClear{
							Dir:      dir,
							Language: language,
						}.Clear()
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
	baseCmd.PersistentFlags().Bool("override-backup", false, "override-backup")
	baseCmd.PersistentFlags().Bool("with-detail", false, "with-detail")
}
