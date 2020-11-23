package leetcode

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Localization struct {
	Language       string   `json:"language"`
	TitleSlug      string   `json:"title_slug"`
	QuestionID     string   `json:"question_id"`
	Title          string   `json:"title"`
	Difficulty     string   `json:"difficulty"`
	Question       string   `json:"question"`
	CodeSnippet    string   `json:"code_snippet"`
	SampleTestCase string   `json:"sample_test_case"`
	Tags           []string `json:"tags"`
}

func (th *Localization) Save(override bool) {
	baseDir := "questions"
	so := SaveOption{
		saveDir:        baseDir,
		title:          th.Title,
		titleSlug:      th.TitleSlug,
		difficulty:     th.Difficulty,
		question:       th.Question,
		codeSnippet:    th.CodeSnippet,
		sampleTestCase: th.SampleTestCase,
		language:       th.Language,
	}

	switch th.Difficulty {
	case "easy":
		th.Difficulty = "简单"
	case "hard":
		th.Difficulty = "困难"
	case "medium":
		th.Difficulty = "中等"
	}

	so.saveDir = path.Join(baseDir, "serial", th.Difficulty, th.QuestionID)

	// 已存在不重新获取，除非强制
	if f, err := os.Stat(path.Join(so.saveDir, so.language)); err == nil && f.IsDir() && !override {
		fmt.Println("--------", so.language, "-------")
		fmt.Println("是否覆盖: ", override)
		fmt.Println("语言: ", so.language)
		fmt.Println("题目已经存在: ", so.saveDir)
		return
	}

	err := so.SaveQuestion()
	if err != nil {
		fmt.Println(err)
	}

	for _, tag := range th.Tags {
		so.saveDir = path.Join(baseDir, "tags", tag, th.Difficulty, th.TitleSlug)
		err := so.SaveQuestion()
		if err != nil {
			fmt.Println(err)
		}
	}
}

type SaveOption struct {
	saveDir, title, titleSlug,
	difficulty, question, codeSnippet,
	sampleTestCase, language string
}

func (th SaveOption) SaveQuestion() (err error) {
	err = os.MkdirAll(th.saveDir, 0755)
	if err != nil {
		return
	}

	switch th.language {
	case "golang":
		err = th.golang()

		// todo other language
	}
	if err != nil {
		return
	}

	th.question = fmt.Sprintf(`## [%s](%s)

%s`,
		th.title,
		fmt.Sprintf("https://leetcode-cn.com/problems/%s/", th.titleSlug),
		th.question)

	err = ioutil.WriteFile(
		path.Join(th.saveDir, "question.md"),
		[]byte(th.question),
		0755)
	if err != nil {
		return
	}

	return
}
