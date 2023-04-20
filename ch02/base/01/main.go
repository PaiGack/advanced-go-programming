// package main

// /*
// #cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
// #cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
// #cgo linux CFLAGS: -DCGO_OS_LINUX=1

// #if defined(CGO_OS_WINDOWS)
//     const char* os = "windows";
// #elif defined(CGO_OS_DARWIN)
//     const char* os = "darwin";
// #elif defined(CGO_OS_LINUX)
//     const char* os = "linux";
// #else
// #	error(unknown os)
// #endif
// */
// import "C"

// func main() {
// 	// cgo_helper.PrintCString(C.cs)
// 	print(C.GoString(C.os))
// }

package main

/*
struct A {
	int i;
	float f;
};
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A
	fmt.Println(a.i)
	fmt.Println(a.f)
}
