package roman_to_int_test

import (
	"algo/roman_to_int"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "XII",
			args: args{s: "XII"},
			want: 12,
		},
		{
			name: "IIX",
			args: args{s: "IIX"},
			want: 8,
		},
		{
			name: "MCMXCIV",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roman_to_int.RomanToInt(tt.args.s); got != tt.want {
				t.Errorf("RomanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
