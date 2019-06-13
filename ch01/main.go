package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func str2byte(str string)  []byte  {
	a:=(*[2]uintptr)(unsafe.Pointer(&str))
	b:=[3]uintptr{a[0],a[1],a[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func byte2str(b[]byte)string  {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	//a := "hello world"
	//b := []byte(a)
	//fmt.Printf("addr A %p addr B %p\n", &a, b)
	////a b输出的地址不同，所以string2byte对底层数据进行了拷贝
	//c := []byte("hello world")
	//d := string(c)
	//fmt.Printf("addr C %p addr D %p\n", c, &d)
	////c d输出的地址不同，所以byte2string对底层数据进行了拷贝
	str:=strings.Repeat("abc",3)
	b:=str2byte(str)
	str2:=byte2str(b)
	fmt.Println(b,str2)
	//fmt.Printf("addr b %p addr str %p addr str2%p\n", b, &str,&str2)
}
