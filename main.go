package main

import (
	"file-expluwuer/src/services"
	"file-expluwuer/src/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	configuration, err := services.GetConfiguration()

	printInfoText(configuration)

	if err != nil {
		log.Fatal(err)
		return
	}

	if configuration.Directory == "" {
		log.Fatal("Please provide a directory")
		return
	}

	run(configuration)
}

func printInfoText(config services.Configuration) {
	fmt.Println("######## file-expluwuer ########")
	fmt.Println("Directory: " + config.Directory)
	fmt.Println("ShowDir: " + strconv.FormatBool(config.ShowDir))
	fmt.Println("AllowedExtensions: " + strings.Join(config.AllowedExtensions, ", "))
	fmt.Println("BlacklistedExtensions: " + strings.Join(config.BlacklistedExtensions, ", "))
	fmt.Println("################################\n\nFiles:")
}

func run(config services.Configuration) {
	fileInfo, err := utils.ReadFiles(config.Directory)
	if err != nil {
		log.Fatal(err)
		return
	}

	var filesList []os.FileInfo
	var directoriesList []os.FileInfo

	for _, file := range fileInfo {
		if file.IsDir() {
			directoriesList = append(directoriesList, file)
			continue
		}

		filesList = append(filesList, file)
	}

	var folderElements []os.FileInfo
	if config.ShowDir {
		folderElementLists := [][]os.FileInfo{directoriesList, filesList}

		for _, folderElementsList := range folderElementLists {
			utils.SortFiles(folderElementsList)
		}

		folderElements = append(directoriesList, filesList...)
	} else {
		utils.SortFiles(filesList)
		folderElements = filesList
	}

	for _, file := range folderElements {
		if file.IsDir() {
			fmt.Println("📁 " + file.Name() + "/")
			continue
		}

		ext := filepath.Ext(file.Name())[1:]

		if utils.IsBlacklistedExtension(config.BlacklistedExtensions, ext) {
			continue
		}

		if utils.IsAllowedExtension(config.AllowedExtensions, ext) {
			icon := "📄"

			switch {
			case utils.IsArchiveExtension(ext):
				icon = "📦"

			case utils.IsAudioExtension(ext):
				icon = "🎵"

			case utils.IsExecutableExtension(ext):
				icon = "🚀"

			case utils.IsFontExtension(ext):
				icon = "🔤"

			case utils.IsImageExtension(ext):
				icon = "🖼️ "

			case utils.IsVideoExtension(ext):
				icon = "🎥"
			}

			fmt.Println(icon + " [" + utils.CalculateSize(file.Size()) + "] " + file.Name())
		}
	}
}
