package main

import (
	"fmt"
	"sync"
)

// in order to actually reach the 'catch' func call, we need to use a waitgroup
// otherwise the wmain func will finish execution before we reach 'catch' and the program will end.

// to fix this we use a waitgroup to handle synchronising our tasks and halting execution until we're ready to proceed
func main() {
	var waitPlease sync.WaitGroup
	waitPlease.Add(1)
	go catch("jim", &waitPlease)
	waitPlease.Wait()

	// if we are working with multiple concurrent calls (duh!)
	names := []string{"bert", "john", "davey"}
	var wait2 sync.WaitGroup
	for _, name := range names {
		wait2.Add(1)
		go catch(name, &wait2)
	}
	wait2.Wait()

	// what if we end up indefinitely waiting?
	var wait3 sync.WaitGroup
	wait3.Add(1)
	//wait3.Done() after Add and before Wait would avoid indefinitely waiting
	//wait3.Done() for a second time makes our Delta negative, which causes an error
	wait3.Wait()
}

func catch(rabbit string, waitPlease *sync.WaitGroup) {
	fmt.Println("chased down rabbit: ", rabbit)
	waitPlease.Done()
}
