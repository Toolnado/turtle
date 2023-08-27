package pkg_os

import (
	"os"
	"path/filepath"
)

func ProcessPath(names ...string) string {
	var path []string
	path = append(path, filepath.Dir(os.Args[0]))
	path = append(path, names...)
	return filepath.Join(path...)
}
