package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	y := &resource{x: 10}
	go func() {
		defer fmt.Println("done first read")
		y.RLock()
		defer y.RUnlock()
		go func() {
			defer fmt.Println("done first write")
			fmt.Println("first write req")
			y.Lock()
			fmt.Println("after first write granted")
			defer y.Unlock()
		}()
		time.Sleep(time.Second)
		go func() {
			defer fmt.Println("done second read")
			fmt.Println("second read req")
			y.RLock()
			fmt.Println("after second read granted")
			defer y.RUnlock()
		}()
		time.Sleep(10 * time.Second)
	}()
	time.Sleep(time.Minute)
}

type resource struct {
	sync.RWMutex
	x int
}
