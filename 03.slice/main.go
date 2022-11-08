package main

import "fmt"

func main() {
	//切片的定义
	var a []int //方式1
	a = []int{1}
	fmt.Println(a)

	var b []int = []int{2} //方式2
	fmt.Println(b)

	c := []int{3} //方式3
	fmt.Println(c)

	//切片增删改查
	//增加
	a = append(c)
	fmt.Println(a, "<><><><>")

}
