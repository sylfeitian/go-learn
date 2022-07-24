package main

import (
	"fmt"
	"time"
)

func day03() {

	go say("world")
	go say("hello")
	say("134")

	s := []int{7, 3, 5, 6, 3}
	c := make(chan int)
	fmt.Println(s[:len(s)/2])
	fmt.Println(s[len(s)/2:])
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道C中接收

	fmt.Println(x, y, x+y)
}

func say(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum发送到通道 c
}
