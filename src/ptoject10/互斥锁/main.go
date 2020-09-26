package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	state := make(map[int]int)
	// mutex := new(sync.Mutex)

	for i := 0; i < 100; i++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)
			// mutex.Lock()
			state[key] = val
			// mutex.Unlock()

			fmt.Println(key, state[key])
			// runtime.Gosched()
		}()
	}

	time.Sleep(time.Second)
}
