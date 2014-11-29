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
	TesseractBin   string `json:"tesseractBin"`
	ConvertBin     string `json:"convertBin"`
	PythonBin      string `json:"pythonBin"`
	UtilsDir       string `json:"utilsDir"`
}

var Config config

func init() {
	file, err := os.Open("config.json")

	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)

	if err != nil {
		log.Fatalf("Config: %s\n", err)
	}
}

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
