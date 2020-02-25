package z

import (
	"os"
	"path"
)

func Exists(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func MakeDirs(s ...string) string {
	root := path.Join(s...)
	if !Exists(root) {
		os.MkdirAll(root, os.ModePerm)
	}
	return root
}
