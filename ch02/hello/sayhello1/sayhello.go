package main

/*
#include <stdio.h>
static void say_hello(const char* s){
	puts(s);
}
*/
import "C"

func main() {
	C.say_hello(C.CString("Hello, World\n"))
}
