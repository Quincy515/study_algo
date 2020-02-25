package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(n)
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil { // 如果树为空
		return [][]int{} // 返回空的list
	}
	var result [][]int  // 定义要返回的结果
	var q []*TreeNode   // 定义辅助队列
	q = append(q, root) // 同时把根节点入队

	for len(q) > 0 { // 当队列不为空时，循环以下操作
		var elem []int              // 定义list保存这一层级访问到的元素
		size := len(q)              // 然后把当前队列大小取出,表示的是当前层级中节点的个数
		for i := 0; i < size; i++ { // 使用for循环执行以下操作
			s := q[0] // 把队首的节点出队
			q = q[1:]
			elem = append(elem, s.Val) // 把队首的节点数值加入list
			if s.Left != nil {         // 如果左子树不为空
				q = append(q, s.Left) // 则把左子树入队
			}
			if s.Right != nil { // 如果右子树不为空
				q = append(q, s.Right) // 则把右子树入队
			}
		}
		// for循环结束后，说明这一层级节点已经访问完
		// 把保存访问结果的list加入到result
		result = append(result, elem)
	}
	// 循环结束后，result保存的是二叉树层序遍历的结果
	// 使用for循环，把沿中心轴对称的元素对调，以此来反转数组
	for i := 0; i < len(result)/2; i++ {
		j := len(result) - i - 1
		result[i], result[j] = result[j], result[i]
	}
	return result // 最后返回result即可
}
