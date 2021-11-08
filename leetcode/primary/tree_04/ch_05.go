package main

import "learner/leetcode/primary/tree_04/common"

/**
	给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

示例 1：

输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：


输入：nums = [1,3]
输出：[3,1]
解释：[1,3] 和 [3,1] 都是高度平衡二叉搜索树。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func sortedArrayToBST(nums []int) *common.TreeNode {
	return helper2(nums,0,len(nums) -1 )
}

func helper2(nums []int,left,right int) *common.TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	node := &common.TreeNode{Val: nums[mid]}
	node.Left = helper2(nums,left,mid - 1)
	node.Right = helper2(nums,mid + 1, right)
	return node
}
