package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// AllTraversals returns preorder, inorder and postorder from ONE walk.
// Each stack frame carries a state counting the visits so far: first visit
// emits preorder and descends left, second emits inorder and descends
// right, third emits postorder and pops.
func AllTraversals(root *TreeNode) (pre, in, post []int) {
	pre, in, post = []int{}, []int{}, []int{}
	if root == nil {
		return
	}
	type frame struct {
		n     *TreeNode
		state int
	}
	stack := []frame{{root, 1}}
	for len(stack) > 0 {
		top := &stack[len(stack)-1]
		switch top.state {
		case 1:
			pre = append(pre, top.n.Val)
			top.state = 2
			if top.n.Left != nil {
				stack = append(stack, frame{top.n.Left, 1})
			}
		case 2:
			in = append(in, top.n.Val)
			top.state = 3
			if top.n.Right != nil {
				stack = append(stack, frame{top.n.Right, 1})
			}
		default:
			post = append(post, top.n.Val)
			stack = stack[:len(stack)-1]
		}
	}
	return
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(1, tn(2, tn(4, nil, nil), tn(5, nil, nil)), tn(3, nil, nil))
	pre, in, post := AllTraversals(root)
	fmt.Println(pre, in, post)
}
