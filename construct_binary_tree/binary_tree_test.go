package construct_binary_tree

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func toPtr[T interface{}](a T) *T {
	return &a
}

func toStrForPtrSlice(elems any) string {
	elsTyp := reflect.TypeOf(elems)
	if elsTyp.Kind() != reflect.Slice {
		return "received input is not a slice"
	}

	res := strings.Builder{}
	res.WriteString("[]")
	elTyp := elsTyp.Elem()
	res.WriteString(elTyp.String())
	res.WriteString("{")

	s := reflect.ValueOf(elems)
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		if v.IsNil() {
			res.WriteString("null")
		} else {
			res.WriteString(fmt.Sprintf("%v", v.Elem().Interface()))
		}

		if i+1 < s.Len() {
			res.WriteString(",")
		}
	}

	res.WriteString("}")
	return res.String()
}

func TestToStrForPtrSlice(t *testing.T) {
	tests := []struct {
		name  string
		slice any
		want  string
	}{
		{
			name:  "not a slice",
			slice: 5,
			want:  "received input is not a slice",
		},
		{
			name:  "pointer slice",
			slice: []*int{toPtr(1), nil},
			want:  "[]*int{1,null}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toStrForPtrSlice(tt.slice); got != tt.want {
				t.Errorf("toStrForPtrSlice() = %s, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want []*int
	}{
		{
			name: "",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: []*int{
				toPtr(3), toPtr(9), nil, nil, toPtr(15), toPtr(7),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildTree(tt.args.preorder, tt.args.inorder)
			if got == nil && tt.want != nil {
				t.Errorf("BuildTree() = [], want %v", toStrForPtrSlice(tt.want))
			}
			printedByDepth := got.PrintByDepth()
			if !reflect.DeepEqual(printedByDepth, tt.want) {
				t.Errorf("BuildTree() = %v, want %v", toStrForPtrSlice(printedByDepth), toStrForPtrSlice(tt.want))
			}
		})
	}
}
