package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	Port     int `json:"port"`
	PactPort int `json:"pactPort"`
}

var C = &Config{}

func init() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	mainPath := strings.TrimSuffix(path, "controllers")
	mainPath = strings.TrimSuffix(mainPath, "repository")
	mainPath = strings.TrimSuffix(mainPath, "service")
	mainPath = strings.TrimSuffix(mainPath, "server")
	mainPath = strings.TrimSuffix(mainPath, "pact")

	file, err := os.Open(mainPath + "/.config/" + env + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(read, C)
	if err != nil {
		panic(err)
	}
}
