package array_intersection

func ArrayIntersection(arr1 []int, arr2 []int) []int {
	hash := make(map[int]int)
	n := 0
	for _, v := range arr1 {
		hash[v]++
	}

	for _, v := range arr2 {
		if hash[v] > 0 {
			hash[v]--
			arr1[n] = v
			n++
		}
	}

	return arr1[:n]
}

func ArrayIntersectionSorted(arr1 []int, arr2 []int) []int {
	i, j, n := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			arr2[n] = arr1[i]
			n++
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return arr2[:n]
}
