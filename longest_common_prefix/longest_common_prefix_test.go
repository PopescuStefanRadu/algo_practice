package longest_common_prefix_test

import (
	"algo/longest_common_prefix"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no common prefix",
			args: args{
				strs: []string{"a", "b", "cad"},
			},
			want: "",
		},
		{
			name: "one string",
			args: args{
				strs: []string{"cad"},
			},
			want: "cad",
		},
		{
			name: "one common prefix",
			args: args{
				strs: []string{"a", "ab", "acad"},
			},
			want: "a",
		},
		{
			name: "multiple prefixes",
			args: args{
				strs: []string{"aba", "abad", "abade"},
			},
			want: "aba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longest_common_prefix.LongestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
