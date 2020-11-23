package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gladmo/leetcode/leetcode"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("例: leetcode get 22")
		os.Exit(1)
		return
	}

	if args[1] != "get" {
		fmt.Println("例: leetcode get 22")
		os.Exit(2)
		return
	}

	param := strings.TrimSpace(args[2])

	if match, err := regexp.MatchString(`\d+`, param); err == nil && match {
		stat, err := leetcode.ProblemID2name(param)
		if err != nil {
			panic(err)
		}

		param = stat.QuestionTitleSlug

		fmt.Println("问题 ID: ", stat.QuestionID)
		fmt.Println("总提交数: ", stat.TotalSubmitted)
		fmt.Println("题目解文章数: ", stat.TotalColumnArticles)
		fmt.Println("是否是新问题: ", stat.IsNewQuestion)
	}

	if strings.HasPrefix(param, "http") ||
		strings.HasPrefix(param, "leetcode-cn.com") {

		re := regexp.MustCompile(`leetcode-cn\.com/problems/(.*)`)

		result := re.FindStringSubmatch(param)
		if len(result) == 2 {
			param = strings.Trim(result[1], "/")
		}
	}

	fmt.Println("问题名称: ", param)

	res, err := leetcode.Fetch(param)
	if err != nil {
		panic(err)
	}

	res.Download(false)
}
