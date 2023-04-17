package main

import "fmt"

func main() {
	// 数组 字符串 切片 底层数据结构都是一样的
	// 虽然数组元素可以被修改，但是数组本身的赋值和函数传递参数都是以整体复制的方式传递的
	var a [3]int
	app(a)
	// [0 0 0]
	fmt.Println(a)

	var b = [...]int{1, 2, 3}
	var c = [...]int{2: 3, 1: 2}
	var d = [...]int{1, 2, 4: 5, 2: 1, 6}
	// 最后元素位置计算，根据上一个计算
	var e = [...]int{1: 2, 2: 3, 6}
	// [1 2 3] [0 2 3] [1 2 1 6 5] [0 2 3 6]
	fmt.Println(b, c, d, e)

	var p = &a
	fmt.Println(a[0], a[1])
	fmt.Println(p[0], p[1])
	for i, v := range p {
		fmt.Println(i, v)
	}

	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}

	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}

	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1

	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{}
	}()
	<-c2
}

func app(a [3]int) {
	a[0] = 1
	a[1] = 2
	a[2] = 3
}
