package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	//TODO : Check if key is in database before delete
	entry := key + "-DELETED\n" // Special tombstone marker
	_, err = fhandle.WriteString(entry)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Split_In_Two(str_to_split, sep string) (string, string) {
	s := strings.Split(str_to_split, sep)
	return s[0], s[1]

}

func get(file_name, key string) string {
	fhandle, err := os.Open(file_name)
	if err != nil {
		if os.IsNotExist(err) {
			fhandle, err = os.Create(file_name)
		}
	}
	checkError(err)
	defer fhandle.Close()
	file_reader := bufio.NewReader(fhandle)
	for {
		data, err := file_reader.ReadString('\n')
		if err == io.EOF {
			// TODO: Return error here?
			return ""
		}
		curr_key, curr_val := Split_In_Two(data, ":")
		if curr_key == key {
			return curr_val[:len(curr_val)-1]
		}
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
	} else if action == "get" {
		fmt.Println(get(file, key))
	} else {
		fmt.Println("Usage: db (add|del|get) key [value]")
		return
	}
}
