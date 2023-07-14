package main

import (
	"file-expluwuer/src/services"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/nexidian/gocliselect"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Clear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	Clear()
	configuration, err := services.GetConfiguration()

	if err != nil {
		log.Fatal(err)
		return
	}

	if configuration.Directory == "" {
		log.Fatal("Please provide a directory")
		return
	}

	if _, err := os.Stat(configuration.Directory); os.IsNotExist(err) {
		log.Fatal("Directory does not exist")
		return
	}

	start(configuration)
}

func start(configuration services.Configuration) {
	printInfoText(configuration)

	menu := gocliselect.NewMenu("File Directory")
	services.MenuAddItems(configuration, menu)
	choice := menu.Display()

	Clear()
	configuration.FirstAction = false

	if choice == "/exit" {
		return
	}

	if choice == "/.." {
		configuration.Directory = filepath.Dir(configuration.Directory)
		start(configuration)
		return
	}

	if strings.HasPrefix(choice, "/dir") {
		configuration.Directory = configuration.Directory + "/" + strings.TrimPrefix(choice, "/dir")
		start(configuration)
		return
	}

	// If nothing else, it's a file

	cmd := exec.Command(`explorer`, `/select,`, configuration.Directory+choice)
	cmd.Run()
}

func printInfoText(config services.Configuration) {
	fmt.Println("######## file-expluwuer ########")
	fmt.Println("Directory: " + config.Directory)
	fmt.Println("ShowDir: " + strconv.FormatBool(config.ShowDir))
	fmt.Println("AllowedExtensions: " + strings.Join(config.AllowedExtensions, ", "))
	fmt.Println("BlacklistedExtensions: " + strings.Join(config.BlacklistedExtensions, ", "))
	fmt.Println("################################\n ")
}
