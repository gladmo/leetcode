package leet

import (
	"fmt"
	"html"
	"path"
	"strings"

	"github.com/lunny/html2md"
)

type QuestionDetail struct {
	Data struct {
		Question Question `json:"question"`
	} `json:"data"`
}

type Question struct {
	QuestionID            string        `json:"questionId"`
	QuestionFrontendID    string        `json:"questionFrontendId"`
	BoundTopicID          int           `json:"boundTopicId"`
	Title                 string        `json:"title"`
	TitleSlug             string        `json:"titleSlug"`
	Content               string        `json:"content"`
	TranslatedTitle       string        `json:"translatedTitle"`
	TranslatedContent     string        `json:"translatedContent"`
	IsPaidOnly            bool          `json:"isPaidOnly"`
	Difficulty            string        `json:"difficulty"`
	Likes                 int           `json:"likes"`
	Dislikes              int           `json:"dislikes"`
	IsLiked               interface{}   `json:"isLiked"`
	SimilarQuestions      string        `json:"similarQuestions"`
	Contributors          []interface{} `json:"contributors"`
	LangToValidPlayground string        `json:"langToValidPlayground"`
	TopicTags             []struct {
		Name           string `json:"name"`
		Slug           string `json:"slug"`
		TranslatedName string `json:"translatedName"`
		Typename       string `json:"__typename"`
	} `json:"topicTags"`
	CompanyTagStats   interface{}   `json:"companyTagStats"`
	CodeSnippets      []CodeSnippet `json:"codeSnippets"`
	Stats             string        `json:"stats"`
	Hints             []interface{} `json:"hints"`
	Solution          interface{}   `json:"solution"`
	Status            interface{}   `json:"status"`
	SampleTestCase    string        `json:"sampleTestCase"`
	MetaData          string        `json:"metaData"`
	JudgerAvailable   bool          `json:"judgerAvailable"`
	JudgeType         string        `json:"judgeType"`
	MysqlSchemas      []interface{} `json:"mysqlSchemas"`
	EnableRunCode     bool          `json:"enableRunCode"`
	EnvInfo           string        `json:"envInfo"`
	Book              interface{}   `json:"book"`
	IsSubscribed      bool          `json:"isSubscribed"`
	IsDailyQuestion   bool          `json:"isDailyQuestion"`
	DailyRecordStatus interface{}   `json:"dailyRecordStatus"`
	EditorType        string        `json:"editorType"`
	UgcQuestionID     interface{}   `json:"ugcQuestionId"`
	Style             string        `json:"style"`
	Typename          string        `json:"__typename"`
}

type CodeSnippet struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
	Typename string `json:"__typename"`
}

func (th *QuestionDetail) Download(override bool) {
	questionMarkdown := html.UnescapeString(html2md.Convert(th.Data.Question.TranslatedContent))
	question := fmt.Sprintf(`## [%s](%s)

%s`,
		th.Data.Question.TranslatedTitle,
		fmt.Sprintf("https://leetcode-cn.com/problems/%s/", th.Data.Question.TitleSlug),
		questionMarkdown)

	var tags []string
	for _, topicTag := range th.Data.Question.TopicTags {
		tags = append(tags, topicTag.TranslatedName)
	}

	difficulty := strings.ToLower(th.Data.Question.Difficulty)
	// var difficulty string
	// switch strings.ToLower(th.Data.Question.Difficulty) {
	// case "easy":
	// 	difficulty = "easy"
	// case "hard":
	// 	difficulty = "hard"
	// case "medium":
	// 	difficulty = "medium"
	// }

	localization := Localization{
		TitleSlug:      th.Data.Question.TitleSlug,
		QuestionID:     th.Data.Question.QuestionID,
		Title:          th.Data.Question.TranslatedTitle,
		Difficulty:     difficulty,
		Question:       question,
		CodeSnippets:   th.Data.Question.CodeSnippets,
		SampleTestCase: th.Data.Question.SampleTestCase,
	}

	baseDir := "questions"

	var saveDirs []string
	// 按题号
	saveDirs = append(saveDirs, path.Join(baseDir, "serial", difficulty, localization.QuestionID))

	// 按标签
	for _, tag := range tags {
		saveDirs = append(saveDirs, path.Join(baseDir, "tags", tag, difficulty, localization.TitleSlug))
	}

	// 待保存问题的目录列表
	localization.SaveDirs = saveDirs

	localization.Save(override)

	UpdateQuestionInfo(th.Data.Question, localization)
}
