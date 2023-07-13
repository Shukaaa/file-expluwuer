package utils

import (
	"os"
	"sort"
)

func SortFiles(files []os.FileInfo) []os.FileInfo {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	return files
}
