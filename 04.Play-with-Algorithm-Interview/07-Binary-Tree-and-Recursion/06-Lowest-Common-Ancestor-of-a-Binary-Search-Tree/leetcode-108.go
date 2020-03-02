package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 辅助函数 输入是原整数数组以及每次构建子树时需要用到的子数组的开始和结束下标，输出是构建出来的二叉搜索树
func arrayToBST(nums []int, start, end int) *TreeNode {
	if start > end { // 如果开始下标大于结束下标
		return nil // 返回空指针
	}
	mid := start + (end-start)/2               // 否则计算开始下标和结束下标的中间值mid
	root := &TreeNode{Val: nums[mid]}          // 并用下标mid对应的数字构建一个树节点root
	root.Left = arrayToBST(nums, start, mid-1) // 然后用mid左边的子数组构建root的左子树,这里更新结束下标为mid-1即可
	root.Right = arrayToBST(nums, mid+1, end)  // 同理用mid右边的子数组构建root的右子树
	return root                                // 最后返回root即可
}

// 递归方法 Time: O(n), Space: O(log(n))
func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil { // 如果数组为空
		return nil // 返回空指针
	}
	return arrayToBST(nums, 0, len(nums)-1)
}

// 定义保存状态信息的数据结构
type Item struct { // 每个Item包含
	Parent     *TreeNode // 一个父亲节点
	Start, End int       // 用于构建子树的数组开始下标和结束下标
	IsLeft     bool      // 用于标识构建的子树是左子树还是右子树
}

func NewItem(parent *TreeNode, start, end int, isLeft bool) *Item {
	return &Item{parent, start, end, isLeft}
}

// 迭代方法 Time: O(n), Space: O(n)
func sortedArrayToBST1(nums []int) *TreeNode {
	if nums == nil { // 如果数组为空
		return nil // 返回空指针
	}
	var stack []*Item    // 定义辅助栈
	dummy := &TreeNode{} // 定义dummy父节点
	// 连同数组的开始结束下标一起入栈，这里把要构建的二叉搜索树作为dummy节点的左子树
	stack = append(stack, NewItem(dummy, 0, len(nums)-1, true))

	for len(stack) > 0 { // 当栈不为空时，不断执行以下操作
		item := stack[len(stack)-1] // 先出栈一个item
		stack = stack[:len(stack)-1]
		if item.Start <= item.End { // 如果它对应的开始下标小于等于结束下标
			mid := item.Start + (item.End-item.Start)/2 // 则计算这两个下标的中间下标
			child := &TreeNode{Val: nums[mid]}          // 利用mid指向的数字构建一个树节点
			if item.IsLeft {                            // 如果当前构建的是左子树
				item.Parent.Left = child // 则把当前构建的节点加到parent的左节点上
			} else { // 否则把它加到parent的右节点上
				item.Parent.Right = child
			}
			// 要把mid左边的子数组构建当前child节点的左子树，于是把这些信息入栈
			stack = append(stack, NewItem(child, item.Start, mid-1, true))
			stack = append(stack, NewItem(child, mid+1, item.End, false)) // 同理把构建右子树的信息也入栈
		}
	}
	return dummy.Left // 把栈里的Item处理完后，返回dummy节点的Left即可
}
