package leetcode

import (
	"testing"
)

func Test_parseCode(t *testing.T) {

	code := `
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

}
`

	type args struct {
		code string
	}
	tests := []struct {
		name        string
		args        args
		wantNewCode string
		wantOk      bool
	}{
		{
			name:        "test",
			args:        args{code: code},
			wantNewCode: "",
			wantOk:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewCode, gotOk := parseCode(tt.args.code)
			if gotNewCode != tt.wantNewCode {
				t.Errorf("parseCode() gotNewCode = %v, want %v", gotNewCode, tt.wantNewCode)
			}
			if gotOk != tt.wantOk {
				t.Errorf("parseCode() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
