package peakindexinmountainarray

func peakIndexInMountainArray(arr []int) int {
	var l int
	var h = len(arr) - 1

	for l < h {
		mid := (l + h) / 2
		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			h = mid
		}
	}

	return l
}
