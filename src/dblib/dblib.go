package db

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func AddEntry(filename, key, value string) int64 {
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
	fileinfo, err := fhandle.Stat()
	CheckError(err)
	loc := fileinfo.Size() + 1
	entry := key + ":" + value + "\n"
	_, err = fhandle.WriteString(entry)
	CheckError(err)
	return loc // Location where entry was written to in file
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
	defer fhandle.Close()
	if err != nil {
		if os.IsNotExist(err) {
			fhandle, err = os.Create(filename)
			return ""
		}
	}
	CheckError(err)
	curr_key, curr_val := "", ""
	var val string
	file_reader := bufio.NewReader(fhandle)
	for {
		data, err := file_reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if strings.Index(data, ":") == -1 {
			// Key has been deleted
			curr_key, curr_val = SplitInTwo(data, "-")
			if curr_key == key {
				val = ""
			}
		} else {
			// Key exists
			curr_key, curr_val = SplitInTwo(data, ":")
			if curr_key == key {
				val = curr_val[:len(curr_val)-1]
			}
		}
	}
	return val
}

// Function for testing purposes
func printFile(filename string) {
	fhandle, _ := os.Open(filename)
	defer fhandle.Close()
	file_reader := bufio.NewReader(fhandle)
	for {
		data, err := file_reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(data)
	}
}

func Compress(filename string) {
	fhandleOrig, err := os.Open(filename)
	CheckError(err)
	defer fhandleOrig.Close()
	keyValMap := make(map[string]string)
	file_reader := bufio.NewReader(fhandleOrig)
	for {
		data, err := file_reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if strings.Index(data, ":") == -1 {
			// Key was deleted
			curr_key, _ := SplitInTwo(data, "-")
			delete(keyValMap, curr_key)
		} else {
			//Key exists
			curr_key, curr_val := SplitInTwo(data, ":")
			// TODO: Get AddEntry to clean input?
			keyValMap[curr_key] = curr_val[:len(curr_val)-1]
		}

	}
	filenameTemp := filename + "_temp"
	_, err = os.Create(filenameTemp)
	CheckError(err)
	for k, v := range keyValMap {
		AddEntry(filenameTemp, k, v) //Add compressed Key-Value pairs to new database
	}
	// TODO: Create a function go generate timestamp strings
	err = os.Rename(filename, filename+"."+genTimeStamp())
	// TODO: Compress the backup file that gets created
	CheckError(err)
	err = os.Rename(filenameTemp, filename) //New file with old writes removed!
	CheckError(err)
}

func genTimeStamp() string {
	//Timestamp of the format YYYYMMDD.HHMMSS
	return time.Now().Format("20060102.0345")
}
