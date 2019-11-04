package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("------")
		go TaskA()
		go TaskB()
		go TaskC()
		time.Sleep(3 * time.Second)
	}
}

func TaskA() {
	time.Sleep(1 * time.Second)
	fmt.Println("Task A")
}
func TaskB() {
	time.Sleep(1 * time.Second)
	fmt.Println("Task B")
}
func TaskC() {
	time.Sleep(1 * time.Second)
	fmt.Println("Task C")
}
