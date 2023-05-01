package construct_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n TreeNode) PrintByDepth() []*int {
	var queue []*TreeNode

	queue = append(queue, &n)

	var res []*int
	for len(queue) != 0 {
		pop := queue[0]
		queue = queue[1:]
		if pop == nil {
			res = append(res, nil)
			continue
		}
		res = append(res, &pop.Val)
		queue = append(queue, pop.Left, pop.Right)
	}
	return res
}