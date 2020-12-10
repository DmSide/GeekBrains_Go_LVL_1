package lesson3

func SortBubble(arr []int) []int {
	tmpArr := make([]int, len(arr))
	copy(tmpArr, arr)

	for j := 1; j < len(arr); j++ {
		f := true
		for i := 0; i < len(arr)-j; i++ {
			if tmpArr[i] > tmpArr[i+1] {
				tmpArr[i], tmpArr[i+1] = tmpArr[i+1], tmpArr[i]
				f = false
			}
		}
		if f {
			return tmpArr
		}
	}
	return tmpArr
}
