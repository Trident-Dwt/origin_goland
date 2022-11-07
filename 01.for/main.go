/*
本节主要掌握for的简单用法，和break与continue的使用
break：跳出循环
continue：跳过之后的代码，开始执行下一个循环
*/
package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("i am round1")
	}

	j := 0
	for ; j < 5; j++ {
		fmt.Println("i am round2")
	}

	k := 0
	for ; true; k++ {
		fmt.Println("i am round3")

		if k == 5 {
			break
		}
	}

	l := 0
	for {
		fmt.Println("i am round4")
		l++
		if l == 5 {
			break
		}
	}

	m := 0
	for {
		fmt.Println("i am round5")
		m++
		if m > 10 {
			break
		}

		if m%2 == 0 {
			continue
		}

		fmt.Println(m)

	}

}
