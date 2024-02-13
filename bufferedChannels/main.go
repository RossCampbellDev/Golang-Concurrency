package main

import "fmt"

func main() {
	// if the channel were of size 1 (or the default of 0) this code would cause a deadlock
	// because main would be waiting to use the channel, yet waiting on the channel to have it's first value received
	// this can be fixed using buffering - eg the size of 2.
	// this allows us to push 2 messages into our channel before it is "full" and deadlocks our main func
	myChannel := make(chan string, 2)
	myChannel <- "First Message"
	myChannel <- "Second Message"

	// fmt.Println(<-myChannel)
	// fmt.Println(<-myChannel)

	close(myChannel)
	for msg := range myChannel {
		fmt.Println(msg)
	}
}
