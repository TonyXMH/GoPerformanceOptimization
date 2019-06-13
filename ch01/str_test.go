package main

import (
	"strings"
	"testing"
)

var str = strings.Repeat("a",1024)

func test1()  {
	b:=[]byte(str)
	_=string(b)
}

func test2()  {
	b:=str2byte(str)
	byte2str(b)
}

func BenchmarkTest1(b *testing.B)  {
	for i:=0;i<b.N ;i++  {
		test1()
	}
}

func BenchmarkTest2(b *testing.B)  {
	for i:=0;i<b.N ;i++  {
		test2()
	}
}
//压测加上内存指标
//go test -v -bench . -benchmem