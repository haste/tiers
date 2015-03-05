package ocr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"image"
	_ "image/jpeg"
	_ "image/png"

	"tiers/conf"
	"tiers/id"
	"tiers/profile"
)

type innovator struct {
	Rank  int     `json:"rank"`
	Good  float32 `json:"good"`
	Total int     `json:"total"`
}

type OCRProfile struct {
	Top    []string `json:"top"`
	Bottom []string `json:"bottom"`
}

var options = map[int]OCRProfile{}

func InitConfig() {
	path := conf.Config.Workdir + "ocr/profiles/"
	dir, err := os.Open(path)
	if err != nil {
		log.Fatalf("ocr init: %s\n", err)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal("ocr init: %s\n", err)
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			var (
				data  OCRProfile
				width int
			)

			name := file.Name()
			tmp := name[0 : len(name)-5]
			if tmp == "fallback" {
				width = 0
			} else if tmp == "template" {
				continue
			} else {
				width, err = strconv.Atoi(tmp)
				if err != nil {
					log.Fatalf("ocr init: %s\n", err)
				}
			}

			dataFile, err := os.Open(path + name)
			if err != nil {
				log.Fatalf("ocr init: %s\n", err)
			}
			defer dataFile.Close()

			decoder := json.NewDecoder(dataFile)
			err = decoder.Decode(&data)
			if err != nil {
				log.Fatalf("ocr init: %s\n", err)
			}

			options[width] = data
		}
	}
}

func sanitizeNum(input []byte) int64 {
	input = regexp.MustCompile(`[lL|\]JI]`).ReplaceAll(input, []byte("1"))
	input = regexp.MustCompile(`[Oo]`).ReplaceAll(input, []byte("0"))

	n := string(input)
	n = strings.Replace(n, "B", "8", -1)
	n = strings.Replace(n, "n", "11", -1)
	n = strings.Replace(n, "H", "11", -1)
	n = strings.Replace(n, ",", "", -1)

	un, _ := strconv.ParseInt(n, 10, 0)
	return un
}

func matchString(res []byte, pattern string) string {
	r := regexp.MustCompile(pattern)

	if r.Match(res) != true {
		return ""
	}

	return string(r.FindSubmatch(res)[1])
}

func matchNum(res []byte, pattern string) int64 {
	r := regexp.MustCompile(pattern)

	if r.Match(res) != true {
		return 0
	}

	return sanitizeNum(r.FindSubmatch(res)[1])
}

func genMatchNum(res []byte, s string) int64 {
	s = regexp.MustCompile(`[Ss]`).ReplaceAllLiteralString(s, "[Ss]")
	s = regexp.MustCompile(`[ae]`).ReplaceAllLiteralString(s, "[aeE8B]")
	s = regexp.MustCompile(`[Pp]`).ReplaceAllLiteralString(s, "[Pp]")
	s = regexp.MustCompile(`[D0]`).ReplaceAllLiteralString(s, "[D0]")
	s = regexp.MustCompile(`[Oou]`).ReplaceAllLiteralString(s, "[0Oou]")
	s = regexp.MustCompile(`[Cc]`).ReplaceAllLiteralString(s, "[Cc]")
	s = regexp.MustCompile(`[ltI]`).ReplaceAllLiteralString(s, "[l|1tI]")
	s = regexp.MustCompile(`\s+`).ReplaceAllLiteralString(s, `\s*`)

	s = strings.Replace(s, `Hn`, "im", -1)
	s = strings.Replace(s, `-`, ".", -1)
	s = strings.Replace(s, `#`, `([0-9LIlJBOonH|,\]]+)`, -1)

	return matchNum(res, s)
}

type OCR struct {
	Filename   string
	OCRProfile int
	Profile    profile.Profile

	tmpName   string
	innovator []byte
}

func (ocr *OCR) Split() {
	ocr.tmpName = id.New()

	cv := exec.Command(conf.Config.PythonBin, []string{
		conf.Config.UtilsDir + "innovator-crop/crop.py",
		conf.Config.Cache,
		ocr.Filename,
		ocr.tmpName,
	}...)

	res, err := cv.Output()
	if err != nil {
		log.Fatal("cv ", err)
	}

	ocr.innovator = res
}

func (ocr *OCR) getArguments(part string, width int) []string {
	var (
		profile OCRProfile
		args    []string
	)

	fileName := conf.Config.Cache + "/" + ocr.tmpName + "_" + part + ".png"
	// non-zero width means we already have a profile defined.
	if width == 0 {
		reader, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()

		m, _, err := image.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}

		width = m.Bounds().Dx()
	}

	if opt, ok := options[width]; ok {
		profile = opt
	} else {
		profile = options[0]
	}

	switch part {
	case "top":
		args = append(args, profile.Top...)
	case "bottom":
		args = append(args, profile.Bottom...)
	}

	return append(args, fileName)
}

func (ocr *OCR) mogrify(kind string) {
	args := ocr.getArguments(kind, ocr.OCRProfile)
	mogrify := exec.Command(conf.Config.ImageMagickBin+"mogrify", args...)

	err := mogrify.Run()
	if err != nil {
		log.Fatal("convert ", err)
	}
}

