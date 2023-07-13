package utils

import (
	"fmt"
	"os"
)

func ReadFiles(directory string) ([]os.FileInfo, error) {
	files, err := os.Open(directory)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer files.Close()

	fileInfo, err := files.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return fileInfo, nil
}
