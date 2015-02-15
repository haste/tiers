package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type convertProfile struct {
	Top    []string `json:"top"`
	Bottom []string `json:"bottom"`
}

var (
	base string
	out  string

	topResize    int
	bottomResize int

	topLevel    int
	bottomLevel int
)

func init() {
	flag.StringVar(&base, "base", "fallback", "The convert profile to modify")
	flag.StringVar(&out, "output", "fallback", "What to save the profle under.")

	flag.IntVar(&topResize, "top-resize", -1, "Top resize value")
	flag.IntVar(&topLevel, "top-level", -1, "Top level value")

	flag.IntVar(&bottomResize, "bottom-resize", -1, "Bottom resize value")
	flag.IntVar(&bottomLevel, "bottom-level", -1, "Bottom level value")

	flag.Parse()
}

func updateValues(data []string, resize, level int) {
	var prev string
	for key, value := range data {
		switch prev {
		case "-resize":
			if resize != -1 {
				data[key] = fmt.Sprintf("%d%%", resize)
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

	updateValues(data.Top, topResize, topLevel)
	updateValues(data.Bottom, bottomResize, bottomLevel)

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
