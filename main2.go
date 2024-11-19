package main

import (
  f "fmt"
	"strconv"
)

func main() {
	numbers := make(chan int, 10)  
	strings := make(chan string, 10)
	
	for i := 0; i < 10; i++ {
		numbers <- i
	}
	close(numbers) 

	go func() {
		for num := range numbers { 
			strings <- strconv.Itoa(num) 
		}
		close(strings) 
	}()

	for str := range strings {
		f.Println(str)
	}
}
