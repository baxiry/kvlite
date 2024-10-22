# kvlite 
kvlite is simple fast embeded kv store,

I discovered by chance while developing the Zaradb that I designed a very similar storage engine with BitCask design.
It was very good .. and it could be relied upon to build a quick, reliable, light database.

The Zara engine was using Arry as a keys in memory, where we could represent _Id with Index directly, the benefit is quick access and reduce memory as much possible. But it was a limited store that could not be used for general purposes.
bitcask engine use HashTable to store keys in memory, the engine becomes useful for general purposes. The only drawback is that it consumes a lot of memory. But this is a small defect compared to its many advantages.
You can see BitCask's advantages through this [paper](https://riak.com/assets/bitcask-intro.pdf).

## usasge:

```go
package main

import (
	"github.com/baxiry/kvlite"
)

func main() {

  db := kvlite.Open("dbName/")
  defer db.Close()

  // insert, or update if key is exist
  db.Put("key", "hello world!")

  // get data by key
  value := db.Get("key")

  println(value) // "hello world"
}
```

## license BSD-3
