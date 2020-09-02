package main

import (
	"fmt"
	"time"
)

func first() {
	for i := 1; i < 9; i++ {
		fmt.Println("first", i)
		time.Sleep(time.Millisecond)
	}
}

func second() {
	for i := 1; i < 9; i++ {
		fmt.Println("second", i)
		time.Sleep(time.Millisecond) //program sleeps for 80miliseconds
	}
}

func main() {
	go first()
	go second()
	time.Sleep(time.Millisecond * 50) //program sleeps for 100miliseconds
	fmt.Println("program ended")
}

/*Race condition is a behavior which occurs where the output depends on non-deterministic ordering. In other words, it is dependent on the sequence of the events. Race condition occur due to the change of sequence when communicating between tasks*/
