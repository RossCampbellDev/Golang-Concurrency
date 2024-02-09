package main

import (
	"fmt"
	"sync"
)

/*
	Pool pattern is for creating a fixed number of... something.  for when resource management is tight
	sync.Pool has 'get' and 'put' functions.
	sync.Pool is used for managing reusable memory - generally memory which is going to be allocated and deallocated a lot

	when we call Put() we are populating the pool with objects that will be reused
	when we call Get() we are trying to retrieve an object from the pool.  if it's not available, we create a new one

	once done with an object, we call Put() again to return it to the pool

	the primary use case is when doing concurrent stuff where we want to reuse objects a lot, but the overall use of
	memory may vary.
	it's important to note that the pool itself, as well as objects in the pool, may be garbage collected
*/

type Person struct {
	Name string
}

func main() {
	pool := &sync.Pool{
		// New is an attribute of sync.Pool
		//	here we set it to be a function which returns an interface
		//	what this does is:
		//	"when the pool is empty and we call sync.Pool.Get(), instead of returning nil
		//	let's print out a statement and then return a new blank Person object"
		New: func() interface{} {
			fmt.Println("creating a new object")
			return &Person{}
		},
	}

	// fill the pool with some objects
	for i := 0; i < 50; i++ {
		obj := &Person{Name: "test"}
		pool.Put(obj)
	}

	// get objects from the pool
	for i := 0; i < 50; i++ {
		obj := pool.Get().(*Person) // get an object of type 'pointer to Person'
		fmt.Printf("retrieved %s\n", obj.Name)
		obj.Name = "" // reset the object's state (for the sake of this example, actual implementations would be different)
		pool.Put(obj) // return the fishy to the pool, blanked
	}
}
