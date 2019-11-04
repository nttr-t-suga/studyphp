package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "KUTSUZAWA is the GOD"
	split := strings.Split(str, "")
	fmt.Println(split[0])
	// for i := 0; i < 10; i++ {
	// 	fmt.Println()
	// }
	for _, s := range split {
		fmt.Println(s)
	}

}
