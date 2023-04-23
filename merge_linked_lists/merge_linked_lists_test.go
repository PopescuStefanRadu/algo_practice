package merge_linked_lists_test

import (
	"algo/merge_linked_lists"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list2 := &merge_linked_lists.ListNode{Val: 1}

	type args struct {
		list1 *merge_linked_lists.ListNode
		list2 *merge_linked_lists.ListNode
	}
	tests := []struct {
		name string
		args args
		want *merge_linked_lists.ListNode
	}{
		{
			name: "two empty lists",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{
			name: "one empty list",
			args: args{
				list1: nil,
				list2: list2,
			},
			want: list2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge_linked_lists.MergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
