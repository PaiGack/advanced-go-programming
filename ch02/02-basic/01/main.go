package main

import (
	// static const char* cs = "hello";
	"C"

	"./cgo_helper"
)

func main() {
	cgo_helper.PrintCString(C.cs)
}
