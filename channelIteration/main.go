package main

import (
	"fmt"
	"math/rand"
)

/*
	we can do something 10 times this way around
	where we loop our call to the concurrent function
func main() {
	channel := make(chan string)
	for i := 0; i < 10; i++ {
		go fetchData(channel)
		fmt.Println(<-channel)
	}
}

func fetchData(channel chan string) {
	score := rand.Intn(10)
	channel <- fmt.Sprint("data: ", score)
}
*/

/*
OR we can do it this way where we loop our receipt FROM the channel
this method lets us have a stack of concurrently-generated values
which we work through back in the main func when they're all ready
*/
func main() {
	channel := make(chan string, 1)
	go fetchData(channel)
	/*
		for i := 0; i < 10; i++ { // extract messages from the channel
			fmt.Printf("value: %s\n", <-channel) // the ordering of 'value' and 'feed' printlns will not be consistent or alternating - they are not synched
		}
	*/

	// if we use this to try and print all messages, we need to close the channel once we've read them all
	// other wise we will arrive at a deadlock, since we're waiting for the loop to end
	// to close the channel, go to the concurrent operation below
	for value := range channel {
		fmt.Printf("value %s\n", value)
	}

	// alternatively:
	for {
		value, open := <-channel
		if !open {
			break
		}
		fmt.Printf("value %s\n", value)
	}
}

func fetchData(channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println("feed in")
		channel <- fmt.Sprint(rand.Intn(10))
	}
	close(channel)
}
