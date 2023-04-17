package validate_stack_sequences

/**
946. Validate Stack Sequences

Given two integer arrays pushed and popped each with distinct values, return true if this could have been the result
of a sequence of push and pop operations on an initially empty stack, or false otherwise.

Example 1:

Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
Output: true
Explanation: We might do the following sequence:
push(1), push(2), push(3), push(4),
pop() -> 4,
push(5),
pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
Example 2:

Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
Output: false
Explanation: 1 cannot be popped before 2.


Constraints:

1 <= pushed.length <= 1000
0 <= pushed[i] <= 1000
All the elements of pushed are unique.
popped.length == pushed.length
popped is a permutation of pushed.
*/

type InTreeNode struct {
	val           int
	previousState *InTreeNode
}

func (t *InTreeNode) Push(val int) *InTreeNode {
	if t == nil {
		return &InTreeNode{
			val: val,
		}
	}

	return &InTreeNode{
		previousState: t,
		val:           val,
	}
}

func (t *InTreeNode) MustPop(expectedVal int) (*InTreeNode, bool) {
	if t == nil {
		return nil, false
	}
	if t.val != expectedVal {
		return nil, false
	}
	return t.previousState, true
}

func ValidateStackSequences(pushed, popped []int) bool {
	var stack *InTreeNode
	return validateStackSequencesRec(stack, pushed, popped)
}

func validateStackSequencesRec(stack *InTreeNode, pushed, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}
	if len(pushed) != 0 {
		if validatePushed(stack, pushed, popped) {
			return true
		}
	}
	if len(popped) != 0 {
		if validatePopped(stack, pushed, popped) {
			return true
		}
	}
	return false
}

func validatePushed(stack *InTreeNode, pushed, popped []int) bool {
	stack = stack.Push(pushed[0])
	return validateStackSequencesRec(stack, pushed[1:], popped)
}

func validatePopped(stack *InTreeNode, pushed, popped []int) bool {
	stack, isValidOp := stack.MustPop(popped[0])
	if !isValidOp {
		return false
	}
	return validateStackSequencesRec(stack, pushed, popped[1:])
}
