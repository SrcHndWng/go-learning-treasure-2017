package utils

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// UpperPath returns upper path from curDir.
func UpperPath(curDir string, hieraCount int) string {
	path := ""
	ary := strings.Split(curDir, "/")
	max := len(ary) - hieraCount
	for i := 0; i < max; i++ {
		path += ary[i]
		if i < max-1 {
			path += "/"
		}
	}
	return path
}

// SubDirs returns sub directory paths.
func SubDirs(dir string) ([]string, error) {
	var paths []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths, nil
}
