package unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	type args struct {
		m      int
		n      int
		memory map[string]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "28",
			args: args{
				m:      3,
				n:      7,
				memory: map[string]int{},
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePaths(tt.args.m, tt.args.n, tt.args.memory); got != tt.want {
				t.Errorf("UniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
