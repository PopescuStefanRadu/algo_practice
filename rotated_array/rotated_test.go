package rotated_array_test

import (
	"algo/rotated_array"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "even",
			args: args{
				nums:   []int{3, 4, 5, 7},
				target: 5,
			},
			want: 2,
		},
		{
			name: "middle",
			args: args{
				nums:   []int{3, 4, 5, 7, 9},
				target: 5,
			},
			want: 2,
		},
		{
			name: "odd",
			args: args{
				nums:   []int{3, 4, 5, 7, 9},
				target: 7,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := rotated_array.BinarySearch(tt.args.nums, tt.args.target, 0, len(tt.args.nums)-1); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindRotationPosition(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	type response struct {
		pos     int
		rotated bool
	}
	tests := []struct {
		name string
		args args
		resp response
	}{
		{
			name: "first element",
			args: args{
				nums:  []int{0, 1, 2, 3, 4},
				left:  0,
				right: 4,
			},
			resp: response{
				pos:     0,
				rotated: false,
			},
		},
		{
			name: "second element",
			args: args{
				nums:  []int{4, 0, 1, 2, 3},
				left:  0,
				right: 4,
			},
			resp: response{
				pos:     1,
				rotated: true,
			},
		},
		{
			name: "last element",
			args: args{
				nums:  []int{1, 2, 3, 4, 0},
				left:  0,
				right: 4,
			},
			resp: response{
				pos:     4,
				rotated: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotationPosition, rotated := rotated_array.FindRotationPosition(tt.args.nums, tt.args.left, tt.args.right)
			if rotationPosition != tt.resp.pos {
				t.Errorf("FindRotationPosition() rotationPosition = %v, want %v", rotationPosition, tt.resp.pos)
			}
			if rotated != tt.resp.rotated {
				t.Errorf("FindRotationPosition() rotated = %v, want %v", rotated, tt.resp.rotated)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "not rotated, first elem",
			args: args{
				nums:   []int{0, 1, 2},
				target: 0,
			},
			want: 0,
		},
		{
			name: "last elem rotated, first elem",
			args: args{
				nums:   []int{1, 2, 0},
				target: 1,
			},
			want: 0,
		},
		{
			name: "last elem rotated, last elem",
			args: args{
				nums:   []int{1, 2, 0},
				target: 0,
			},
			want: 2,
		},
		{
			name: "first elem rotated, last elem",
			args: args{
				nums:   []int{2, 0, 1},
				target: 0,
			},
			want: 1,
		},
		{
			name: "not found",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "2 elems rotated, not found",
			args: args{
				nums:   []int{1, 3},
				target: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotated_array.Search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
