package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

const directoryName = "gh-mirror"
const configFileName = "gh-mirror.json"

type Config struct {
	SleepDuration int      `json:"sleepDuration"`
	Repos         []string `json:"repos"`
}

func getConfigPath() string {
	configPath := filepath.Join(getRootPath(), configFileName)
	return configPath
}

func createConfig() {
	var config Config
	config.SleepDuration = 30
	config.Repos = []string{}
	writeConfig(config)
	fmt.Printf("Created config file: %v\n", getConfigPath())
}

func readConfig() Config {
	var config Config
	data, err := os.ReadFile(getConfigPath())
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}
	json.Unmarshal(data, &config)
	return config
}

func writeConfig(config Config) {
	configPath := getConfigPath()
	data, _ := json.MarshalIndent(config, "", " ")
	err := os.WriteFile(configPath, data, os.ModePerm)
	if err != nil {
		log.Fatalf("could not write config: %v", err)
	}
}

func addRepoToConfig(repo string) {
	fmt.Println("adding repo", repo)
	config := readConfig()
	for _, cfgRepo := range config.Repos {
		if repo == cfgRepo {
			// prevent duplicates
			return
		}
	}
	config.Repos = append(config.Repos, repo)
	sort.Strings(config.Repos)
	writeConfig(config)
}

func getRootPath() string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		log.Fatal("$HOME is not set")
	}
	rootPath := filepath.Join(homeDir, directoryName)
	return rootPath
}

func createRoot() {
	rootPath := getRootPath()
	err := os.Mkdir(rootPath, os.ModePerm)
	if err != nil {
		log.Fatalf("could not create directory %v: %v", rootPath, err)
	}
	fmt.Printf("Created directory: %v\n", rootPath)
}
