package main

import (
	"fmt"
	"kvlite"
	"time"
)

var max = 1000000

// main
func main() {

	//tables := make(map[string]*kvlite.Database)

	db := kvlite.Open("users/")

	defer db.Close()

	// test writing
	start := time.Now()

	for i := 0; i < max; i++ {
		key := "users:" + fmt.Sprint(i)

		db.Put(key, "hello m"+key)

	}

	fmt.Println("put done in :", time.Since(start))

	fmt.Println()

	start = time.Now()

	ln := 0
	for i := 0; i < max; i++ {
		d := db.Get("users:" + fmt.Sprint(i))
		ln += len(d)
	}

	fmt.Println("get done in :", time.Since(start))
	fmt.Println("len data : ", ln)

	time.Sleep(time.Second * 1)

	fmt.Println("data : ", db.Get("users:8899"))

}
