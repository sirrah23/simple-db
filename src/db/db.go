package main

import (
	"fmt"
	"os"
)

func add_entry(file_name, key, value string) {
	//TODO : Abstract file opening stuff into a Open_Database(...) function
	fhandle, err := os.OpenFile(file_name, os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		if os.IsNotExist(err) {
			fhandle, err = os.Create(file_name)
		}
	}
	checkError(err)
	//TODO : Create a Close_Database(...) function for this
	defer fhandle.Close()
	entry := key + ":" + value + "\n"
	_, err = fhandle.WriteString(entry)
	checkError(err)
}

func del_entry(file_name, key string) {
	//TODO : Abstract file opening stuff into a Open_Database(...) function
	fhandle, err := os.OpenFile(file_name, os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		if os.IsNotExist(err) {
			fhandle, err = os.Create(file_name)
		}
	}
	checkError(err)
	//TODO : Create a Close_Database(...) function for this
	defer fhandle.Close()
	entry := key + "-DELETED\n" // Special tombstone marker
	_, err = fhandle.WriteString(entry)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

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
		add_entry(file, key, value)
	} else if action == "del" {
		del_entry(file, key)
	} else {
		fmt.Println("Usage: db (add|del) key [value]")
		return
	}
}
