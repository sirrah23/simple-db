package main

import (
	"bufio"
	"dblib"
	"fmt"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func cleanNewLine(s string) string {
	return strings.Replace(s, "\n", "", -1)
}

func main() {
	//TODO : Validate syntax (no ':' in key or value)
	//TODO : Create a usage string printing submodule
	if len(os.Args) != 2 {
		fmt.Println("Usage: db <database name>")
		return
	}
	databaseFile := os.Args[1]
	if _, err := os.Stat(databaseFile); os.IsNotExist(err) {
		_, err := os.Create(databaseFile) //fhandle
		CheckError(err)
	}
	reader := bufio.NewReader(os.Stdin)
	var userInput []string
	var index db.Index
	for {
		fmt.Print("db> ")
		text, err := reader.ReadString('\n')
		CheckError(err)
		userInput = strings.Split(text, " ")
		for ind, elem := range userInput {
			userInput[ind] = cleanNewLine(elem)
		}
		action := userInput[0]
		switch action {
		case "add":
			if len(userInput) != 3 {
				fmt.Println("Invalid command...")
			} else {
				key := userInput[1]
				val := userInput[2]
				db.AddEntry(databaseFile, key, val)
			}
		case "get":
			if len(userInput) != 2 {
				fmt.Println("Invalid command...")
			} else {
				key := userInput[1]
				fmt.Println(db.Get(databaseFile, key))
			}
		case "del":
			if len(userInput) != 2 {
				fmt.Println("Invalid command...")
			} else {
				key := userInput[1]
				db.DelEntry(databaseFile, key)
			}
		case "compress":
			if len(userInput) != 1 {
				fmt.Println("Invalid command...")
			} else {
				db.Compress(databaseFile)
			}
		case "quit":
			os.Exit(0)
		case "index":
			index = db.CreateIndex(databaseFile)
			index.PopulateIndex()
			fmt.Println(index)
		default:
			fmt.Println("Invalid command")
		}
	}
}
