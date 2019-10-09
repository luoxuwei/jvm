package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath Entry
	userClassPath Entry
}

func Parse(jreOption, cpOption string) *ClassPath {
	classPath := ClassPath{}
	classPath.parseBootAndExtClasspath(jreOption)
    classPath.parseUserClasspath(cpOption)
	return &classPath
}

func (self ClassPath) parseBootAndExtClasspath(jreOption string) {
    jreDir := getJreDir(jreOption)
    jreLibPath := filepath.Join(jreDir, "lib", "*")
    self.bootClassPath = newWildcardEntry(jreLibPath)
    jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
    self.extClassPath = newWildcardEntry(jreExtPath)
}

func (self ClassPath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "./"
	}
	self.userClassPath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre");
	}

	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false;
		}
	}
	return true
}