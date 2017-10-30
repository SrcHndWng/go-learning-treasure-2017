package main

import (
	"log"
	"os"
	"sync"
	"time"

	"./utils"
)

func main() {
	curDir, _ := os.Getwd()
	path := utils.UpperPath(curDir, 1)
	paths, err := utils.SubDirs(path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("----- du start. -----")
	var wg sync.WaitGroup
	var results []utils.DirInfo
	for _, path := range paths {
		wg.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			log.Printf("path = %s\n", path)
			infos, err := utils.Du(path)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, infos...)
			wg.Done()
		}()
		wg.Wait()
	}

	log.Println("----- print results -----")
	for _, result := range results {
		log.Printf("result = %v\n", result)
	}
}
