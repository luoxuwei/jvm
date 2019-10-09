package classpath

import (
	"os"
	"path/filepath"
)

func newWildcardEntry(name string) CompositeEntry {
	baseDir := name[:len(name)-1]

	compositeEntry := []Entry{}
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		entry := newEntry(path)
		if entry != nil {
			compositeEntry = append(compositeEntry, entry)
		}

		return nil
	}
	filepath.Walk(baseDir, walkFunc)
	return compositeEntry
}
