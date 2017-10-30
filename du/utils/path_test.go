package utils

import (
	"log"
	"os"
	"testing"
)

func TestSubDirs(t *testing.T) {
	curDir, _ := os.Getwd()
	dirPath := UpperPath(curDir, 2)
	paths, err := SubDirs(dirPath)
	if err != nil {
		t.Error("SubDirs returns error.")
	}

	log.Println("----- SubDirs result -----")
	for _, path := range paths {
		log.Printf("path = %s\n", path)
	}
}
