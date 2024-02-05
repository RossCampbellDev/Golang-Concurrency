package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	mutexes are used to lock resources so that concurrent threads don't access them at the wrong times and cause race conditions etc
	in Go, the Mutex type comes from the sync package
	go routines can acquire the 'lock' by calling mutext.Lock(), then release with Unlock()

	if we have a package-accessible variable that our concurrent routines are accessing/modifying
	then we want to use Lock() before accessing it, and Unlock() after accessing it
*/

var (
	globalInt int
	lock      sync.Mutex
	rwLock    sync.RWMutex // many can hold a read lock, only one can hold write lock
)

func main() {
	var waitPlease sync.WaitGroup

	// if our max value for i is much higher, we'll find it starts to not add up as much as we think it should
	// this is due to not having a mutex, so we are repeating some steps
	// eg one thread may update the value of globalInt AFTER another has read it, but BEFORE it has updated it
	for i := 0; i < 1000; i++ {
		waitPlease.Add(1)
		go increment(&waitPlease)
	}
	waitPlease.Wait()
	fmt.Println(globalInt)

	readAndWrite()
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done() // decrement our delta each time we run this routine
	lock.Lock()     // acquire a lock on the vars we access
	globalInt++
	lock.Unlock()
}

// what the below funcs demonstrate is the use of a read-write lock, as opposed to regular lock
// if we are doing lots of reads, we can use read lock so that many reads can be performed
// without blocking access to the resource for READING.
// WRITE lock access and write access is, however, blocked while there is a read lock
func readAndWrite() {
	go read()
	go read()
	go read()
	go read()
	go write()

	time.Sleep(5 * time.Second)
	fmt.Println("done read and write")
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("read locking")
	time.Sleep(time.Second)
	fmt.Println("read unlock")
}

func write() {
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("write lock")
	time.Sleep(time.Second)
	fmt.Println("write unlock")
}
