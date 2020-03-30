package main

import "fmt"

// 全局常量
const (
	consta = 1
	constb bool = true
)

// 枚举iota
const (
	iotaa = iota
	iotab = iota
	iotac = iota
)

// 局部常量
func main() {
	const c string = "c"
	const d = "d"
	e := "e"
	f, g := "f", "g"
	fmt.Println("c:" + c + ", d:" + d + ", e:" + e + ", f:" + f + ", g:" + g)
}
