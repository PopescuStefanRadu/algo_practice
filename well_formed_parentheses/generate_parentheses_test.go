package well_formed_parentheses_test

import (
	"algo/well_formed_parentheses"
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "one",
			args: args{
				n: 1,
			},
			want: []string{"()"},
		},
		{
			name: "three",
			args: args{
				n: 3,
			},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := well_formed_parentheses.GenerateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
