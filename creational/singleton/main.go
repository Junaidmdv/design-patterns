package main

import (
	"fmt"
	"sync"
)

// The Singleton Pattern is a creational design pattern used to ensure that only one instance of a particular
//
//	object exists during the entire runtime of an application, and it provides a global point of access to that instance.

var lock = &sync.Mutex{}

type Database struct {
	//instance
}

var db *Database

func getInstance() *Database {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			db = &Database{}
			fmt.Println("new databse instance is created")
		} else {
			fmt.Println("databse instance is already created")
		}
	} else {
		fmt.Println("databse instance is already created")
	}
	return db
}

func main() {

	for i := 0; i < 100; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
