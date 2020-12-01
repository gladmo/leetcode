package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/745/golang/solution"
)

func main() {
	/*

			["c","f","j"]
		"a"

	*/

	tests := []struct {
		name   string
		input1 []byte
		input2 byte
		want   byte
	}{
		{
			name:   "test-cfj-a",
			input1: []byte("cfj"),
			input2: 'a',
			want:   'c',
		},
		{
			name:   "test-cfj-c",
			input1: []byte("cfj"),
			input2: 'c',
			want:   'f',
		},
		{
			name:   "test-cfj-d",
			input1: []byte("cfj"),
			input2: 'd',
			want:   'f',
		},
		{
			name:   "test-cfj-g",
			input1: []byte("cfj"),
			input2: 'g',
			want:   'j',
		},
		{
			name:   "test-cfj-j",
			input1: []byte("cfj"),
			input2: 'j',
			want:   'c',
		},
		{
			name:   "test-cfj-k",
			input1: []byte("cfj"),
			input2: 'k',
			want:   'c',
		},
		{
			name:   "test-eeeeeennnn-e",
			input1: []byte("eeeeeennnn"),
			input2: 'e',
			want:   'n',
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
