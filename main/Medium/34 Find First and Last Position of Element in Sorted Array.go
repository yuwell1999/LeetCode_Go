package Medium

func searchRange(nums []int, target int) []int {
	// 二分查找如下值
	// 第一个大于等于target的下标
	// 第一个大于target的下标，然后将下标减1
	leftIndex := binarySearch(nums, target, true)
	rightIndex := binarySearch(nums, target, false) - 1
	if leftIndex <= rightIndex && rightIndex < len(nums) && nums[leftIndex] == target && nums[rightIndex] == target {
		return []int{leftIndex, rightIndex}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, target int, lower bool) int {
	left, right, ans := 0, len(nums)-1, len(nums)
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		if nums[mid] > target || (lower && nums[mid] >= target) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}
