package store

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

const storeKey = `leetcode.questions`

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

// Stats leetcode db stats
func Stats() {
	leetcode, err := leetcodeDB()
	if err != nil {
		return
	}
	defer leetcode.Close()

	b, _ := json.Marshal(leetcode.Stats())
	fmt.Println(string(b))
}

func (th Store) Bytes() []byte {
	b, _ := json.Marshal(th)
	return b
}

func UpdateQuestionInfo(store Store) error {
	leetcode, err := leetcodeDB()
	if err != nil {
		return err
	}
	defer leetcode.Close()

	return leetcode.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(storeKey))
		if err != nil {
			return err
		}

		keyTitle := []byte(fmt.Sprint(store.TitleSlug))
		keyID := []byte(fmt.Sprint(store.QuestionID))
		b.Put(keyID, store.Bytes())
		return b.Put(keyTitle, store.Bytes())
	})
}

func QuestionInfo(titleSlug string) (info Store, err error) {
	leetcode, err := leetcodeDB()
	if err != nil {
		return
	}
	defer leetcode.Close()

	err = leetcode.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(storeKey))

		if b == nil {
			return nil
		}

		data := b.Get([]byte(titleSlug))
		if len(data) == 0 {
			return nil
		}

		err = json.Unmarshal(data, &info)
		return err
	})

	return
}

func AllQuestionTitleSlug() (titles []string, err error) {
	leetcode, err := leetcodeDB()
	if err != nil {
		return
	}
	defer leetcode.Close()

	err = leetcode.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(storeKey))

		if b == nil {
			return nil
		}

		return b.ForEach(func(k, v []byte) error {
			var s Store
			err = json.Unmarshal(v, &s)
			if err != nil {
				return err
			}

			titles = append(titles, s.TitleSlug)
			return nil
		})
	})

	return
}
