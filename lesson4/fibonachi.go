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
	fib.fibMap[2] = 2
	fib.fibMap[3] = 3
}

func (fib *FibStruct) Fib(n int) int {
	if n <= 0 {
		return 0
	}

	val, ok := fib.fibMap[n]
	if ok {
		return val
	}

	return fib.Fib(n-1) + fib.Fib(n-2)
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
