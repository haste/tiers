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

type convertProfile struct {
	Top    []string `json:"top"`
	Bottom []string `json:"bottom"`
}

var convertOptions = map[int]convertProfile{}

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
				data  convertProfile
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

			convertOptions[width] = data
		}
	}
}

func sanitizeNum(input []byte) int64 {
	input = regexp.MustCompile(`[lL|\]JI]`).ReplaceAll(input, []byte("1"))
	input = regexp.MustCompile(`[Oo]`).ReplaceAll(input, []byte("0"))

	n := string(input)
	n = strings.Replace(n, "B", "8", -1)
	n = strings.Replace(n, "n", "11", -1)
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
	s = regexp.MustCompile(`[Oo]`).ReplaceAllLiteralString(s, "[0Oo]")
	s = regexp.MustCompile(`[Cc]`).ReplaceAllLiteralString(s, "[Cc]")
	s = regexp.MustCompile(`[ltI]`).ReplaceAllLiteralString(s, "[l|1tI]")
	s = regexp.MustCompile(`\s+`).ReplaceAllLiteralString(s, `\s*`)

	s = strings.Replace(s, `Hn`, "im", -1)
	s = strings.Replace(s, `-`, ".", -1)
	s = strings.Replace(s, `#`, `([0-9LIlJBOon|,\]]+)`, -1)

	return matchNum(res, s)
}

func buildProfile(top, bottom []byte) profile.Profile {
	var p profile.Profile

	p.Nick = matchString(top, "([a-zA-Z0-9]+)[^\n]*\\s*[^\n]*LVL")
	p.Level = int(genMatchNum(top, "LVL #"))
	p.AP = genMatchNum(top, "# AP")

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

	return p
}

func handleInnovator(p *profile.Profile, data innovator) {
	if data.Good > 0 && data.Rank >= 0 {
		p.InnovatorLevel = profile.BadgeRanks["Innovator"][data.Rank]
	}
}

func getConvertArgs(in, out, part string, width int) []string {
	var (
		profile     convertProfile
		convertArgs = []string{in}
	)

	if opt, ok := convertOptions[width]; ok {
		profile = opt
	} else {
		profile = convertOptions[0]
	}

	switch part {
	case "top":
		convertArgs = append(convertArgs, profile.Top...)
	case "bottom":
		convertArgs = append(convertArgs, profile.Bottom...)
	}

	return append(convertArgs, out)
}

func runOpenCV(in, out string) []byte {
	cv := exec.Command(conf.Config.PythonBin, []string{
		conf.Config.UtilsDir + "innovator-crop/crop.py",
		conf.Config.Cache,
		in,
		out,
	}...)

	res, err := cv.Output()
	if err != nil {
		log.Fatal("cv ", err)
	}

	return res
}

func runConvert(in, out, part string, width int) {
	convertArgs := getConvertArgs(in, out, part, width)
	convert := exec.Command(conf.Config.ConvertBin, convertArgs...)

	err := convert.Run()
	if err != nil {
		log.Fatal("convert ", err)
	}
}

func runTesseract(in string) []byte {
	tesseract := exec.Command(conf.Config.TesseractBin, []string{
		"-psm",
		"4",
		"-l", "eng",
		in,
		"stdout",
		"ingress",
	}...)

	res, err := tesseract.Output()
	if err != nil {
		log.Fatal("tesseract ", err)
	}

	return res
}

func runOCR(fileName string) profile.Profile {
	tmpId := id.New()
	tmpFormat := fmt.Sprintf("%s/%s_%%s_%%s.png", conf.Config.Cache, tmpId)

	cvTop := fmt.Sprintf(tmpFormat, "cv", "top")
	cvBottom := fmt.Sprintf(tmpFormat, "cv", "bottom")
	tmpTop := fmt.Sprintf(tmpFormat, "tmp", "top")
	tmpBottom := fmt.Sprintf(tmpFormat, "tmp", "bottom")

	res := runOpenCV(fileName, tmpId)

	var innovator innovator
	decoder := json.NewDecoder(bytes.NewReader(res))
	err := decoder.Decode(&innovator)

	reader, err := os.Open(cvTop)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	width := m.Bounds().Dx()
	runConvert(cvTop, tmpTop, "top", width)
	runConvert(cvBottom, tmpBottom, "bottom", width)

	top := runTesseract(tmpTop)
	bottom := runTesseract(tmpBottom)

	p := buildProfile(top, bottom)

	handleInnovator(&p, innovator)

	os.Remove(cvTop)
	os.Remove(cvBottom)
	os.Remove(tmpTop)
	os.Remove(tmpBottom)

	return p
}

// XXX: Should probably return an error as well
func OCR(fileName string) profile.Profile {
	return runOCR(fileName)
}
