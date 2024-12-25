package Sort

func MergeSort(a, b []int) []int {
	i := 0
	j := 0
	res := make([]int, 0, len(a)+len(b))
	for i < len(a) || j < len(b) {
		if i >= len(a) {
			res = append(res, b[j:]...)
			break
		}

		if j >= len(b) {
			res = append(res, a[i:]...)
			break
		}
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
			continue
		} else {
			res = append(res, b[j])
			j++
			continue
		}
	}
	return res
}
