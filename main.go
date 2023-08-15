package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.RWMutex
	var data int
	var wg sync.WaitGroup

	// Writers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			data++
		}()
	}

	// Readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.RLock()
			defer mutex.RUnlock()
			fmt.Println("Read data:", data)
		}()
	}

	wg.Wait()
}

// package main
//
// import (
// 	"fmt"
// 	"sync"
// )
//
// func gen(nums ...int) <-chan int {
// 	out := make(chan int, len(nums))
// 	go func() {
// 		defer close(out)
// 		for _, v := range nums {
// 			out <- v
// 		}
// 	}()
//
// 	return out
// }
//
// func square(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for v := range in {
// 			out <- v * v
// 		}
// 	}()
//
// 	return out
// }
//
// func display(res <-chan int) {
// 	for v := range res {
// 		fmt.Println("Lmao ", v)
// 	}
// }
//
// func merge(channels ...<-chan int) <-chan int {
// 	out := make(chan int, 1)
// 	var wg sync.WaitGroup
// 	output := func(c <-chan int) {
// 		defer wg.Done()
// 		for v := range c {
// 			out <- v
// 		}
// 	}
//
// 	for _, v := range channels {
// 		go output(v)
// 	}
//
// 	wg.Add(len(channels))
// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
//
// 	return out
// }
//
// func main() {
// 	in := gen(2, 3, 4, 5, 6, 7, 8, 9)
// 	out1 := square(in)
// 	out2 := square(in)
// 	display(merge(out1, out2))
// }
