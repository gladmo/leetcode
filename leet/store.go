package leet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/olekukonko/tablewriter"
)

var store = make(map[string]Store)
var storeFile = path.Join("questions", "store.json")

func init() {
	updateCache()
}

func updateCache() (err error) {
	b, err := ioutil.ReadFile(storeFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &store)
	if err != nil {
		return
	}

	return
}

type Store struct {
	Title           string   `json:"title"`
	TranslatedTitle string   `json:"translated_title"`
	QuestionID      string   `json:"question_id"`
	Languages       []string `json:"language"`
	Tags            []string `json:"tags"`

	Difficulty string   `json:"difficulty"`
	SaveDir    []string `json:"save_dir"`
	TitleSlug  string   `json:"title_slug"`
	Question   string   `json:"question"`
}

func (th Store) Print(withDetail bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	table.Append([]string{"标题", th.Title})
	table.Append([]string{"问题ID", th.QuestionID})
	table.Append([]string{"中文标题", th.TranslatedTitle})
	// table.Append([]string{"语言列表", strings.Join(th.Languages, ",")})
	table.Append([]string{"标签", strings.Join(th.Tags, ",")})
	table.Append([]string{"难度", th.Difficulty})
	codeLine := fmt.Sprintf("%s/README.md:%d:1", th.SaveDir[0], 1)
	table.Append([]string{"题目描述", codeLine})

	table.Render()

	if withDetail {
		fmt.Println(th.Question)
	}
}

func GetQuestionInfo(titleSlug string) (info Store) {
	s, ok := store[titleSlug]
	if !ok {
		return
	}

	return s
}

func GetAllQuestionTitleSlug() (res []string) {
	for titleSlug := range store {
		res = append(res, titleSlug)
	}

	return
}

func UpdateQuestionInfo(que Question, local Localization) {
	err := updateCache()
	if err != nil {
		fmt.Println(err.Error())
	}

	var tags []string
	for _, topicTag := range que.TopicTags {
		tags = append(tags, topicTag.TranslatedName)
	}

	var langs []string
	for _, lan := range que.CodeSnippets {
		langs = append(langs, lan.LangSlug)
	}

	info := Store{
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

	store[info.TitleSlug] = info
	b, err := json.Marshal(store)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = ioutil.WriteFile(storeFile, b, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
}
