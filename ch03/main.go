package main

import (
	"runtime"
	"runtime/debug"
	"time"
)

//map扩容的代价是数据的拷贝与重新哈希
const capacity = 1000000
var d interface{}

func value()interface{}  {
	m:=make(map[int]int,capacity)
	for i:=0;i<capacity ;i++  {
		m[i]=i
	}
	return m
}

func pointer()interface{}  {
	m:=make(map[int]*int,capacity)
	for i:=0;i<capacity ;i++  {
		m[i]=&i
	}
	return m
}

func gcMap(f func()interface{})  {
	d= f()
	for i:=0;i<20 ;i++  {
		runtime.GC()
		time.Sleep(time.Second)
	}
}

func deleteMap()interface{}  {
	m:=make(map[int]int,capacity)
	for i:=0;i<capacity ;i++  {
		m[i]=i
	}
	for k,_:=range m{
		delete(m,k)
	}
	return m
}
func freeMap(f func()interface{})  {
	//_=f()
	d=f()//删除对象内容但是不会删除所占的内存
	for i:=0;i<20 ;i++  {
		debug.FreeOSMemory()
		time.Sleep(time.Second)
	}
}

func main()  {
	//gcMap(value)
	//gcMap(pointer)
	freeMap(deleteMap)
}
//go build -o test && GODEBUG="gctrace=1" ./test
//如果用map做cache本地内存缓存的话，在内存压力过大的时候，可以定时开辟新的map来释放内存，就想mysql中你删除表的数据，但是表占用的空间也不会变小，只有重建表才能减少。