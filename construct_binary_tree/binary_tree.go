package construct_binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return BuildTree(preorder, inorder)
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	inOrderPos := map[int]int{}

	for pos, v := range inorder {
		inOrderPos[v] = pos
	}

	fillNode(root, 0, preorder, inOrderPos)

	return root
}

func fillNode(node *TreeNode, pos int, preOrder []int, inOrderPos map[int]int) (int, bool) {
	pos++
	if pos > len(preOrder)-1 {
		return 0, false
	}

	newNode := &TreeNode{
		Val: preOrder[pos],
	}

	if inOrderPos[node.Val] < inOrderPos[newNode.Val] {
		node.Right = newNode
		return fillNode(newNode, pos, preOrder, inOrderPos) // complete the right child
	}

	node.Left = newNode
	pos, hasMore := fillNode(newNode, pos, preOrder, inOrderPos) // complete the left child
	if !hasMore {
		return 0, false
	}

	return fillNode(node, pos, preOrder, inOrderPos) // complete the current node (possibly with a new right child)
}

// inorder left, par, right
// preorder par, left, right
// preorder [3,9,20,15,7] inorder [9,3,15,20,7]
// we read the preorder, first we get the root 3
// we read next -> it can be either the left or the right
// we read from inorder -> we need to find 9 and see whether it's prior to 3 or after 3
//       - if it's prior that means that it's the leftmost node
//       - if it's after that means that it's the rightmost node
