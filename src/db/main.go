package main

import (
	"dblib"
	"fmt"
	"os"
)

func main() {
	//TODO : Validate syntax (no ':' in key or value)
	if len(os.Args) == 1 {
		fmt.Println("Usage: db (add|del) key [value]")
		return
	}
	action := os.Args[1]
	key := os.Args[2]
	value := ""
	if len(os.Args) > 3 {
		value = os.Args[3]
	}
	file := "./my_db"
	if action == "add" {
		db.AddEntry(file, key, value)
	} else if action == "del" {
		db.DelEntry(file, key)
	} else if action == "get" {
		fmt.Println(db.Get(file, key))
	} else {
		fmt.Println("Usage: db (add|del|get) key [value]")
		return
	}
}
