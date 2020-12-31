package lesson2

import (
	"reflect"
	"testing"
)

var GlobalF []int

// ...\Projects\GeekBrains_Go_LVL_1\lesson2>go test -bench=BenchmarkSimpleNumbersV1 -benchmem> v1.txt
func BenchmarkSimpleNumbersV1(b *testing.B) {
	x := []int{1}

	b.SetBytes(1024)
	for i := 0; i < b.N; i++ {
		x = SimpleNumbersV1(10000)
	}

	GlobalF = x
}

// ...\Projects\GeekBrains_Go_LVL_1\lesson2>go test -bench=BenchmarkSimpleNumbersV2 -benchmem> v2.txt
func BenchmarkSimpleNumbersV2(b *testing.B) {
	x := []int{1}

	b.SetBytes(1024)
	for i := 0; i < b.N; i++ {
		x = SimpleNumbersV2(10000)
	}

	GlobalF = x
}

func TestSimpleNumbersV1(t *testing.T) {
	name := "20 Simple Numbers"
	want := []int{2, 3, 5, 7, 11, 13, 17, 19}
	t.Run(name, func(t *testing.T) {
		got := SimpleNumbersV1(20)
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("%s: expected: %v, got: %v", name, want, got)
		}

	})
}

func TestSimpleNumbersV2(t *testing.T) {
	name := "20 Simple Numbers"
	want := []int{2, 3, 5, 7, 11, 13, 17, 19}
	t.Run(name, func(t *testing.T) {
		got := SimpleNumbersV2(20)
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("%s: expected: %v, got: %v", name, want, got)
		}

	})
}
