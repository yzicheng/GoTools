package mid

import (
	"fmt"
	"testing"
)

// 输入： arr = [1,3,5,7,2,4,6,8], k = 4
// 输出： [1,2,3,4]
func Test1714(t *testing.T) {
	arr := []int{1, 3, 5, 7, 2, 4, 6, 8}
	k := 4
	arr = smallestK(arr, k)
	fmt.Println(arr)
}

//	func smallestK(arr []int, k int) []int {
//		i := 1
//		for ; i < len(arr); i++ {
//			j := i
//			for j > 0 && arr[j-1] > arr[j] {
//				arr[j-1], arr[j] = arr[j], arr[j-1]
//				j--
//			}
//		}
//		return arr[:k]
//	}
func smallestK(arr []int, k int) []int {
	quickStore(arr, 0, len(arr)-1)
	return arr[:k]
}
func quickStore(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickStore(arr, low, pivotIndex-1)
		quickStore(arr, pivotIndex+1, high)
	}
}
func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low + 1
	j := high
	for i <= j {
		for i <= j && arr[i] < pivot {
			i++
		}
		for i <= j && arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	arr[low], arr[j] = arr[j], arr[low]
	return j
}
