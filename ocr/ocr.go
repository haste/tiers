package ocr

import (
	"bytes"
	"encoding/json"
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
	"tiers/profile"
)

type innovator struct {
	Rank  int     `json:"rank"`
	Good  float32 `json:"good"`
	Total int     `json:"total"`
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
	s = regexp.MustCompile(`\s+`).ReplaceAllLiteralString(s, `\s*`)

	s = strings.Replace(s, `l`, "[l|1]", -1)
	s = strings.Replace(s, `-`, ".", -1)
	s = strings.Replace(s, `#`, `([0-9LIlJBOon|,\]]+)`, -1)

	return matchNum(res, s)
}

func buildProfile(res []byte) profile.Profile {
	var p profile.Profile

	p.Nick = matchString(res, "([a-zA-Z0-9]+)[^\n]*\\s*[^\n]*LVL")
	p.Level = int(genMatchNum(res, "LVL #"))
	p.AP = genMatchNum(res, "# AP")

	// Discovery
	p.UniquePortalsVisited = genMatchNum(res, "Unique Portals Visited #")
	p.PortalsDiscovered = genMatchNum(res, "Portals Discovered #")
	p.XMCollected = genMatchNum(res, "XM Collected # XM")

	// Health
	p.DistanceWalked = genMatchNum(res, "Distance Walked # km")

	// Building
	p.ResonatorsDeployed = genMatchNum(res, "Resonators Deployed #")
	p.LinksCreated = genMatchNum(res, "Links Created #")
	p.ControlFieldsCreated = genMatchNum(res, "Control Fields Created #")
	p.MindUnitsCaptured = genMatchNum(res, "Mind Units Captured # MUs")
	p.LongestLinkEverCreated = genMatchNum(res, "Longest Link Ever Created # km")
	p.LargestControlField = genMatchNum(res, "Largest Control Field # MUs")
	p.XMRecharged = genMatchNum(res, "XM Recharged # XM")
	p.PortalsCaptured = genMatchNum(res, "Portals Captured #")
	p.UniquePortalsCaptured = genMatchNum(res, "Unique Portals Captured #")
	p.ModsDeployed = genMatchNum(res, "Mods Deployed #")

	// Combat
	p.ResonatorsDestroyed = genMatchNum(res, "Resonators Destroyed #")
	p.PortalsNeutralized = genMatchNum(res, "Portals Neutralized #")
	p.EnemyLinksDestroyed = genMatchNum(res, "Enemy Links Destroyed #")
	p.EnemyControlFieldsDestroyed = genMatchNum(res, "Enemy Control Fields Destroyed #")

	// Defense
	p.MaxTimePortalHeld = genMatchNum(res, "Max Time Portal Held # days")
	p.MaxTimeLinkMaintained = genMatchNum(res, "Max Time Link Maintained # days")
	p.MaxLinkLengthXDays = genMatchNum(res, "Max Link Length x Days # km-days")
	p.MaxTimeFieldHeld = genMatchNum(res, "Max Time Field Held # days")
	p.LargestFieldMUsXDays = genMatchNum(res, "Largest Field MUs x Days # MU-days")

	// Missions
	p.UniqueMissionsCompleted = genMatchNum(res, "Unique Missions Completed #")

	// Resource Gathering
	p.Hacks = genMatchNum(res, "Hacks #")

	return p
}

func handleInnovator(p *profile.Profile, data innovator) {
	if data.Good > 0 && data.Rank >= 0 {
		p.InnovatorLevel = profile.BadgeRanks["Innovator"][data.Rank]
	}
}

func runOCR(fileName string) profile.Profile {
	cvFile := conf.Config.Cache + "cv_" + fileName
	tmpFile := conf.Config.Cache + "tmp_" + fileName

	cv := exec.Command(conf.Config.PythonBin, []string{
		conf.Config.UtilsDir + "innovator-crop/crop.py",
		conf.Config.Cache,
		fileName,
	}...)

	res, err := cv.Output()
	if err != nil {
		log.Fatal("cv ", err)
	}

	var innovator innovator
	decoder := json.NewDecoder(bytes.NewReader(res))
	err = decoder.Decode(&innovator)

	reader, err := os.Open(cvFile)
	if err != nil {
		log.Fatal(err)
	}

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	width := m.Bounds().Dx()

	convertArgs := []string{cvFile}
	if width == 1200 {
		convertArgs = append(convertArgs,
			"-resize",
			"105%",
			"-level",
			"45%",
			"-colorspace",
			"gray",
			"+dither",
			"-colors",
			"2",
			"-negate",
		)
	} else if width == 1080 {
		convertArgs = append(convertArgs,
			"-resize",
			"100%",
			"-level",
			"25%",
			"-colorspace",
			"gray",
			"-blur",
			"1x65535",
			"-negate",
			"-sharpen",
			"1x65535",
		)
	} else if width == 768 || width == 720 {
		convertArgs = append(convertArgs,
			"-resize",
			"175%",
			"-level",
			"15%",
			"-colorspace",
			"gray",
			"+dither",
			"-colors",
			"2",
			"-negate",
		)
	} else if width == 480 {
		convertArgs = append(convertArgs,
			"-resize",
			"140%",
			"-level",
			"30%",
			"-colorspace",
			"gray",
			"+dither",
			"-colors",
			"2",
			"-negate",
		)
	} else {
		//} else if width == 640 {
		convertArgs = append(convertArgs,
			"-resize",
			"150%",
			"-level",
			"25%",
			"-colorspace",
			"gray",
			"+dither",
			"-colors",
			"2",
			"-negate",
		)
	}
	convertArgs = append(convertArgs, tmpFile)
	convert := exec.Command(conf.Config.ConvertBin, convertArgs...)

	err = convert.Run()
	if err != nil {
		log.Fatal("convert ", err)
	}

	tesseract := exec.Command(conf.Config.TesseractBin, []string{
		"-psm",
		"4",
		"-l", "eng",
		tmpFile,
		"stdout",
		"ingress",
	}...)

	res, err = tesseract.Output()
	if err != nil {
		log.Fatal("tesseract ", err)
	}

	os.Remove(cvFile)
	os.Remove(tmpFile)

	p := buildProfile(res)

	handleInnovator(&p, innovator)

	return p
}

// XXX: Should probably return an error as well
func OCR(fileName string) profile.Profile {
	return runOCR(fileName)
}
