package store

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

// Question 问题描述
type Question struct {
	QuestionID string    `json:"question_id"`
	Title      string    `json:"title"`
	Remark     string    `json:"remark"`
	CreatedAt  time.Time `json:"created_at"`
}

// NewMarkQuestion 问题书签
func NewMarkQuestion(questionID, title, remark string) Question {
	return Question{
		QuestionID: questionID,
		Title:      title,
		Remark:     remark,
		CreatedAt:  time.Now(),
	}
}

type Questions []Question

func (th Questions) Len() int {
	return len(th)
}

func (th Questions) Less(i, j int) bool {
	ii, _ := strconv.Atoi(th[i].QuestionID)
	jj, _ := strconv.Atoi(th[j].QuestionID)
	return ii < jj
}

func (th Questions) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}

func (th *Questions) Add(q Question) {
	for _, question := range *th {
		if question.QuestionID == q.QuestionID {
			return
		}
	}

	*th = append(*th, q)
}

func (th Questions) Bytes() []byte {
	b, _ := json.Marshal(th)
	return b
}

func MarkQuestion(question Question) error {
	private, err := privateDB()
	if err != nil {
		return err
	}
	defer private.Close()

	return private.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("mark.questions"))
		if err != nil {
			return err
		}

		key := []byte(fmt.Sprint(question.QuestionID))
		ss := b.Get(key)

		var qs Questions
		if len(ss) != 0 {
			err = json.Unmarshal(ss, &qs)
			if err != nil {
				return err
			}
		}

		qs.Add(question)

		return b.Put(key, qs.Bytes())
	})
}

func ListMarkQuestions() (qs Questions, err error) {
	private, err := privateDB()
	if err != nil {
		return
	}
	defer private.Close()

	err = private.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("mark.questions"))

		return b.ForEach(func(k, v []byte) error {
			var q Questions
			err = json.Unmarshal(v, &q)
			if err != nil {
				return err
			}

			qs = append(qs, q...)
			return nil
		})
	})

	sort.Sort(qs)
	return
}

// markSolution 题解书签
type markSolution struct {
	Solution `json:"solution"`

	CreatedAt time.Time `json:"created_at"`
	Remark    string    `json:"remark"`
}

func NewMarkSolution(remark string, s Solution) markSolution {
	return markSolution{
		Solution:  s,
		CreatedAt: time.Now(),
		Remark:    remark,
	}
}

type MarkSolutions []markSolution

func (th MarkSolutions) Len() int {
	return len(th)
}

func (th MarkSolutions) Less(i, j int) bool {
	ii, _ := strconv.Atoi(th[i].QuestionID)
	jj, _ := strconv.Atoi(th[j].QuestionID)
	return ii < jj
}

func (th MarkSolutions) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}

func (th *MarkSolutions) Add(ms markSolution) {
	for _, solution := range *th {
		if ms.CodeHash == ms.CodeHash &&
			ms.Remark == solution.Remark {
			return
		}
	}

	*th = append(*th, ms)
}

func (th MarkSolutions) Bytes() []byte {
	b, _ := json.Marshal(th)
	return b
}

func MarkSolution(solution markSolution) error {
	private, err := privateDB()
	if err != nil {
		return err
	}
	defer private.Close()

	return private.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("mark.solutions"))
		if err != nil {
			return err
		}

		key := []byte(fmt.Sprint(solution.QuestionID))
		ss := b.Get(key)

		var ms MarkSolutions
		if len(ss) != 0 {
			err = json.Unmarshal(ss, &ms)
			if err != nil {
				return err
			}
		}

		ms.Add(solution)

		return b.Put(key, ms.Bytes())
	})
}

func ListMarkSolutions() (ss MarkSolutions, err error) {
	private, err := privateDB()
	if err != nil {
		return
	}
	defer private.Close()

	err = private.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("mark.solutions"))

		if b == nil {
			return nil
		}

		return b.ForEach(func(k, v []byte) error {
			var s MarkSolutions
			err = json.Unmarshal(v, &s)
			if err != nil {
				return err
			}

			ss = append(ss, s...)
			return nil
		})
	})

	sort.Sort(ss)
	return
}
