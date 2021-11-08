package main

import (
	"bradfield/gorocks/rocksdb"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	db := rocksdb.NewRocksDB("/tmp/rocks")
	defer db.Close()
	fmt.Println("Usage:")
	fmt.Println("GET <key>")
	fmt.Println("DEL <key>")
	fmt.Print("PUT <key> <value>\n\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Exiting")
			os.Exit(1)
		}
		line = strings.TrimSpace(line)
		action := strings.Fields(line)
		action_len := len(action)
		if action_len == 0 {
			continue
		}
		switch action[0] {
		case "GET":
			if action_len != 2 {
				fmt.Printf("Invalid GET: contains %d argument(s)", action_len)
				continue
			}
			val, err := db.Get(action[1])
			if err != nil {
				fmt.Printf("Error: %s", err)
				continue
			}
			if val == nil {
				fmt.Printf("\n")
				continue
			}
			fmt.Printf("%s\n", *val)
		case "DEL":
			if action_len != 2 {
				fmt.Printf("Invalid GET: contains %d argument(s)", action_len)
				continue
			}
			err := db.Delete(action[1])
			if err != nil {
				fmt.Printf("Error: %s", err)
				continue
			}
			fmt.Println("ok")
		case "PUT":
			if action_len != 3 {
				fmt.Printf("Invalid GET: contains %d argument(s)", action_len)
				continue
			}
			err := db.Put(action[1], action[2])
			if err != nil {
				fmt.Printf("%s", err)
			}
		default:
			fmt.Println("Invalid command")
		}
	}
}
