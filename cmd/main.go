package main

import (
	"fmt"
	"kvlite"
	"time"
)

var max = 10

// main
func main() {

	// first db
	db := kvlite.Open("db1/")
	defer db.Close()

	s := time.Now()
	for i := 0; i < max; i++ {
		key := fmt.Sprint(i)
		db.Set(key, "hello world:"+key)
	}
	fmt.Println(time.Since(s))
	s = time.Now()

	// set data

	l := 0
	for i := 0; i < max; i++ {
		l += len(db.Get(fmt.Sprint(i)))
	}
	fmt.Println(time.Since(s))

	data := db.Get("3")
	fmt.Println("len & data:", l, data)

}
