package main

import (
  f "fmt"
	"strconv"
	"sync"
)

func main() {
	numbers := make(chan int, 10)
	strings := make(chan string, 10)

	for i := 0; i < 10; i++ {
		numbers <- i
	}
	close(numbers)

	var wait sync.WaitGroup

	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			for num := range numbers {
				strings <- strconv.Itoa(num)
			}
		}()
	}

	go func() {
		wait.Wait()
		close(strings)
	}()

	for str := range strings {
		f.Println(str)
	}
}
