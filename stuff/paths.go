package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func WalkMatch(root string) ([]string, error) {
	var matches []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		var extArr = [6]string{"*.txt", "*.xlsx"}
		for _, element := range extArr {
			//fmt.Println(element)
			matched, err := filepath.Match(element, filepath.Base(path))
			if err != nil {
				return err
			} else if matched {
				matches = append(matches, path)
			}
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return matches, nil

}
func main() {
	start := time.Now()

	txtFiles, txtError := WalkMatch("C:\\temp\\lazulum")
	if txtError != nil {
		fmt.Print(txtError)
	}
	fmt.Print(txtFiles)
	elapsed := time.Since(start)
	log.Printf("Scan took %s", elapsed)
}
