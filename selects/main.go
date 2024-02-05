package main

import "fmt"

func main() {
	doggo1, doggo2 := make(chan string), make(chan string)

	go makeBossDoggo(doggo1, "hi doggo 1")
	go makeBossDoggo(doggo2, "hi doggo 2")

	// see which doggo becomes the boss doggo
	// the cases are reviewed in a pseudo-random order
	// so while both cases will come out true, the one we run will be quite random
	select {
	case message := <-doggo1:
		fmt.Println(message)
	case message := <-doggo2:
		fmt.Println(message)
	}
}

// send a message to specified channel
func makeBossDoggo(doggo chan string, message string) {
	doggo <- message
}
