package sort

func BubbleSort(input []int) []int {
	arr := append([]int{}, input...) // 복사본 생성
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func SelectionSort(input []int) []int {
	arr := append([]int{}, input...) // 복사본 생성
	n := len(arr)
	for i := 0; i < n; i++ {
		min := arr[i]
		for j := i + 1; j < n; j++ {
			if min > arr[j] {
				temp := arr[j]
				arr[j] = min
				min = temp
			}
		}
		arr[i] = min
	}
	return arr
}
