package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("go\n")
	var x = 1
	fmt.Println(x)
	if x < 5 {
		fmt.Println("x<5")
	}
	var y string
	y = "str"
	fmt.Println(y)
	num := 5
	fmt.Println(num)

	// hour := time.Now().Hour()
	// if hour >= 6 && hour < 12 {
	// 	fmt.Println("朝")
	// } else if hour < 17 {
	// 	fmt.Println("昼")
	// } else {
	// 	fmt.Println("夜")
	// }
	if hour := time.Now().Hour(); hour >= 6 && hour < 12 {
		fmt.Println("朝")
	} else if hour < 19 {
		fmt.Println("昼")
	} else {
		fmt.Println("夜")
	}
	dayOfWeek := "月"
	switch dayOfWeek {
	case "土":
		fmt.Println("大概は休みです。")
	case "日":
		fmt.Println("休み")
	default:
		fmt.Println("仕事や")
	}
}