func (ocr *OCR) tesseract(fileName string) []byte {
	tesseract := exec.Command(conf.Config.TesseractBin, []string{
		"-psm",
		"4",
		"-l", "eng",
		fileName,
		"stdout",
		"ingress",
	}...)

	res, err := tesseract.Output()
	if err != nil {
		log.Fatal("tesseract ", err)
	}

	return res
}

func (ocr *OCR) ProcessInnovator() {
	var data innovator
	decoder := json.NewDecoder(bytes.NewReader(ocr.innovator))
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal("json ", err)
	}

	p := &ocr.Profile
	if data.Good > 0 && data.Rank >= 0 {
		p.InnovatorLevel = profile.BadgeRanks["Innovator"][data.Rank]
	}
}

func (ocr *OCR) ProcessTop() {
	fileName := conf.Config.Cache + "/" + ocr.tmpName + "_top.png"
	ocr.mogrify("top")
	top := ocr.tesseract(fileName)
	ocr.buildProfileTop(top)
}

func (ocr *OCR) ProcessBottom() {
	fileName := conf.Config.Cache + "/" + ocr.tmpName + "_bottom.png"
	ocr.mogrify("bottom")
	bottom := ocr.tesseract(fileName)
	ocr.buildProfileBottom(bottom)
}

func (ocr *OCR) Process() {
	ocr.ProcessInnovator()
	ocr.ProcessTop()
	ocr.ProcessBottom()
}

func (ocr *OCR) CleanUp() {
	os.Remove(conf.Config.Cache + "/" + ocr.tmpName + "_top.png")
	os.Remove(conf.Config.Cache + "/" + ocr.tmpName + "_bottom.png")
}

func (ocr *OCR) buildProfileTop(top []byte) {
	p := &ocr.Profile

	fmt.Printf("%s\n", top)

	p.Nick = matchString(top, "([a-zA-Z0-9]+)[^\n]*\\s*[^\n]*LVL")
	p.Level = int(genMatchNum(top, "LVL #"))
	p.AP = genMatchNum(top, "# AP")
}

func (ocr *OCR) buildProfileBottom(bottom []byte) {
	p := &ocr.Profile

	fmt.Printf("%s\n", bottom)

	// Discovery
	p.UniquePortalsVisited = genMatchNum(bottom, "Unique Portals Visited #")
	p.PortalsDiscovered = genMatchNum(bottom, "Portals Discovered #")
	p.XMCollected = genMatchNum(bottom, "XM Collected # XM")

	// Health
	p.DistanceWalked = genMatchNum(bottom, "Distance Walked # km")

	// Building
	p.ResonatorsDeployed = genMatchNum(bottom, "Resonators Deployed #")
	p.LinksCreated = genMatchNum(bottom, "Links Created #")
	p.ControlFieldsCreated = genMatchNum(bottom, "Control Fields Created #")
	p.MindUnitsCaptured = genMatchNum(bottom, "Mind Units Captured # MUs")
	p.LongestLinkEverCreated = genMatchNum(bottom, "Longest Link Ever Created # km")
	p.LargestControlField = genMatchNum(bottom, "Largest Control Field # MUs")
	p.XMRecharged = genMatchNum(bottom, "XM Recharged # XM")
	p.PortalsCaptured = genMatchNum(bottom, "Portals Captured #")
	p.UniquePortalsCaptured = genMatchNum(bottom, "Unique Portals Captured #")
	p.ModsDeployed = genMatchNum(bottom, "Mods Deployed #")

	// Combat
	p.ResonatorsDestroyed = genMatchNum(bottom, "Resonators Destroyed #")
	p.PortalsNeutralized = genMatchNum(bottom, "Portals Neutralized #")
	p.EnemyLinksDestroyed = genMatchNum(bottom, "Enemy Links Destroyed #")
	p.EnemyControlFieldsDestroyed = genMatchNum(bottom, "Enemy Control Fields Destroyed #")

	// Defense
	p.MaxTimePortalHeld = genMatchNum(bottom, "Max Time Portal Held # days")
	p.MaxTimeLinkMaintained = genMatchNum(bottom, "Max Time Link Maintained # days")
	p.MaxLinkLengthXDays = genMatchNum(bottom, "Max Link Length x Days # km-days")
	p.MaxTimeFieldHeld = genMatchNum(bottom, "Max Time Field Held # days")
	p.LargestFieldMUsXDays = genMatchNum(bottom, "Largest Field MUs x Days # MU-days")

	// Missions
	p.UniqueMissionsCompleted = genMatchNum(bottom, "Unique Missions Completed #")

	// Resource Gathering
	p.Hacks = genMatchNum(bottom, "Hacks #")
	p.GlyphHackPoints = genMatchNum(bottom, "Glyph Hack Points #")

	// Mentoring
	p.AgentsSuccessfullyRecruited = genMatchNum(bottom, "Agents Successfully Recruited #")
}

func New(fileName string, override int) *OCR {
	ocr := new(OCR)
	ocr.Filename = fileName
	ocr.OCRProfile = override

	return ocr
}
