package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Once is an object that will perform exactly one action
// eg if a task can ONLY be done once, but many routines are chasing it (for example) we may use Once

var missionCompleted bool

func main() {
	var waitPls sync.WaitGroup

	var once sync.Once

	for i := 0; i < 10; i++ {
		waitPls.Add(1)
		go func() {
			if foundTreasure() {
				// with the original approach, we will repeatedly try to mark the mission as completed - for no reason
				// however if we call sync.Once.Do() and pass in the func to only do once, we will only call it one time
				once.Do(markMissionCompleted)
			}
			waitPls.Done()
		}()
	}
	waitPls.Wait()
	checkComplete()
}

func markMissionCompleted() {
	missionCompleted = true
}

func checkComplete() {
	if missionCompleted {
		fmt.Println("COMPLETE")
	} else {
		fmt.Println("nah")
	}
}

func foundTreasure() bool {
	return rand.Intn(10) == 0
}
