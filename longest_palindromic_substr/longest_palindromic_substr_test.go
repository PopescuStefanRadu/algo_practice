package longest_palindromic_substr_test

import (
	"algo/longest_palindromic_substr"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "(asd{ffdsa)asdff}",
			args: args{"asdffdsaasdff"},
			want: "ffdsaasdff",
		},
		{
			name: "{ad(da}d)",
			args: args{"addad"},
			want: "adda",
		},
		{
			name: "ada",
			args: args{"ada"},
			want: "ada",
		},
		{
			name: "dd",
			args: args{"dd"},
			want: "dd",
		},
		{
			name: "d",
			args: args{"d"},
			want: "d",
		},
		{
			name: "empty str",
			args: args{""},
			want: "",
		},
		{
			name: "cbbd",
			args: args{"cbbd"},
			want: "bb",
		},
		{
			name: "ac",
			args: args{"ac"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longest_palindromic_substr.LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
