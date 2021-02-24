package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

//通配符类路径不能递归匹配子目录下的JAR文件
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //remove *
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
