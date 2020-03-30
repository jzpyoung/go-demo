package main

import "fmt"

// 全局变量
var (
	vara = 1
	varb bool = true
)

// 局部变量
func main() {
	var varc string = "varc"
	var vard = "vard"
	vare := "vare"
	varf, varg := "varf", "varg"
	fmt.Println("varc:" + varc + ", vard:" + vard + ", vare:" + vare + ", varf:" + varf + ", varg:" + varg)
}
