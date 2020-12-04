package array

import (
	"reflect"
	"sort"
)

func OnlyOrderDifference(A, B [][]int) bool {
	var resA []int
	var resB []int

	for _, ints := range A {
		resA = append(resA, ints...)
	}

	for _, ints := range B {
		resB = append(resB, ints...)
	}

	sort.Ints(resA)
	sort.Ints(resB)

	return reflect.DeepEqual(resA, resB)
}
