package command

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/store"
)

var testCmd = &cobra.Command{
	Use:   "test question_id|leetcode_url",
	Short: "测试本地代码，保存题解",
	Long:  "测试本地代码，并保存在你的题解列表中，可通过 solution 子命令查看",
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
		remark, _ := cmd.PersistentFlags().GetString("remark")

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

		var code []byte
		var commandName string
		var commandParams []string
		var sourceDir string
		switch language {
		case "golang":
			// go run questions/serial/medium/133/golang/main.go
			commandName = "go"
			fileDir := fmt.Sprintf(`%s/golang/main.go`, codeDir)
			commandParams = append(commandParams, "run", fileDir)

			sourceDir = fmt.Sprintf(`%s/golang/solution/%s.go`, codeDir, info.TitleSlug)
			b, _ := ioutil.ReadFile(sourceDir)
			code, _ = format.Source(b)
		}

		evaluation := ""
		t := time.Now()
		c := exec.Command(commandName, commandParams...)
		result, err := c.CombinedOutput()
		if err != nil {
			fmt.Println(err.Error())
			evaluation = "Compilation failed"
		} else {
			success := bytes.Count(result, []byte("PASS"))
			fail := bytes.Count(result, []byte("FAILED"))
			evaluation = fmt.Sprintf("%d/%d", success, success+fail)
		}

		solution := store.NewSolution(
			info.QuestionID,
			language,
			sourceDir,
			string(code),
			string(result),
			evaluation,
			remark,
			time.Since(t),
		)
		err = store.AddSolution(solution)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(result))
	},
}

func init() {
	testCmd.PersistentFlags().Bool("serial", true, "serial")
	testCmd.PersistentFlags().String("lang", "golang", "lang")
	testCmd.PersistentFlags().String("tag", "", "tag")
	testCmd.PersistentFlags().String("remark", "", "为测试添加备注")
}
