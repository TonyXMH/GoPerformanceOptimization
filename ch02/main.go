package main

const (
	SmartSize = 1024
	LargeSize = 1024 * 1024
)

func arraySmall() [SmartSize]int {
	ret := [SmartSize]int{}
	for i := 0; i < SmartSize; i++ {
		ret[i] = i
	}
	return ret
}

func arrayLarge() [LargeSize]int {
	ret := [LargeSize]int{}
	for i := 0; i < LargeSize; i++ {
		ret[i] = i
	}
	return ret
}

func slice(size int) []int {
	ret := make([]int, size)
	for i := 0; i < size; i++ {
		ret[i] = i
	}
	return ret
}
