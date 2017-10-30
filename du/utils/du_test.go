package utils

import (
	"log"
	"os"
	"testing"
)

func TestDu(t *testing.T) {
	curDir, _ := os.Getwd()
	path := UpperPath(curDir, 1)
	infos, err := Du(path)
	if err != nil {
		t.Error("Du returns error.")
	}

	log.Println("----- Du result -----")
	for _, info := range infos {
		log.Printf("info = %v\n", info)
	}
}
