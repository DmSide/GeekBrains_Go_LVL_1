package lesson2

func SimpleNumbersV1(max int) []int {
	result := []int{2}
	for i := 3; i < max; i++ {
		c := true
		for _, res := range result {
			if i%res == 0 {
				c = false
			}
		}
		if c {
			result = append(result, i)
		}
	}
	return result
}

func filter(k int, ss []int, test func(int) bool) (ret []int) {
	for ind, s := range ss {
		if test(s) || ind <= k {
			ret = append(ret, s)
		}
	}
	return
}

func SimpleNumbersV2(max int) []int {
	result := make([]int, max-1)
	for i := 0; i < max-1; i++ {
		result[i] = i + 2
	}
	k := 0
	for k < len(result) {
		result = filter(k, result, func(s int) bool { return s%result[k] != 0 })
		k++
	}

	return result
}
