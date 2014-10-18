package conf

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Address        string `json:"address"`
	Database       string `json:"database"`
	SMTP           string `json:"smtp"`
	CookieHashKey  []byte `json:"cookieHashKey"`
	CookieBlockKey []byte `json:"cookieBlockKey`
	Cert           string `json:"cert"`
	Key            string `json:"key"`
	Cache          string `json:"cache"`
}

var Config config

func Load(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Error loading config file: %s\n", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)

	if err != nil {
		log.Fatalf("Config: %s\n", err)
	}
}
