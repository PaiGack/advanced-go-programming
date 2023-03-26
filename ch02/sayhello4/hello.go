package main

import "C"

import "fmt"

// go 导出的接口，也可以实现
//export say_hello
func say_hello(s *C.char) {
	fmt.Print(C.GoString(s))
}
