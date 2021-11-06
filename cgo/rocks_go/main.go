package main

import (
	"bradfield/gorocks/rocksdb"
	"fmt"
)

func main() {
	db := rocksdb.NewRocksDB("/tmp/rocks")
	defer db.Close()
	db.Put("hello", "there")
	db.Put("general", "kenobi")
	fmt.Println(db.Get("hello"))
	fmt.Println(db.Get("general"))
}
