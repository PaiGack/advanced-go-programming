package main

// void say_hello(char* s);
import "C"

import "fmt"

func main() {
	C.say_hello(C.CString("Hello, World\n"))
}

// export 注释的函数名称，必须和 go 函数名称一致（大小写也必须一致）
//export say_hello
func say_hello(s *C.char) {
	fmt.Print(C.GoString(s))
}
