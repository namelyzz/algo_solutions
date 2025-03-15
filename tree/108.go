package tree

/*
思路：利用二分查找思想递归构建平衡二叉搜索树
关键点1：有序数组的中点是天然的根节点，可以保证左右子树节点数平衡
关键点2：二叉搜索树的中序遍历就是有序数组
关键点3：递归可以自然地处理子树构建问题
关键点4：每次取中点作为根节点，可以保证树的高度最小化（题目要求转换为平衡二叉搜索树，注意平衡这个词）
总的来说，就是考察二叉搜索树的特性

基准情况：如果数组为空，返回空树
寻找中点：取数组中间位置作为根节点
递归构建左右子树
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}
