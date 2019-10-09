package classpath

import (
	"github.com/pkg/errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	composteEntry := []Entry{}

	for _, p := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(p)
		composteEntry = append(composteEntry, entry)
	}
	return composteEntry
}

func (self CompositeEntry) readClass(name string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(name)
		if data != nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("can not find class "+name)
}

func (self CompositeEntry) String() string {
	pathList := make([]string, len(self))
	for i, entry := range self {
		pathList[i] = entry.String()
	}
    return strings.Join(pathList, pathListSeparator)
}


