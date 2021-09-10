package main

import "fmt"
import "sync"

type T int

func main() {
	var slice []T
	var wg sync.WaitGroup

	queue := make(chan T, 1)


	wg.Add(100)
	for i := 0; i < 100; i++ {
	go func(i int) {
	queue <- T(i)
	}(i)
}

	go func() {
	for t := range queue {
	slice = append(slice, t)
	wg.Done()   // ** move the `Done()` call here
	}
	}()

	wg.Wait()

	fmt.Println(slice)
}