package db

import (
	"bufio"
	"io"
	"os"
	"testing"
)

var dbfile = "./test_db"

func TestInsertOne(t *testing.T) {
	defer CleanUp()
	key, value := "hello", "world"
	AddEntry(dbfile, key, value)
	insertedValue := Get(dbfile, "hello")
	if insertedValue != "world" {
		t.Fail()
	}
}

func TestInsertTwo(t *testing.T) {
	defer CleanUp()
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
	defer CleanUp()
	keyToGet := "hello"
	value := Get(dbfile, keyToGet)
	if value != "" {
		t.Fail()
	}
}

func TestDeleteExists(t *testing.T) {
	defer CleanUp()
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

func CleanUp() {
	err := os.Remove(dbfile)
	if err != nil {
		panic("Could't do a cleanup!!!!!")
	}
}
