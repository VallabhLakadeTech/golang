package programs

import "fmt"

// "In a slice of int(s), find the last occurance of an element
// [Note: Aim for a complexity of O(log n)]"
func LastOccurence() {

	nums := []int{1, 2, 2, 2, 3, 4, 5, 5, 5, 6, 7}
	target := 6
	fmt.Printf("%v occured last at position %v", target, findLastOccurence(nums, target))
}

func findLastOccurence(nums []int, target int) (position int) {
	start, end := 0, len(nums)-1
	position = -1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			position = mid
			start = mid + 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return

}
