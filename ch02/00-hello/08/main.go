package main

import (
	// _GoString_ 预定义的 C 语言类型
	// 注释要和主体部分分开

	// void SayHello(_GoString_ s);
	"C"
	"fmt"

	"runtime"
	"strconv"
	"strings"
)

func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}

func main() {
	fmt.Printf("main %d\n", GetGoid())
	C.SayHello("Hello, World\n")
	fmt.Printf("main %d\n", GetGoid())
}

// export 紧跟 // 不能有空格
//
//export SayHello
func SayHello(s string) {
	fmt.Printf("SayHello %d\n", GetGoid())
	fmt.Print(s)
	fmt.Printf("SayHello %d\n", GetGoid())
}
