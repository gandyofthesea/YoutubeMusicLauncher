package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const url = "https://music.youtube.com"
const configFileName = "config.json"

func main() {
	path := readConfig()
	openBrowser(path)
}

func openBrowser(path string) {
	cmd := exec.Command(path, "--new-window "+url)
	err := cmd.Start()

	if err != nil {
		log.Fatal("Error opening browser: ", err)
	}
}

func readConfig() string {
	if !fileExists() {
		createConfigJson()
	}
	content, err := os.ReadFile(configFileName)

	if err != nil {
		log.Fatal("Config file read error")
	}

	var result string
	err = json.Unmarshal(content, &result)

	if err != nil {
		log.Fatal(err)
	}
	return result
}

func createConfigJson() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the target path: ")
	target, _ := reader.ReadString('\n')
	targetQ := strings.Trim(target, "\r\n\"")

	file, _ := os.Create(configFileName)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	configAsJson, _ := json.Marshal(targetQ)
	_, err := file.Write(configAsJson)

	if err != nil {
		log.Fatal(err)
	}
}

func fileExists() bool {
	_, err := os.Stat(configFileName)
	return !errors.Is(err, os.ErrNotExist)
}
