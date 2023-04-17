package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	s2 := s[:0]

	fmt.Println(len(s), cap(s))
	fmt.Println(len(s2), cap(s2))

	fmt.Printf("%p %p\n", &s, &s2)
	s2 = append(s2, 1)
	fmt.Printf("%p %p\n", &s, &s2)

	// fmt.Println(a[len(a)-2 : len(a) : 7]) // panic
	// fmt.Println(a[len(a)-2 : len(a) : 9]) // panic
	fmt.Println(a[len(a)-2 : len(a) : 8])

	a2 := []int{1, 2, 3}
	a2 = a2[:copy(a2, a2[1:])] // 删除开头 1 个元素
	fmt.Println(a2)

	// copy()
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	n := copy(slice2, slice1[4:]) // 只会复制slice1的前3个元素到slice2中
	// copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice2)
	fmt.Println(n)
}

var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}

func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以 int 方式给 float64 排序
	sort.Ints(b)
}

func SortFloat64FastV2(a []float64) {
	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	// 以 int 方式给 float64 排序
	sort.Ints(c)
}
