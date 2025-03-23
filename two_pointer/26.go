package two_pointer

/*
因为数组是 有序的，所以我们可以利用这一点来避免使用额外的空间。
通过 双指针（slow 和 fast）的方法，我们可以在原数组中进行修改，从而去除重复项。
slow 和 fast 两个指针分别从数组的开头和第二个元素开始。slow 指向当前有效的数组位置，fast 用来遍历整个数组。
当 fast 指针遍历完数组后，slow 指针指向的就是最后一个有效元素的索引，新的数组长度就是 slow + 1。
*/
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	slow, fast := 0, 1
	for fast < n {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow + 1
}
