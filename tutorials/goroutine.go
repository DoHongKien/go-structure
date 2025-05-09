package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
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

func downloadFile(url string, ch chan string) {
	// Random duration between 1 and 3 seconds
	duration := time.Duration(rand.IntN(3) + 1) * time.Second

	time.Sleep(duration)
	ch <- fmt.Sprintf("Downloaded %s in %v seconds", url, duration.Seconds())
}

func main2() {
	fileUrls := []string{}
	for i := 1; i <= 1000; i++ {
		fileUrls = append(fileUrls, fmt.Sprintf("file%d.txt", i))
	}

	ch := make(chan string)

	for _, url := range fileUrls {
		go downloadFile(url, ch)
	}

	for range fileUrls {
		result := <-ch
		fmt.Println(result)
	}

	fmt.Println("🎉 All files downloaded")
}
