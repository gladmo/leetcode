package leet

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"

	"github.com/gladmo/leetcode/store"
)

func InfoPrint(info store.Store, withDetail bool) {
	if info.QuestionID == "" {
		fmt.Println("问题还未下载到本地，请先使用 get 命令下载")
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	table.Append([]string{"标题", info.Title})
	table.Append([]string{"问题ID", info.QuestionID})
	table.Append([]string{"中文标题", info.TranslatedTitle})
	// table.Append([]string{"语言列表", strings.Join(info.Languages, ",")})
	table.Append([]string{"标签", strings.Join(info.Tags, ",")})
	table.Append([]string{"难度", info.Difficulty})
	codeLine := fmt.Sprintf("%s/README.md:2", info.SaveDir[0])
	table.Append([]string{"题目描述", codeLine})
	goFileDir := fmt.Sprintf(
		"%s/golang/solution/%s.go",
		info.SaveDir[0],
		info.TitleSlug)
	goCodeLine := fmt.Sprintf(
		"%s:%d",
		goFileDir,
		CustomerCodeLineByFile(goFileDir),
	)
	table.Append([]string{"Go代码", goCodeLine})

	table.Render()

	if withDetail {
		fmt.Println(info.Question)
	}
}

func GetQuestionInfoByID(questionID string) (info store.Store) {
	var err error
	info, err = store.QuestionInfo(questionID)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func GetQuestionInfo(titleSlug string) (info store.Store) {
	var err error
	info, err = store.QuestionInfo(titleSlug)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func GetAllQuestionTitleSlug() (res []string) {
	var err error
	res, err = store.AllQuestionTitleSlug()
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func UpdateQuestionInfo(que Question, local Localization) {
	var tags []string
	for _, topicTag := range que.TopicTags {
		tags = append(tags, topicTag.TranslatedName)
	}

	var langs []string
	for _, lan := range que.CodeSnippets {
		langs = append(langs, lan.LangSlug)
	}

	info := store.Store{
		Title:           que.Title,
		TranslatedTitle: que.TranslatedTitle,
		QuestionID:      local.QuestionID,
		Languages:       langs,
		Tags:            tags,
		Difficulty:      local.Difficulty,
		SaveDir:         local.SaveDirs,
		TitleSlug:       local.TitleSlug,
		Question:        local.Question,
	}

	err := store.UpdateQuestionInfo(info)
	if err != nil {
		fmt.Println(err.Error())
	}
}
