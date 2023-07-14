package services

import (
	"file-expluwuer/src/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/nexidian/gocliselect"
)

func MenuAddItems(config Configuration, menu *gocliselect.Menu) {
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

	if config.Directory != "/" {
		menu.AddItem("📁 ../", "/..")
	}

	for _, file := range folderElements {
		if file.IsDir() {
			menu.AddItem("📁 ["+utils.CalculateSize(file.Size())+"] "+file.Name()+"/", "/dir"+file.Name())
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

			menu.AddItem(icon+" ["+utils.CalculateSize(file.Size())+"] "+file.Name(), file.Name())
		}
	}

	menu.AddItem("⬅️  Exit", "/exit")
}