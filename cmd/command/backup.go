package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/gladmo/leetcode/leet"
)

var backupCmd = &cobra.Command{
	Use:   "backup [question_id|leetcode_url]",
	Short: "[Deprecated]backup you complete questions to solutions",
	Long:  "将你实现的算法备份到 solutions 目录，贡献代码者此命令很有用",
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
			info.Print(withDetail)

			for _, dir := range info.SaveDir {
				err = leet.BackupClear{
					Dir: dir,
				}.Backup(override)
				if err != nil {
					fmt.Println(err.Error())
				}
			}

		} else {
			titles := leet.GetAllQuestionTitleSlug()

			for _, title := range titles {
				info := leet.GetQuestionInfo(title)
				for _, dir := range info.SaveDir {
					err = leet.BackupClear{
						Dir: dir,
					}.Backup(override)
					if err != nil {
						fmt.Println(err.Error())
					}
				}
			}
		}

	},
}

func init() {
	backupCmd.PersistentFlags().Bool("override-backup", false, "override-backup")
	backupCmd.PersistentFlags().Bool("with-detail", false, "with-detail")
}
