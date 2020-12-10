package lesson4

import (
	"fmt"
	"sort"
)

type FibStruct struct {
	fibMap map[int]int
}

func (fib *FibStruct) Init() {
	fib.fibMap = make(map[int]int)
	fib.fibMap[1] = 1
}

func (fib *FibStruct) Fib(n int) int {
	if n <= 0 {
		return 0
	}

	var (
		f1, f2 int
	)

	f1 = fib.fibMap[n-1]
	if f1 == 0 {
		f1 = fib.Fib(n - 1)
		fib.fibMap[n-1] = f1
	}

	f2 = fib.fibMap[n-2]
	if f2 == 0 {
		f2 = fib.Fib(n - 2)
		fib.fibMap[n-2] = f2
	}

	fib.fibMap[n] = f1 + f2
	return f1 + f2
}

func (fib *FibStruct) Print() {
	keys := make([]int, 0, len(fib.fibMap))
	for k := range fib.fibMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%v: %v\n", k, fib.fibMap[k])
	}
}
