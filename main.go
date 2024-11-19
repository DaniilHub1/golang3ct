package main

import (
  f "fmt"
	"sync"
)

var globalCounter int

var mutex sync.Mutex

func increment() {
	mutex.Lock()
	globalCounter++
	mutex.Unlock()
}

func main() {
	var wait sync.WaitGroup

	for i := 0; i < 5; i++ {
		wait.Add(1) 
		go func() {
			increment()
			wait.Done() 
		}()
	}

	wait.Wait()

	f.Println("Результат увелечения глобальной переменной:", globalCounter)
}