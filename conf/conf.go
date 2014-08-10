package conf

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Database     string `json:"database"`
	CookieSecret string `json:"cookie-secret"`
}

var Config config

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Error loading config file: %s\n", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)

	if err != nil {
		log.Fatalf("Config: %s\n", err)
	}
}
