package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type ConfigDir struct {
	Path         string `json:"Path"`
	Extension    string `json:"Extension"`
	FriendlyName string `json:"FriendlyName"`
}

type Config struct {
	Directories []ConfigDir `json:"Directories"`
}

type DirectoryCount struct {
	Count int `json:"Count"`
}

type FileCountResponse struct {
	DirectoriesCounts map[string]DirectoryCount
}

func parseConfig() Config {
	jsonFile, e := os.Open("filecount-api-config.json")
	if e != nil {
		panic(e)
	}
	defer jsonFile.Close()
	var config Config
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)

	return config
}

func listDirectoriesCounts(w http.ResponseWriter, r *http.Request) {
	config := parseConfig()
	directoriesCountsMap := make(map[string]DirectoryCount)
	for _, dir := range config.Directories {
		directoriesCountsMap[dir.FriendlyName] = DirectoryCount{Count: countFilesInDirectory(dir.Path, dir.Extension)}
	}

	response := FileCountResponse{directoriesCountsMap}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleRequests(port string) {
	http.HandleFunc("/", listDirectoriesCounts)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func countFilesInDirectory(directory string, extension string) int {
	files, e := os.ReadDir(directory)

	var count = 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if extension != "" {
			if !strings.HasSuffix(file.Name(), extension) {
				continue
			}
		}
		count++
	}

	if e != nil {
		panic(e)
	}

	return count
}

func main() {
	port := os.Args[1]

	handleRequests(port)
}
