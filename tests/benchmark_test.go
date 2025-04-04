package tests

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	mutex := sync.Mutex{}
	var result = 0
	
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Simulate some work by adding numbers
			for j := range 1_000_000 {
				mutex.Lock()
				result += Add(j, j+1)
				mutex.Unlock()
			}
		}
	})

	fmt.Println("Final result:", result)
}

func BenchmarkAdd1(b *testing.B) {
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        localSum := 0  // Each goroutine maintains its own sum
        for pb.Next() {
            // Do just one operation per Next() call
            j := rand.Intn(1000)  // Example value
            localSum += Add(j, j+1)
        }
        fmt.Println("Local sum:", localSum)  // Print the local sum for each goroutine
    })
}