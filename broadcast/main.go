package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	sync.Cond - set up a condition which goroutines wait for before execution
	each Cond has a Locker (a mutex) which must be held when changing the condition
	or calling the wait method
*/

/*
	broadcasting is similar to signal - but for multiple goroutines.  all goroutines are
	waiting for the signal before they will start
*/

var ready bool

func main() {
	ready = false
	readyToGo()
	broadcastStart()
}

func readyToGo() {
	cond := sync.NewCond(&sync.Mutex{}) // returns the address of the new cond object

	go begin(cond)
	cond.L.Lock() // L is our 'locker' - mutex
	for !ready {
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Println("Finito!")
}

func begin(cond *sync.Cond) {
	time.Sleep(time.Second * 3)
	ready = true
	cond.Signal() // wakes up one goroutine that is waiting on cond
}

// don't return until the function has complete
// which only happens once the broadcaster signal has been received
func prepareForBroadcast(fn func(), broadcaster *sync.Cond) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Done()
		broadcaster.L.Lock()
		defer broadcaster.L.Unlock()
		broadcaster.Wait()
		fn()
	}()
	wg.Wait()
}

func broadcastStart() {
	broadcaster := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(3)

	prepareForBroadcast(func() {
		fmt.Println("job 1 done")
		wg.Done()
	}, broadcaster)
	prepareForBroadcast(func() {
		fmt.Println("job 2 done")
		wg.Done()
	}, broadcaster)
	prepareForBroadcast(func() {
		fmt.Println("job 3 done")
		wg.Done()
	}, broadcaster)

	broadcaster.Broadcast() // finally, send a broadcast signal to the cond that all our functions are waiting on
	wg.Wait()               // don't return until we've done our broadcast
	fmt.Println("all done")
}
