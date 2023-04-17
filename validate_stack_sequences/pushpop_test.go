package validate_stack_sequences_test

import (
	"algo/validate_stack_sequences"
	"fmt"
	"testing"
)

func TestValidateStackSequences(t *testing.T) {
	tests := []struct {
		name           string
		expectedResult bool
		pushLis        []int
		popLis         []int
	}{
		{
			name:           "Test1",
			expectedResult: true,
			pushLis:        []int{1, 2, 3, 4, 5},
			popLis:         []int{4, 5, 3, 2, 1},
		},
		{
			name:           "Test2",
			expectedResult: false,
			pushLis:        []int{1, 2, 3, 4, 5},
			popLis:         []int{4, 3, 5, 1, 2},
		},
		{
			name:           "Test3",
			expectedResult: true,
			pushLis:        []int{1, 2, 3},
			popLis:         []int{1, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.name), func(t *testing.T) {
			isValid := validate_stack_sequences.ValidateStackSequences(tt.pushLis, tt.popLis)
			if isValid != tt.expectedResult {
				t.Fatalf("expected: %v, got: %v", tt.expectedResult, isValid)
			}
		})
	}

}
