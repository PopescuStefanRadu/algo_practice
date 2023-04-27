package sort_colors

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			reflect.DeepEqual(tt.args.nums, tt.expected)
		})
	}
}
