package main

import (
	"fmt"
	"kvlite"
	"time"
)

var max = 100000

// main
func main() {

	// first db
	db := kvlite.Open("db1/")

	l := 0

	s := time.Now()
	for i := 0; i < max; i++ {
		key := "users:" + fmt.Sprint(i)
		db.Set(key, "hello from db1 id"+key)
	}
	fmt.Println(time.Since(s))

	time.Sleep(time.Second)
	s = time.Now()

	l = 0
	for i := 0; i < max; i++ {
		l += len(db.Get("users:" + fmt.Sprint(i)))
	}

	fmt.Println(time.Since(s))

	fmt.Println("\n\ndata & size: ", db.Get("users:2999"), l)
	time.Sleep(time.Second)

}
