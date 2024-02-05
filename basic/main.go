package main

import (
	"fmt"
	"time"
)

func main() {
	// time what we're doing to see concurrency in action
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	doggos := []string{"Jazz", "Jura", "Amba", "Bede", "Cassie"}

	/* Naive linear approach
	for _, doggo := range doggos {
		pat(doggo)
	}
	*/

	// Concurrent approach
	for _, doggo := range doggos {
		go pat(doggo)
	}

	time.Sleep(time.Second * 2) // a hack because we might exit main function before we're actually finished with our goroutine
}

func pat(dog string) {
	fmt.Printf("^^ %s patted!  tail wagging!\n", dog)
}
