package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

// 这个接口表示类路径项
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

//根据参数内容创建不同类型的entry
func newEntry(path string) Entry {
	//class path 包含;
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	//class path以*结尾
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	//class path以.jar结尾
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	//classpath是个目录
	return newDirEntry(path)
}
