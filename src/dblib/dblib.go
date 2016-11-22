package db

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func AddEntry(filename, key, value string) {
	//TODO : Abstract file opening stuff into a Open_Database(...) function
	fhandle, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		if os.IsNotExist(err) {
			fhandle, err = os.Create(filename)
		}
	}
	CheckError(err)
	//TODO : Create a Close_Database(...) function for this
	defer fhandle.Close()
	entry := key + ":" + value + "\n"
	_, err = fhandle.WriteString(entry)
	CheckError(err)
}

func DelEntry(filename, key string) {
	//TODO : Abstract file opening stuff into a Open_Database(...) function
	fhandle, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		if os.IsNotExist(err) {
			fhandle, err = os.Create(filename)
		}
	}
	CheckError(err)
	//TODO : Create a Close_Database(...) function for this
	defer fhandle.Close()
	//TODO : Check if key is in database before delete
	entry := key + "-DELETED\n" // Special tombstone marker
	_, err = fhandle.WriteString(entry)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func SplitInTwo(strtosplit, sep string) (string, string) {
	s := strings.Split(strtosplit, sep)
	return s[0], s[1]

}

func Get(filename, key string) string {
	fhandle, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fhandle, err = os.Create(filename)
		}
	}
	CheckError(err)
	defer fhandle.Close()
	file_reader := bufio.NewReader(fhandle)
	for {
		data, err := file_reader.ReadString('\n')
		if err == io.EOF {
			// TODO: Return error here?
			return ""
		}
		curr_key, curr_val := SplitInTwo(data, ":")
		if curr_key == key {
			return curr_val[:len(curr_val)-1]
		}
	}
}
