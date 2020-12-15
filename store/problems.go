package store

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

const allProblemsKey = `leetcode.problems`

func UpdateProblems(info []QuestionStats) error {
	return private.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(allProblemsKey))
		if err != nil {
			return err
		}

		for _, stats := range info {
			key := []byte(fmt.Sprint(stats.QuestionID))
			err = b.Put(key, stats.Bytes())
		}

		return err
	})
}

func ProblemsTTL(expireAt time.Time) error {
	return private.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(allProblemsKey))
		if err != nil {
			return err
		}

		key := []byte(fmt.Sprint("TTL"))
		return b.Put(key, []byte(expireAt.Format(time.RFC3339)))
	})
}

func ProblemInfoIsExpire() (ok bool, err error) {
	err = private.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(allProblemsKey))
		if err != nil {
			return err
		}

		key := []byte(fmt.Sprint("TTL"))
		data := b.Get(key)
		if len(data) == 0 {
			ok = false
			return nil
		}
		t, err := time.Parse(time.RFC3339, string(data))
		if err != nil {
			return err
		}
		ok = time.Now().After(t)
		return nil
	})

	return
}

func GetProblemsInfo(questionID string) (info QuestionStats, err error) {
	err = private.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(allProblemsKey))

		if b == nil {
			err = fmt.Errorf(`problem id: %s,  not found`, questionID)
			return err
		}

		data := b.Get([]byte(questionID))
		if len(data) == 0 {
			err = fmt.Errorf(`problem id: %s,  not found`, questionID)
			return err
		}
		err = json.Unmarshal(data, &info)
		return err
	})
	return
}

type QuestionStats struct {
	QuestionID          int    `json:"question_id"`
	QuestionTitle       string `json:"question__title"`
	QuestionTitleSlug   string `json:"question__title_slug"`
	QuestionHide        bool   `json:"question__hide"`
	TotalAcs            int    `json:"total_acs"`
	TotalSubmitted      int    `json:"total_submitted"`
	TotalColumnArticles int    `json:"total_column_articles"`
	FrontendQuestionID  string `json:"frontend_question_id"`
	IsNewQuestion       bool   `json:"is_new_question"`
}

func (th *QuestionStats) Bytes() []byte {
	b, _ := json.Marshal(th)
	return b
}
