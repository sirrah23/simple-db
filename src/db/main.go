package main

import (
	"dblib"
	"fmt"
	"os"
)

func main() {
	//TODO : Validate syntax (no ':' in key or value)
	//TODO : Create a usage string printing submodule
	if len(os.Args) == 1 {
		fmt.Println("Usage: db (add|del|get) key [value]")
		return
	}
	action := os.Args[1]
	/*
		key := os.Args[2]
		value := ""
	*/
	file := "./my_db"
	if action == "add" {
		if len(os.Args) < 4 {
			fmt.Println("Usage: db (add|del|get) key [value]")
			return
		}
		key := os.Args[2]
		value := os.Args[3]
		db.AddEntry(file, key, value)
	} else if action == "del" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: db (add|del|get) key [value]")
			return
		}
		key := os.Args[2]
		db.DelEntry(file, key)
	} else if action == "get" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: db (add|del|get) key [value]")
			return
		}
		key := os.Args[2]
		fmt.Println(db.Get(file, key))
	} else if action == "compress" {
		db.Compress(file)
	} else {
		fmt.Println("Usage: db (add|del|get) key [value]")
		return
	}
}
