package main

import (
	"fmt"
)

func fatRate(h float64, w float64, y int, ss string) {
	var BIM float64
	BIM = w / (h * h)
	var sex int
	if "男" == ss {
		sex = 1
	} else {
		sex = 0
	}
	var result = (1.2*BIM + 0.23*float64(y) - 5.4 - 10.8*float64(sex)) / 100

	if 1 == sex {
		if result < 0.1 {
			fmt.Println("太瘦了")
		} else if result >= 0.1 && result <= 0.5 {
			fmt.Println("还行，保持")
		} else {
			fmt.Println("没救了，等死吧")
		}
	} else {
		if result < 0.15 {
			fmt.Println("太瘦了")
		} else if result >= 0.15 && result <= 0.45 {
			fmt.Println("还行，保持")
		} else {
			fmt.Println("没救了，等死吧")
		}
	}
}

func main() {
	for {
		var h float64
		fmt.Println("height：")
		fmt.Scanln(&h)
		var w float64
		fmt.Println("weight：")
		fmt.Scanln(&w)
		var y int
		fmt.Println("age：")
		fmt.Scanln(&y)
		var ss string
		fmt.Println("sex：")
		fmt.Scanln(&ss)
		fatRate(h, w, y, ss)
		fmt.Println("are you want continue?")
		var aaa string
		fmt.Scanln(&aaa)

		if aaa != "y" {
			break
		}
	}
}
