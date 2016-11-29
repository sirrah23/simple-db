package db

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var dbfile = "./test_db"

func TestInsertOne(t *testing.T) {
	defer CleanUp(dbfile)
	key, value := "hello", "world"
	AddEntry(dbfile, key, value)
	insertedValue := Get(dbfile, "hello")
	if insertedValue != "world" {
		t.Fail()
	}
}

func TestInsertTwo(t *testing.T) {
	defer CleanUp(dbfile)
	key1, value1 := "hello", "world"
	key2, value2 := "This is a", "Test"
	AddEntry(dbfile, key1, value1)
	AddEntry(dbfile, key2, value2)
	insertedValue := Get(dbfile, "hello")
	if insertedValue != "world" {
		t.Fail()
	}
	insertedValue = Get(dbfile, "This is a")
	if insertedValue != "Test" {
		t.Fail()
	}
}

func TestGetNotExist(t *testing.T) {
	defer CleanUp(dbfile)
	keyToGet := "hello"
	value := Get(dbfile, keyToGet)
	if value != "" {
		t.Fail()
	}
}

func TestGetMultiple(t *testing.T) {
	// Get should always get the last write
	defer CleanUp(dbfile)
	key1, value1 := "hello", "world"
	key2, value2 := "hello", "friend"
	AddEntry(dbfile, key1, value1)
	AddEntry(dbfile, key2, value2)
	insertedValue := Get(dbfile, "hello")
	if insertedValue != "friend" {
		t.Fail()
	}
}

func TestDeleteExists(t *testing.T) {
	defer CleanUp(dbfile)
	key, value := "hello", "world"
	AddEntry(dbfile, key, value)
	DelEntry(dbfile, "hello")
	fhandle, err := os.Open(dbfile)
	if err != nil {
		panic("Couldn't open the database!")
	}
	fileReader := bufio.NewReader(fhandle)
	data, err := fileReader.ReadString('\n')
	// Second line has deleted update
	data, err = fileReader.ReadString('\n')
	if err != nil {
		// Shouldn't be getting errors when reading the data here...
		t.Fail()
	}
	if data != "hello-DELETED\n" {
		// Key marked as deleted
		t.Fail()
	}
	data, err = fileReader.ReadString('\n')
	if err != io.EOF {
		// Should be nothing else in the file...
		t.Fail()
	}
}

func TestEmptyCompression(t *testing.T) {
	os.Create(dbfile)
	Compress(dbfile)
	output, err := exec.Command("bash", "-c", "ls "+dbfile+"*").Output()
	CheckError(err)
	filesInDir := strings.Split(string(output), "\n")
	filesInDir = filesInDir[:len(filesInDir)-1]
	// New DB file and Backup
	if len(filesInDir) != 2 {
		t.Fail()
	}
	for i := range filesInDir {
		CleanUp(filesInDir[i])
	}
}

func CleanUp(filename string) {
	err := os.Remove(filename)
	if err != nil {
		panic("Could't do a cleanup!!!!!")
	}
}
