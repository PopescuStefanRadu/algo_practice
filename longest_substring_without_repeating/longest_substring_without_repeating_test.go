package longest_substring_without_repeating_test

import (
	"algo/longest_substring_without_repeating"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "aba",
			args: args{
				s: "aba",
			},
			want: 2,
		},
		{
			name: "aBARDa",
			args: args{
				s: "abarda",
			},
			want: 4,
		},
		{
			name: "abaRDASTF",
			args: args{
				s: "abardastf",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longest_substring_without_repeating.LengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
