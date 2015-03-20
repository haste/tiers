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
	Nick   []string `json:"nick"`
	Level  []string `json:"level"`
	AP     []string `json:"ap"`
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

func max(x, y int64) int64 {
	if x > y {
		return x
	}

	return y
}

func sanitizeNum(input []byte) int64 {
	input = regexp.MustCompile(`[lL|\]IitJ]`).ReplaceAll(input, []byte("1"))
	input = regexp.MustCompile(`[DOo]`).ReplaceAll(input, []byte("0"))
	input = regexp.MustCompile(`[Ss]`).ReplaceAll(input, []byte("5"))

	n := string(input)
	n = strings.Replace(n, "Q", "9", -1)
	n = strings.Replace(n, "B", "8", -1)
	n = strings.Replace(n, "Z", "7", -1)
	n = strings.Replace(n, "G", "6", -1)
	n = strings.Replace(n, "A", "4", -1)
	n = strings.Replace(n, "n", "11", -1)
	n = strings.Replace(n, "H", "11", -1)
	n = strings.Replace(n, ",", "", -1)
	n = strings.Replace(n, " ", "", -1)

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

func genMatch(s string) string {
	s = regexp.MustCompile(`[Ss]`).ReplaceAllLiteralString(s, "[Ss5]+")
	s = regexp.MustCompile(`[aeCcOoUu]`).ReplaceAllLiteralString(s, "[aecE8B9Cc0OoUu]+")
	s = regexp.MustCompile(`[Pp]`).ReplaceAllLiteralString(s, "[Pp]+")
	s = regexp.MustCompile(`[ltIfi]`).ReplaceAllLiteralString(s, "[l|1tIfi]+")
	s = regexp.MustCompile(`\s+`).ReplaceAllLiteralString(s, `\s*`)

	s = strings.Replace(s, "V", "[Vvu]", -1)
	s = strings.Replace(s, `D`, "[D00]+", -1)
	s = strings.Replace(s, `r`, "[rt]+", -1)
	s = strings.Replace(s, `Hn`, "im", -1)
	s = strings.Replace(s, `-`, ".", -1)
	s = strings.Replace(s, `#`, `([0-9QLIiltJBZADOonHSsG|,\] ]+)`, -1)

	return s
}

func genMatchNum(res []byte, s string) int64 {
	return matchNum(res, genMatch(s))
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
	case "nick":
		args = append(args, profile.Nick...)
	case "level":
		args = append(args, profile.Level...)
	case "ap":
		args = append(args, profile.AP...)
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

func (ocr *OCR) tesseract(fileName, kind string) []byte {
	args := []string{"-psm"}

	if kind == "nick" {
		args = append(args, "8")
	} else {
		args = append(args, "4")
	}

	args = append(args, []string{
		"-l", "eng",
		fileName,
		"stdout",
		"ingress",
	}...)

	tesseract := exec.Command(conf.Config.TesseractBin, args...)
	res, err := tesseract.Output()
	if err != nil {
		log.Fatal("tesseract ", err)
	}

	return res
}

func (ocr *OCR) replaces(in []byte) []byte {
	in = regexp.MustCompile(genMatch(`Enenwljnks`)).ReplaceAllLiteral(in, []byte("Enemy Links"))
	in = regexp.MustCompile(genMatch(`Enenlenks`)).ReplaceAllLiteral(in, []byte("Enemy Links"))
	in = regexp.MustCompile(genMatch(`EnmnyLmks`)).ReplaceAllLiteral(in, []byte("Enemy Links"))

	in = regexp.MustCompile(genMatch(`Resonahus`)).ReplaceAllLiteral(in, []byte("Resonators"))
	in = regexp.MustCompile(genMatch(`Resonauus`)).ReplaceAllLiteral(in, []byte("Resonators"))
	in = regexp.MustCompile(genMatch(`ResonauHs`)).ReplaceAllLiteral(in, []byte("Resonators"))

	in = regexp.MustCompile(genMatch(`MalemeLlnk`)).ReplaceAllLiteral(in, []byte("Max Time Link"))
	in = regexp.MustCompile(genMatch(`Maleme`)).ReplaceAllLiteral(in, []byte("Max Time"))

	// MaxThneLuutMaunamed
	in = regexp.MustCompile(genMatch(`MaxThneLuutMaunamed`)).ReplaceAllLiteral(in, []byte("Max Time Link Maintained"))
	// MaxThneLnntMannamed
	in = regexp.MustCompile(genMatch(`MaxThneLnntMannamed`)).ReplaceAllLiteral(in, []byte("Max Time Link Maintained"))
	// MaXTuneLhnrMaunmned
	in = regexp.MustCompile(genMatch(`MaXTuneLhnrMaunmned`)).ReplaceAllLiteral(in, []byte("Max Time Link Maintained"))

	// MaxTHnekajHem
	// MaxTHnekaiHem
	in = regexp.MustCompile(`M[aecE8B9Cc0]xT[Hh]n[aecE8B9Cc0]k[aecE8B9Cc0][ij]H[aecE8B9Cc0][mM]`).ReplaceAllLiteral(in, []byte("Max Time Field Held"))
	// MaxTHnekamed
	in = regexp.MustCompile(`M[aecE8B9Cc0]xT[Hh]n[aecE8B9Cc0]k[aecE8B9Cc0]m[aecE8B9Cc0]d`).ReplaceAllLiteral(in, []byte("Max Time Field Held"))

	// MaxTHnePonalHdd
	// MaxTHnePonalHem
	// MaxTnnePonalHem
	// MaxTHnePonaled
	// MaxTunePodalHdd
	in = regexp.MustCompile(`M[aecE8B9Cc0]xT[HhnUu]*[aecE8B9Cc0][Pp]o[Nnd]*[aecE8B9Cc0]l[HhnaecE8B9Cc0d]*[mMd]*`).ReplaceAllLiteral(in, []byte("Max Time Portal Held"))

	// 131th Hack Points
	// 131th HaCk Points
	in = regexp.MustCompile(`131th\s*H[aecE8B9Cc0][aecE8B9Cc0]k\s*[Pp]oin[l|1tI][Ss5]`).ReplaceAllLiteral(in, []byte("Glyph Hack Points"))

	// MamekLengH1xDays
	// MamekLengU1xDays
	// MamekLengN1xDays
	in = regexp.MustCompile(`M[aecE8B9Cc0]m[aecE8B9Cc0]kL[aecE8B9Cc0]ng[HhUuNn]1xD[aecE8B9Cc0]y[Ss5]`).ReplaceAllLiteral(in, []byte("Max Link Length x Days"))

	// HBCKS
	in = regexp.MustCompile(`[Hh][AaB][Cc][Kk][5Ss]`).ReplaceAllLiteral(in, []byte("Hacks"))

	in = bytes.Replace(in, []byte("rn"), []byte("m"), -1)

	in = bytes.Replace(in, []byte("081310de"), []byte("Deployed"), -1)

	in = bytes.Replace(in, []byte("Desnoyed"), []byte("Destroyed"), -1)
	in = bytes.Replace(in, []byte("Deonyed"), []byte("Destroyed"), -1)
	in = bytes.Replace(in, []byte("08500de"), []byte("Destroyed"), -1)
	in = bytes.Replace(in, []byte("Deshoyed"), []byte("Destroyed"), -1)
	in = bytes.Replace(in, []byte("DesUOyed"), []byte("Destroyed"), -1)

	in = bytes.Replace(in, []byte("M08"), []byte("MUs"), -1)

	in = bytes.Replace(in, []byte("knbdays"), []byte("km-days"), -1)
	in = bytes.Replace(in, []byte("knrdays"), []byte("km-days"), -1)
	in = bytes.Replace(in, []byte("kanays"), []byte("km-days"), -1)

	return in
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

func (ocr *OCR) ProcessNick() {
	fileName := conf.Config.Cache + "/" + ocr.tmpName + "_nick.png"
	ocr.mogrify("nick")
	nick := ocr.tesseract(fileName, "nick")
	ocr.buildProfileNick(nick)
}

func (ocr *OCR) ProcessLevel() {
	fileName := conf.Config.Cache + "/" + ocr.tmpName + "_level.png"
	ocr.mogrify("level")
	level := ocr.tesseract(fileName, "level")
	ocr.buildProfileLevel(level)
}

func (ocr *OCR) ProcessAP() {
	fileName := conf.Config.Cache + "/" + ocr.tmpName + "_ap.png"
	ocr.mogrify("ap")
	ap := ocr.tesseract(fileName, "ap")
	ocr.buildProfileAP(ap)
}

func (ocr *OCR) ProcessTop() {
	ocr.ProcessNick()
	ocr.ProcessLevel()
	ocr.ProcessAP()
}

func (ocr *OCR) ProcessBottom() {
	fileName := conf.Config.Cache + "/" + ocr.tmpName + "_bottom.png"
	ocr.mogrify("bottom")
	bottom := ocr.tesseract(fileName, "bottom")
	bottom = ocr.replaces(bottom)
	ocr.buildProfileBottom(bottom)
}

func (ocr *OCR) Process() {
	ocr.ProcessInnovator()
	ocr.ProcessTop()
	ocr.ProcessBottom()
}

func (ocr *OCR) CleanUp() {
	os.Remove(conf.Config.Cache + "/" + ocr.tmpName + "_nick.png")
	os.Remove(conf.Config.Cache + "/" + ocr.tmpName + "_level.png")
	os.Remove(conf.Config.Cache + "/" + ocr.tmpName + "_ap.png")
	os.Remove(conf.Config.Cache + "/" + ocr.tmpName + "_bottom.png")
}

func (ocr *OCR) buildProfileNick(nick []byte) {
	p := &ocr.Profile

	p.Nick = matchString(nick, "([a-zA-Z0-9]+)")
}

func (ocr *OCR) buildProfileLevel(level []byte) {
	p := &ocr.Profile

	p.Level = int(genMatchNum(level, "L V L #"))
}

func (ocr *OCR) buildProfileAP(ap []byte) {
	p := &ocr.Profile

	p.AP = genMatchNum(ap, "# A P")
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
	p.ConsecutiveDaysHacking = max(
		genMatchNum(bottom, "Consecutive Days Hacking # days"),
		genMatchNum(bottom, "Longest Hacking Streak # days"),
	)

	// Mentoring
	p.AgentsSuccessfullyRecruited = genMatchNum(bottom, "Agents Successfully Recruited #")
}

func New(fileName string, override int) *OCR {
	ocr := new(OCR)
	ocr.Filename = fileName
	ocr.OCRProfile = override

	return ocr
}
