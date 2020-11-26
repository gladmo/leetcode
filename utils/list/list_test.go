package list

import (
	"testing"
)

func TestNode_String(t *testing.T) {
	type fields struct {
		Val  int
		Next *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "[1,2,3,4]",
			fields: fields{
				Val: 1,
				Next: &Node{
					Val: 2,
					Next: &Node{
						Val: 3,
						Next: &Node{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: "[1,2,3,4]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := &Node{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := th.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
