package main

import (
	"fmt"
	"time"
)

func main() {
	overtime()
}

func demo() {
	a := time.After(2 * time.Second)
	<-a
	fmt.Println("timer receive")

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("timer receive")
	})
}

func overtime() {
	c := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		<-c
	}()
	select {
	case c <- 1:
		fmt.Println("channel ..")
	case <-time.After(2 * time.Second):
		close(c)
		fmt.Println("timeout...")
		time.Sleep(time.Second * 4)
	}
}
