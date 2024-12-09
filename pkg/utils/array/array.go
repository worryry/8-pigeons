package array

import "sort"

func InArray(target string, arr []string) bool {
	sort.Strings(arr)
	index := sort.SearchStrings(arr, target)
	if index < len(arr) && arr[index] == target {
		return true
	}
	return false
}
