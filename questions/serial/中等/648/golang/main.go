package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/648/golang/solution"
)

func main() {
	/*

			["cat","bat","rat"]
		"the cattle was rattled by the battery"

	*/

	tests := []struct {
		name   string
		input1 []string
		input2 string
		want   string
	}{
		{
			name:   "test-[cat,bat,rat]",
			input1: []string{"cat", "bat", "rat"},
			input2: "the cattle was rattled by the battery",
			want:   "the cat was rat by the bat",
		},
		{
			name:   "test-[a,b,c]",
			input1: []string{"a", "b", "c"},
			input2: "aadsfasf absbs bbab cadsfafs",
			want:   "a a b c",
		},
		{
			name:   "test-[a,aa,aaa,aaaa]",
			input1: []string{"a", "aa", "aaa", "aaaa"},
			input2: "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa",
			want:   "a a a a a a a a bbb baba a",
		},
		{
			name:   "test-[catt,cat,bat,rat]",
			input1: []string{"catt", "cat", "bat", "rat"},
			input2: "the cattle was rattled by the battery",
			want:   "the cat was rat by the bat",
		},
		{
			name:   "test-[ac,ab]",
			input1: []string{"ac", "ab"},
			input2: "it is abnormal that this solution is accepted",
			want:   "it is ab that this solution is ac",
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input1, test.input2)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input1, test.input2)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
