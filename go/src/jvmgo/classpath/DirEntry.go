package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string)  *DirEntry {
    absDir, err := filepath.Abs(path)
    if err != nil {
    	panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(name string) ([]byte, Entry, error)  {
    fileName := filepath.Join(self.absDir, name)
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
    	return nil, self, err
	}
    return data, self, nil
}

func (self *DirEntry) String() string  {
	return self.absDir
}