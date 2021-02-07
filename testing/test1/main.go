package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	w := &sync.WaitGroup{}
	w.Add(1)

	w.Add(1)

	w.Add(1)

	w.Add(1)

	go gfunc(w, func() {
		fmt.Println(1)
		time.Sleep(time.Second * 3)
	})
	go gfunc(w, func() {
		fmt.Println(2)
		time.Sleep(time.Second * 4)
	})
	go gfunc(w, func() {
		fmt.Println(3)
		time.Sleep(time.Second * 4)
	})
	go gfunc(w, func() {
		fmt.Println(4)
		time.Sleep(time.Second * 2)
	})
	time.Sleep(time.Second * 1)
	w.Wait()
	fmt.Println("done")
}

func gfunc(w *sync.WaitGroup, fn func()) {
	fn()
	w.Done()
}
