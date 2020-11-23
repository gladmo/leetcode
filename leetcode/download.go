package leetcode

import (
	"html"
	"strings"

	"github.com/lunny/html2md"
)

type QuestionDetail struct {
	Data struct {
		Question struct {
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
			CompanyTagStats interface{} `json:"companyTagStats"`
			CodeSnippets    []struct {
				Lang     string `json:"lang"`
				LangSlug string `json:"langSlug"`
				Code     string `json:"code"`
				Typename string `json:"__typename"`
			} `json:"codeSnippets"`
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
		} `json:"question"`
	} `json:"data"`
}

func (th *QuestionDetail) Download(override bool) {
	questionMarkdown := html.UnescapeString(html2md.Convert(th.Data.Question.TranslatedContent))

	var tags []string
	for _, topicTag := range th.Data.Question.TopicTags {
		tags = append(tags, topicTag.TranslatedName)
	}

	for _, item := range th.Data.Question.CodeSnippets {
		localization := Localization{
			Language:       item.LangSlug,
			TitleSlug:      th.Data.Question.TitleSlug,
			QuestionID:     th.Data.Question.QuestionID,
			Title:          th.Data.Question.TranslatedTitle,
			Difficulty:     strings.ToLower(th.Data.Question.Difficulty),
			Question:       questionMarkdown,
			CodeSnippet:    item.Code,
			SampleTestCase: th.Data.Question.SampleTestCase,
			Tags:           tags,
		}

		localization.Save(override)
	}
}
