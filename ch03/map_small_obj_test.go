package main

import "testing"

const Size = 500000

func testSmallObjVal()interface{}  {
	m:=make(map[int]int,Size)
	for i:=0;i<Size ;i++  {
		m[i] = i
	}
	return m
}

func testSmallObjPointer()interface{}  {
	m:=make(map[int]*int,Size)
	for i:=0;i<Size ;i++  {
		m[i] = &i
	}
	return m
}

func BenchmarkSmallObjVal(b *testing.B)  {
	for i:=0;i<b.N;i++ {
		testSmallObjVal()
	}
}

func BenchmarkSmallObjPointer(b *testing.B)  {
	for i:=0;i<b.N;i++{
		testSmallObjPointer()
	}
}
