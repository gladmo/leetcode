package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/gladmo/leetcode/leet"
)

var testCmd = &cobra.Command{
	Use:   "test question_id|leetcode_url",
	Short: "test you code and analyse",
	Long:  "测试本地代码",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			cmd.Println("参数异常")
			cmd.Help()
			os.Exit(1)
			return
		}

		param := leet.Parse(strings.TrimSpace(args[0]))

		serial, err := cmd.PersistentFlags().GetBool("serial")
		if err != nil {
			fmt.Println(err.Error())
		}

		tag, err := cmd.PersistentFlags().GetString("tag")
		if err != nil {
			fmt.Println(err.Error())
		}
		if tag != "" {
			serial = false
		}

		language, err := cmd.PersistentFlags().GetString("lang")
		if err != nil {
			language = "golang"
		}

		var codeDir string
		info := leet.GetQuestionInfo(param)
		for _, s := range info.SaveDir {
			if serial && strings.Contains(s, "serial") {
				codeDir = s
				break
			}

			if tag != "" && strings.Contains(s, tag) {
				codeDir = s
				break
			}
		}

		if codeDir == "" {
			fmt.Println("question not found")
			os.Exit(2)
		}

		var commandName string
		var commandParams []string
		switch language {
		case "golang":
			// go run questions/serial/中等/133/golang/main.go
			commandName = "go"
			commandParams = append(commandParams, "run", fmt.Sprintf(`%s/golang/main.go`, codeDir))
		}

		// fmt.Println(commandName, strings.Join(commandParams, " "))

		c := exec.Command(commandName, commandParams...)
		result, err := c.CombinedOutput()
		if err != nil {
			fmt.Println(string(result))
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(result))
	},
}

func init() {
	testCmd.PersistentFlags().Bool("serial", true, "serial")
	testCmd.PersistentFlags().String("lang", "golang", "lang")
	testCmd.PersistentFlags().String("tag", "", "tag")
}
