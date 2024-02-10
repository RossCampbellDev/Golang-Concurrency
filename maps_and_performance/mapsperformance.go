package main

import (
	"fmt"
	"sync"
)

/*
	sync.Map
	a regular map maps some value to another
	the sync package map: helps with concurrency when doing map type stuff

	regular maps do not support concurrency - race condition problems

	sync.Map is more performant than using a regular map with mutexes etc
*/

func main() {
	syncMap := sync.Map{}

	syncMap.Store(0, "hey")
	syncMap.Store(1, "you")
	syncMap.Store(2, "now")

	syncValue, syncOk := syncMap.Load(0)
	fmt.Printf("val is %s and loadOk returned %t\n", syncValue, syncOk)

	// syncMap.Delete(0)

	// syncValue, loaded := syncMap.LoadAndDelete(0) // deletes the value for a key, and returns the previous value if there was one.  'loaded' shows if there was a val to load
	// fmt.Printf("val %s deleted, and loaded returned: %t\n", syncValue, loaded)

	syncValue, loaded := syncMap.LoadOrStore(0, "hey") // if there is no value for the given key, it stores the given one.  if there is one, it returns it into syncValue
	// loaded is true if the value was stored
	// note, it may change the order of the Range() below
	fmt.Printf("loadorstore: %s and %t\n", syncValue, loaded)

	// ranging over a sync.Map is a little more awkward than a regular map
	// remember:  no guarantee of order
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
