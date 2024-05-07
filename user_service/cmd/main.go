package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"user-service/web"
)

type Config struct {
	Port string `json:"port"`
}

var config Config

func LoadConfig(fileName string) (Config, error) {
	var currentConfig Config
	configFile, err := os.Open(fileName)
	defer configFile.Close()
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return currentConfig, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&currentConfig)

	if err != nil {
		log.Fatal("Error parsing JSON data:", err)
		return currentConfig, err
	}
	return currentConfig, err
}

func main() {
	var err error
	config, err = LoadConfig("../config.json")
	if err != nil {
		log.Fatal("Error loading config", err)
	}

	mux := web.StartServer()
	err = http.ListenAndServe(":"+config.Port, mux)
	if err != nil {
		log.Fatal("Error loading config", err)
	}

}
