package tests

import (
	"fmt"
	"sync"
)

func Channel() {
	// Create a channel named "ch" for string values
	ch := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- fmt.Sprintf("Hello %d", i)
		}
		close(ch)
	}()

	fmt.Println("Receiving from channel named 'ch':")
	for message := range ch {
		fmt.Println(message)
	}
}

func Mutex() {
	// Create a mutex named "m" for synchronization
	m := &sync.Mutex{}
	// Create a wait group to wait for goroutines to finish
	wg := sync.WaitGroup{}
	// Create a shared variable named "counter"
	counter := 0

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Lock()
			counter++
			m.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value with Mutex:", counter)
}