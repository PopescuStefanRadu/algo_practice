package first_missing_positive_test

import (
	"algo/first_missing_positive"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "	",
			args: args{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := first_missing_positive.FirstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("FirstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
