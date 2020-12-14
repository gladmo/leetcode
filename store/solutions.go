package store

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

// Solution 解题内容
type Solution struct {
	QuestionID  string        `json:"question_id"` // 问题id
	Language    string        `json:"language"`    // 语言
	CodeHash    string        `json:"code_hash"`   // 实现代码格式化后的hash
	SourceDir   string        `json:"source_dir"`  // 代码存放目录
	Code        string        `json:"code"`        // 实现的代码
	Result      string        `json:"result"`      // 测试返回结果
	Times       int           `json:"times"`       // 执行次数
	Consumption time.Duration `json:"consumption"` // 消耗
	Evaluation  string        `json:"evaluation"`  // 运行评价
	Remark      string        `json:"remark"`      // 备注
	CreatedAt   time.Time     `json:"create_at"`   // 创建内容时间
}

// NewSolution 创建题解
func NewSolution(questionID, lang, sourceDir, code, result, evaluation, remark string, consumption time.Duration) Solution {
	codeHash := md5.Sum([]byte(code))

	return Solution{
		QuestionID:  questionID,
		Language:    lang,
		CodeHash:    hex.EncodeToString(codeHash[:]),
		SourceDir:   sourceDir,
		Code:        code,
		Result:      result,
		Consumption: consumption,
		Evaluation:  strings.TrimSpace(evaluation),
		Remark:      remark,
		CreatedAt:   time.Now(),
		Times:       1,
	}
}

// 解题集
type answer []Solution

func (th *answer) Exists(solution Solution) bool {
	for idx, s := range *th {
		if s.CodeHash == solution.CodeHash {
			(*th)[idx].Times++
			(*th)[idx].Consumption = solution.Consumption
			(*th)[idx].Evaluation = solution.Evaluation
			(*th)[idx].Remark = solution.Remark
			(*th)[idx].Result = solution.Result
			return true
		}
	}

	return false
}

func (th *answer) Append(solution Solution) {
	*th = append(*th, solution)
}

func (th answer) Times() (times int) {
	for _, solution := range th {
		times += solution.Times
	}
	return
}

func (th answer) FirstTime() time.Time {
	return th[0].CreatedAt
}

func (th answer) LastTime() time.Time {
	return th[len(th)-1].CreatedAt
}

func (th answer) Bytes() []byte {
	b, _ := json.Marshal(th)
	return b
}

func (th *answer) Tidy() {
	var newAnswer answer
	for _, solution := range *th {
		tmp := strings.Split(solution.Evaluation, "/")
		if len(tmp) == 2 &&
			tmp[0] == tmp[1] &&
			strings.TrimSpace(tmp[0]) != "0" {
			newAnswer = append(newAnswer, solution)
		}
	}
	*th = newAnswer
}

// AddSolution 添加解题答案
func AddSolution(solution Solution) error {
	return private.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("solutions"))
		if err != nil {
			return err
		}

		key := []byte(fmt.Sprint(solution.QuestionID))
		ss := b.Get(key)

		var ans answer
		if len(ss) != 0 {
			err = json.Unmarshal(ss, &ans)
			if err != nil {
				return err
			}
		}

		if !ans.Exists(solution) {
			ans.Append(solution)
		}

		return b.Put(key, ans.Bytes())
	})
}

type SolutionsList struct {
	QuestionID string    `json:"question_id"`
	Version    int       `json:"version"`
	Times      int       `json:"times"` // 执行次数
	FirstTime  time.Time `json:"first_time"`
	LastTime   time.Time `json:"last_time"`
}

func ListSolution() (sl []SolutionsList, err error) {
	err = private.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("solutions"))

		var solutionsList SolutionsList
		return b.ForEach(func(k, v []byte) error {
			solutionsList.QuestionID = string(k)

			var ans answer
			err = json.Unmarshal(v, &ans)
			if err != nil {
				return err
			}
			solutionsList.Times = ans.Times()
			solutionsList.FirstTime = ans.FirstTime()
			solutionsList.LastTime = ans.LastTime()
			solutionsList.Version = len(ans)

			sl = append(sl, solutionsList)
			return nil
		})
	})

	return
}

// GetSolution 获取测试过的题解
func GetSolution(questionID string) (ss []Solution, err error) {
	err = private.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("solutions"))

		res := b.Get([]byte(questionID))
		if len(res) == 0 {
			return fmt.Errorf("question id: %s solutions not found", questionID)
		}

		return json.Unmarshal(res, &ss)
	})

	return
}

// RemoveSolution delete solution
func RemoveSolution(questionID string) error {
	return private.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("solutions"))

		return b.Delete([]byte(questionID))
	})
}

func TidySolution(questionID string) error {
	return private.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("solutions"))

		key := []byte(questionID)
		data := b.Get(key)
		if len(data) == 0 {
			return nil
		}

		var ans answer
		err := json.Unmarshal(data, &ans)
		if err != nil {
			return err
		}

		ans.Tidy()

		return b.Put(key, ans.Bytes())
	})
}
