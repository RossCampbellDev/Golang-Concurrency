package main

import (
	"fmt"
	"sync/atomic"
)

/*
	sync.Atomic provides low level primitives for implementing our own synchronisation

*/

func main() {
	var sum int64

	// the AddInt64 line is essentially the same as running:
	/*
		lock mutex
		increment sum
		unlock mutex
	*/
	fmt.Println(sum)
	atomic.AddInt64(&sum, 1) // pass in the variable we are adding to, and then the delta (eg how much to add)
	fmt.Println(sum)

	// alternatively we can establish an atomic variable
	// which the linter recommends over the above, particularly for 32bit CPUs
	var v atomic.Int64
	v.Add(5)

	var newInt int64
	atomic.StoreInt64(&newInt, 1337)
}
