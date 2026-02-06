package twopointer

func TwoPointer[T any](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	l := 0
	r := len(arr) - 1

	for l <= r {
		hold := arr[l]
		arr[l] = arr[r]
		arr[r] = hold

		l += 1
		r -= 1
	}

	return arr
}
