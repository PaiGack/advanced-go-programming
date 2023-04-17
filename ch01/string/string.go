package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// type StringHeader struct {
//  Data uintptr
//  Len int
// }
//

func main() {
	var data = [...]byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd'}
	fmt.Println(data)

	s := "hello, world"
	hello := s[:5]
	world := s[7:]
	fmt.Println(hello, world)

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len)
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len)

	fmt.Printf("%#v\n", []byte("Hello, 世界"))
	fmt.Printf("%#v\n", []rune("Hello, 世界"))

	fmt.Println("\xe4\xb8\x96") // 打印: 世
	fmt.Println("\xe7\x95\x8c") // 打印: 界

	fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc") // �界abc

	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
		fmt.Println(i, c)
	}
	// 0 65533  // \uFFFD, 对应 �
	// 1 0      // 空字符
	// 2 0      // 空字符
	// 3 30028  // 界
	// 6 97     // a
	// 7 98     // b
	// 8 99     // c

	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}

}
