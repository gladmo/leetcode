package leet

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Localization struct {
	TitleSlug      string        `json:"title_slug"`
	QuestionID     string        `json:"question_id"`
	Title          string        `json:"title"`
	Difficulty     string        `json:"difficulty"`
	Question       string        `json:"question"`
	CodeSnippets   []CodeSnippet `json:"code_snippets"`
	SampleTestCase string        `json:"sample_test_case"`

	SaveDirs []string `json:"save_dirs"`
}

func (th *Localization) Save(override bool) {
	for _, dir := range th.SaveDirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			continue
		}

		err = ioutil.WriteFile(
			path.Join(dir, "question.md"),
			[]byte(th.Question),
			0755)
		if err != nil {
			return
		}

		for _, snippet := range th.CodeSnippets {
			saveDir := path.Join(dir, snippet.LangSlug)

			// 已存在不重新获取，除非强制
			if f, err := os.Stat(saveDir); err == nil && f.IsDir() && !override {
				fmt.Println("--------", snippet.LangSlug, "-------")
				fmt.Println("是否覆盖: ", override)
				fmt.Println("语言: ", snippet.LangSlug)
				fmt.Println("题目已经存在: ", dir)
				continue
			}

			so := SaveOption{
				SaveDir:        saveDir,
				Title:          th.Title,
				TitleSlug:      th.TitleSlug,
				Difficulty:     th.Difficulty,
				Question:       th.Question,
				CodeSnippet:    snippet.Code,
				SampleTestCase: th.SampleTestCase,
				Language:       snippet.LangSlug,
			}

			err := so.SaveQuestion()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

type SaveOption struct {
	SaveDir, Title, TitleSlug,
	Difficulty, Question, CodeSnippet,
	SampleTestCase, Language string
}

func (th SaveOption) SaveQuestion() (err error) {
	switch th.Language {
	case "golang":
		return th.golang()

		// todo other language
	}

	return
}
