package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := make(chan int, 10)
	strings := make(chan string, 10)

	for i := 0; i < 10; i++ {
		numbers <- i
	}
	close(numbers)

	for i := 0; i < 10; i++ {
		go func() {
			for num := range numbers {
				strings <- strconv.Itoa(num)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-strings)
	}
	close(strings)
}
