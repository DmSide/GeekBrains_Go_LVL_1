package lesson3

func SortInsert(arr []int) []int {
	_len := len(arr)
	a := make([]int, _len)
	copy(a, arr)

	for i := 1; i < _len; i++ {
		x := a[i]
		j := i
		for j > 0 && a[j-1] > x {
			a[j] = a[j-1]
			j--
		}
		a[j] = x
	}
	return a
}
