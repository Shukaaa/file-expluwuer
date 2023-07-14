package services

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Configuration struct {
	FirstAction           bool
	Directory             string
	ShowDir               bool
	AllowedExtensions     []string
	BlacklistedExtensions []string
}

func GetConfiguration() (Configuration, error) {
	configuration := Configuration{
		FirstAction:           true,
		Directory:             "",
		ShowDir:               true,
		AllowedExtensions:     []string{"*"},
		BlacklistedExtensions: []string{""},
	}

	for _, arg := range os.Args {
		switch {
		case strings.HasPrefix(arg, "dir="):
			directory := strings.Split(arg, "=")[1]

			absPath, err := filepath.Abs(directory)
			if err != nil {
				return configuration, err
			}

			configuration.Directory = absPath

		case strings.HasPrefix(arg, "showDir="):
			boolValue, err := strconv.ParseBool(strings.Split(arg, "=")[1])
			if err != nil {
				return configuration, err
			}
			configuration.ShowDir = boolValue

		case strings.HasPrefix(arg, "allowedExtensions="):
			configuration.AllowedExtensions = strings.Split(strings.Split(arg, "=")[1], ",")

		case strings.HasPrefix(arg, "blacklistedExtensions="):
			configuration.BlacklistedExtensions = strings.Split(strings.Split(arg, "=")[1], ",")
		}
	}

	return configuration, nil
}
