package main

import (
	"fmt"
	"time"
)

/*
	the synchronisation issue between main and our concurrenct tasks needs resolving.
	the main func may finish execution before our concurrent tasks.
	we may not want this
	to get around it, we might have both main and our tasks accessing shared memory to check and wait upon a flag.
	but this is not ideal.
	instead:  use channels
*/

/*
	channels exist between the two processes in this example.
	the concurrent tasks feed back through the channel and the main func can see this feedback
*/

/* SYNTAX
MAKE a channel -> pass channel to concurrent task -> send values back -> react
we use the left arrow to show data going into the channel:  `myChannel <- someData`
we use it again at the other 'end' of the channel to receive: `<-myChannel`
think:  send TO the channel with an arrow on the right of the channel.  receive FROM with an arrow on the left
*/

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	// make a channel
	myChannel := make(chan bool)

	doggo := "Jura"
	go pat(doggo, myChannel)

	// receive our channel vals
	fmt.Println(<-myChannel)
}

func pat(doggo string, patted chan bool) {
	time.Sleep(time.Second)
	fmt.Printf("Patted %s!  tail wagging!", doggo)
	patted <- true // pass a bool back through our channel
}
