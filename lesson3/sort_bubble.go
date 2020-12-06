package lesson3

import "fmt"

func SortBubble(arr []int) []int {
	_len := len(arr)
	a := make([]int, _len)
	t := copy(a, arr)
	fmt.Println(t)

	for j := 1; j <= _len-1; j++ {
		f := true
		for i := 0; i <= _len-1-j; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				f = false
			}
		}
		if f {
			return a
		}
	}
	return a
}
