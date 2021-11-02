package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DemoInfo struct {
	Age  int
	Text string
}

func main() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

func race() {
	done := make(chan bool)
	m := make(map[string]string)
	m["name"] = "world"
	go func() {
		m["name"] = "data race"
		done <- true
	}()
	fmt.Println("Hello,", m["name"])
	<-done
}

func norace() {
	d := &DemoInfo{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func(info *DemoInfo) {
		info.Age = 3
		wg.Done()
	}(d)
	go func(info *DemoInfo) {
		info.Text = "3"
		wg.Done()
	}(d)
	wg.Wait()
	fmt.Println(d)
}
