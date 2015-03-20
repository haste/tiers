package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type convertProfile struct {
	Nick   []string `json:"nick"`
	Level  []string `json:"level"`
	AP     []string `json:"ap"`
	Bottom []string `json:"bottom"`
}

var (
	base string
	out  string

	nickResize   int
	levelResize  int
	apResize     int
	bottomResize int

	nickThreshold   int
	levelThreshold  int
	apThreshold     int
	bottomThreshold int

	nickLevel   int
	levelLevel  int
	apLevel     int
	bottomLevel int
)

func init() {
	flag.StringVar(&base, "base", "fallback", "The convert profile to modify")
	flag.StringVar(&out, "output", "fallback", "What to save the profle under.")

	flag.IntVar(&nickResize, "nick-resize", -1, "Nick resize value")
	flag.IntVar(&nickThreshold, "nick-threshold", -1, "Nick threshold value")
	flag.IntVar(&nickLevel, "nick-level", -1, "Nick level value")

	flag.IntVar(&levelResize, "level-resize", -1, "Level resize value")
	flag.IntVar(&levelThreshold, "level-threshold", -1, "Level threshold value")
	flag.IntVar(&levelLevel, "level-level", -1, "Level level value")

	flag.IntVar(&apResize, "ap-resize", -1, "AP resize value")
	flag.IntVar(&apThreshold, "ap-threshold", -1, "AP threshold value")
	flag.IntVar(&apLevel, "ap-level", -1, "AP level value")

	flag.IntVar(&bottomResize, "bottom-resize", -1, "Bottom resize value")
	flag.IntVar(&bottomThreshold, "bottom-threshold", -1, "Top threshold value")
	flag.IntVar(&bottomLevel, "bottom-level", -1, "Bottom level value")

	flag.Parse()
}

func updateValues(data []string, resize, threshold int, level int) {
	var prev string
	for key, value := range data {
		switch prev {
		case "-resize":
			if resize != -1 {
				data[key] = fmt.Sprintf("%d%%", resize)
			}
		case "-threshold":
			if threshold != -1 {
				data[key] = fmt.Sprintf("%d%%", threshold)
			}
		case "-level":
			if level != -1 {
				data[key] = fmt.Sprintf("%d%%", level)
			}
		}

		prev = value
	}
}

func main() {
	var data convertProfile
	path := "ocr/profiles/"

	in, err := os.Open(path + base + ".json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(in)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	in.Close()

	updateValues(data.Nick, nickResize, nickThreshold, nickLevel)
	updateValues(data.Level, levelResize, levelThreshold, levelLevel)
	updateValues(data.AP, apResize, apThreshold, apLevel)
	updateValues(data.Bottom, bottomResize, bottomThreshold, bottomLevel)

	out, err := os.OpenFile(path+out+".json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	out.Write(b)
	out.Write([]byte("\n"))
	out.Close()
}
