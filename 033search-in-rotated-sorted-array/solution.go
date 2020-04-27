package _33search_in_rotated_sorted_array


/**

由于题目要求时间复杂度是logn，则采用二分查找，整体上可以划分为以下两种情况：

nums[left] <= nums[mid]
此时前半部分是升序序列
(1) 若nums[left] <= target <= nums[mid],则表示target在前半部分有序的范围内存在
right = mid - 1
(2) 若target < nums[left] <= nums[mid], 则表示target小于前半部分升序序列中最小值，不能再前半部分序列中查找，需要在后半部分序列中查找
left = mid + 1
(3) 若nums[left] <= nums[mid] < target, 则表示target大于前半部分升序序列中最大值，也不能再前半部分序列中查找，需要在后半部分序列中查找
left = mid + 1

nums[left] > nums[mid] 即：nums[mid] <= nums[right]
此时后半部分序列有序
(1) 若nums[mid] <= target <= nums[right],则表示target在后半部分有序列的范围内存在
left = mid + 1
(2) 若target < nums[mid] <= nums[right],则表示target小于后半部分升序序列中最小值，不能再后半部分序列中查找，需要在前半部分序列中查找
right = mid - 1
(3) 若nums[mid] <= nums[right] < target, 则表示target大于后半部分升序序列中最大值，也不能再后半部分序列中查找，需要在前半部分序列中查找
right = mid - 1

作者：liu-zhuan-nian-shao
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/jian-dan-luo-lie-mei-yi-chong-qing-kuang-by-liu-zh/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right]))  {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

