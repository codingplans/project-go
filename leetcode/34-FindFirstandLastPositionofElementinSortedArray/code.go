package _00_init_code


func searchRange(arr []int, target int) []int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := int(uint((right + left)) >> 1)
		println(mid, right, left)
		if arr[mid] == target {
			left, right = mid, mid
			for left-1 >= 0 && arr[left-1] == target {
				left--
			}
			for right+1 < len(arr) && arr[right+1] == target {
				right++
			}
			return []int{left, right}
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		println(left, right, mid)

	}
	return []int{-1, -1}
}
