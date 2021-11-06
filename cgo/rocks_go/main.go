package main

import (
	"bradfield/gorocks/rocksdb"
	"fmt"
	"log"
)

func main() {
	db := rocksdb.NewRocksDB("/tmp/rocks")
	defer db.Close()
	err := db.Put("hello", "there")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Put("general", "kenobi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db.Get("hello"))
	fmt.Println(db.Get("general"))
}
