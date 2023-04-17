package main

// #include <hello.h>
import "C"

func main() {
	C.say_hello(C.CString("Hello, world\n"))
}
