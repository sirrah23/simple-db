package db

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Index struct {
	filename string
	indmap   map[string]int64
}

func CreateIndex(filename string) *Index {
	return &Index{filename, make(map[string]int64)}
}

func (i *Index) PopulateIndex() {
	fhandle, err := os.Open(i.filename)
	CheckError(err)
	finfo, err := fhandle.Stat()
	CheckError(err)
	fsize := finfo.Size()
	fileReader := bufio.NewReader(fhandle)
	currKey := ""
	currPos := int64(0)
	for {
		data, err := fileReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		CheckError(err)
		currLoc := currPos
		if strings.Index(data, ":") == -1 {
			// Key has been deleted
			currKey, _ = SplitInTwo(data, "-")
			if _, ok := i.indmap[currKey]; ok {
				delete(i.indmap, currKey)
			}
		} else {
			// Key exists
			currKey, _ = SplitInTwo(data, ":")
			i.indmap[currKey] = currLoc
		}
		buffered := int64(fileReader.Buffered())
		currPos = fsize - buffered
	}
}

func (i *Index) GetVal(key string) (val string) {
	val = ""
	fhandle, err := os.Open(i.filename)
	CheckError(err)
	fileReader := bufio.NewReader(fhandle)
	loc, ok := i.indmap[key]
	if !ok {
		return
	}
	_, err = fileReader.Discard(int(loc)) //-1?
	CheckError(err)
	data, err := fileReader.ReadString('\n')
	CheckError(err)
	if strings.Index(data, ":") == -1 {
		return
	}
	_, val = SplitInTwo(data, ":")
	val = val[:len(val)-1]
	return
}

func (i *Index) AddEntry(key string, loc int64) {
	i.indmap[key] = loc
}
