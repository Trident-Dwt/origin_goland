/*
数组特性。长度固定
值类型一致

作业要求   完成一个日历表
并且按照一周的宽度显示日历表
*/
package main

import (
	"fmt"
)

func main() {

	//a数组的声明
	var a [3]int
	fmt.Println(a)
	//a的赋值
	a = [3]int{1, 2, 3}
	fmt.Println(a)
	//a的修改值
	a = [3]int{11, 22, 3}
	fmt.Println(a)

	personInfo1 := [3]string{"小强", "男", "在职1"}
	personInfo2 := [3]string{"小强", "男", "在职2"}
	personInfo3 := [3]string{"小强", "男", "在职3"}

	fmt.Println(&personInfo1)
	fmt.Println(&personInfo2)
	fmt.Println(&personInfo3)

	newpersonInfo := [...][3]string{
		personInfo1,
		{"小强", "男", "在职1"},
		{"小强", "男", "在职1"},
		{"小强", "男", "在职1"},
	}
	fmt.Println("ddddddddd", newpersonInfo)
	for _, val := range newpersonInfo {
		fmt.Println(val)
	}
	//制作一个日历表
	//year := 0
	//month := 0
	//day := 0
	//
	//var isLeapYear int
	//if year%4 == 0 {
	//	isLeapYear = 1
	//} else {
	//	isLeapYear = 0
	//}

}
