package plus_one_test

import (
	"algo/plus_one"
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "11 + 1",
			args: args{
				digits: []int{1, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "999 + 1",
			args: args{
				digits: []int{9, 9, 9},
			},
			want: []int{1, 0, 0, 0},
		},
		{
			name: "189 + 1",
			args: args{
				digits: []int{1, 8, 9},
			},
			want: []int{1, 9, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plus_one.PlusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
