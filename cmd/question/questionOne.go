package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		// 將取出的節點從移除
		queue = queue[1:]

		// 交換節點的左右子樹
		node.Left, node.Right = node.Right, node.Left

		// 將交換後的左右子樹節點加入隊列中，以便下一輪交換
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}
