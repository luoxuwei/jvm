package classpath

import "path/filepath"

type DirEntry struct {
	absDir string
}

func (self *DirEntry) newDirEntry(path string)  *DirEntry {
    absDir, err := filepath.Abs(path)
    if err != nil {
    	panic(err)
	}
	return &DirEntry{absDir}
}