package median_of_two_sorted_arrs_test

import (
	"algo/median_of_two_sorted_arrs"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1,2,3,5,6,8,9,11",
			args: args{
				nums1: []int{1, 3, 5, 8},
				nums2: []int{2, 6, 9, 11},
			},
			want: 5.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := median_of_two_sorted_arrs.FindMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("FindMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
