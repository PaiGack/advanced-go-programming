package main

// void say_hello(const char* s);
import "C"

func main() {
	C.say_hello(C.CString("Hello, world\n"))
}
