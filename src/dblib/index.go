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

func CreateIndex(filename string) Index {
	return Index{filename, make(map[string]int64)}
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
