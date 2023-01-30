package sort

import "sync"

const Threshold = 1 << 20

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	left := MergeSort(items[:len(items)/2])
	right := MergeSort(items[len(items)/2:])
	return merge(left, right)
}

func MergeSortParallel(s []int) []int {
	len := len(s)
	// left := make(chan []int)
	left := []int{}
	if len > 1 {
		middle := len / 2

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			// left <- MergeSortParallel(s[:middle])
			left = MergeSortParallel(s[:middle])
		}()

		// go func() {
		// 	right <- MergeSortParallel(s[middle:])
		// }()

		right := MergeSortParallel(s[middle:])

		wg.Wait()
		// return merge(<-left, right)
		return merge(left, right)
	}
	return s
}

func merge(a []int, b []int) []int {
	Result := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			Result = append(Result, a[i])
			i++
		} else {
			Result = append(Result, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		Result = append(Result, a[i])
	}
	for ; j < len(b); j++ {
		Result = append(Result, b[j])
	}
	return Result
}
