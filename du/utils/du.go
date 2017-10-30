package utils

import (
	"io/ioutil"
	"path/filepath"
)

// DirInfo has a info of directory.
type DirInfo struct {
	path       string
	filesCount int
	size       int64
}

var infos []DirInfo

// Du gets directory path and returns directory info.
func Du(dir string) ([]DirInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var count int
	var size int64
	for _, file := range files {
		if file.IsDir() {
			Du(filepath.Join(dir, file.Name()))
		} else {
			count++
			size += file.Size()
		}
	}
	infos = append(infos, DirInfo{path: dir, filesCount: count, size: size})
	return infos, nil
}
