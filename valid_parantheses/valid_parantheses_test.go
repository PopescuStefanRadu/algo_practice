package valid_parantheses_test

import (
	"algo/valid_parantheses"
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple parenthesis",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "unopened parenthesis",
			args: args{
				s: ")",
			},
			want: false,
		},
		{
			name: "unclosed parenthesis",
			args: args{
				s: "(",
			},
			want: false,
		},
		{
			name: "separate parnetheses",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "complex valid scenario",
			args: args{
				s: "([{()[()]}])()",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid_parantheses.IsValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("IsValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
