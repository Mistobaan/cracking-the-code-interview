package chapterVI

func swap_min_max(array []int) {
	if len(array) < 2 {
		return
	}

	min_idx, max_idx := 0, 0

	for i := 0; i < len(array); i += 1 {
		if array[i] < array[min_idx] {
			min_idx = i
		}

		if array[i] > array[max_idx] {
			max_idx = i
		}
	}

	array[min_idx], array[max_idx] = array[max_idx], array[min_idx]
}
