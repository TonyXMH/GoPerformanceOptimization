package main

import "testing"

func test(m map[int]int) {
	for i := 0; i < 10000; i++ {
		m[i] = i
	}
}
func BenchmarkMap(b *testing.B) {
	for i:=0;i<b.N ;i++  {
		b.StopTimer()
		m:=make(map[int]int)
		b.StartTimer()
		test(m)
	}
}

func BenchmarkCapMap(b*testing.B)  {
	for i:=0;i<b.N ;i++  {
		b.StopTimer()
		m:=make(map[int]int,10000)
		b.StartTimer()
		test(m)
	}
}

//go test -v -bench . -benchmem
//map扩容的代价是数据的拷贝与重新哈希