package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/gladmo/leetcode/leet"
)

var clearCmd = &cobra.Command{
	Use:   "clear [question_id|leetcode_url]",
	Short: "set questions to default (backup all unbanked)",
	Long:  "将本地问题恢复为初始状态",
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
		backup, err := cmd.PersistentFlags().GetBool("backup")
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
				if backup {
					err = leet.BackupClear{
						Dir: dir,
					}.Backup(override)
					if err != nil {
						fmt.Println(err.Error())
					}
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
				for _, dir := range info.SaveDir {

					if backup {
						err = leet.BackupClear{
							Dir: dir,
						}.Backup(override)
						if err != nil {
							fmt.Println(err.Error())
						}
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
	clearCmd.PersistentFlags().Bool("backup", false, "backup")
	clearCmd.PersistentFlags().Bool("override-backup", false, "override-backup")
	clearCmd.PersistentFlags().Bool("with-detail", false, "with-detail")
}
