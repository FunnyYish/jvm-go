package classpath

import (
	"os"
	"path/filepath"
)

// Classpath 代表classpath抽象
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

//Parse 从命令行参数中解析出classpath
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	//jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)

}
func getJreDir(jreOption string) string {
	if jreOption != "" && exist(jreOption) {
		return jreOption

	}
	if exist("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("无法找到jre目录")
}
func exist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	return self.userClasspath.readClass(className)
}
func (self *Classpath) String() string {
	return self.userClasspath.String()
}
