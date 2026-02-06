package binarysearch

func BinarySearch(arr []int, target int) int {
	l := 0
	r := len(arr)

	return binarySearchLogic(arr, target, l, r)
}

func binarySearchLogic(arr []int, target int, l int, r int) int {
	for l < r {
		m := (l + r) / 2

		switch {
		case arr[m] == target:
			return m
		case arr[m] < target:
			l = m + 1
		default:
			r = m
		}
	}

	return -1
}
