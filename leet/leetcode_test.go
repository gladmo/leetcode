package leet

import (
	"testing"
)

func TestParseFromURL(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "solution1",
			args: args{
				param: "https://leetcode-cn.com/problems/k-th-symbol-in-grammar/solution/",
			},
			want: "k-th-symbol-in-grammar",
		},
		{
			name: "solution2",
			args: args{
				param: "leetcode-cn.com/problems/k-th-symbol-in-grammar/solution/",
			},
			want: "k-th-symbol-in-grammar",
		},
		{
			name: "solution3",
			args: args{
				param: "https://leetcode-cn.com/problems/k-th-symbol-in-grammar/",
			},
			want: "k-th-symbol-in-grammar",
		},
		{
			name: "solution4",
			args: args{
				param: "https://leetcode-cn.com/problems/k-th-symbol-in-grammar",
			},
			want: "k-th-symbol-in-grammar",
		},
		{
			name: "solution5",
			args: args{
				param: "leetcode-cn.com/problems/k-th-symbol-in-grammar",
			},
			want: "k-th-symbol-in-grammar",
		},
		{
			name: "solution6",
			args: args{
				param: "leetcode-cn.com/problems/k-th-symbol-in-grammar/",
			},
			want: "k-th-symbol-in-grammar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseFromURL(tt.args.param); got != tt.want {
				t.Errorf("ParseFromURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
