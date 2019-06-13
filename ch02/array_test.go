package main

import "testing"

func BenchmarkArraySmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arraySmall()
	}
}

func BenchmarkArrayLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayLarge()
	}
}
func BenchmarkSliceSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice(SmartSize)
	}
}
func BenchmarkSliceLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice(LargeSize)
	}
}
//go test -v -bench . -benchmem
//小对象在复制上的开销远小于在堆上分配和回收