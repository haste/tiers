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

	"tiers/conf"
	"tiers/profile"
)

type innovator struct {
	Rank  int `json:"rank"`
	Good  int `json:"good"`
	Total int `json:"total"`
}

func sanitizeNum(input []byte) int64 {
	input = regexp.MustCompile(`[lL|\]JI]`).ReplaceAll(input, []byte("1"))

	n := string(input)
	n = strings.Replace(n, "o", "0", -1)
	n = strings.Replace(n, "B", "8", -1)
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
	s = regexp.MustCompile(`[ae]`).ReplaceAllLiteralString(s, "[ae]")
	s = regexp.MustCompile(`[Pp]`).ReplaceAllLiteralString(s, "[Pp]")
	s = regexp.MustCompile(`[Oo]`).ReplaceAllLiteralString(s, "[0Oo]")
	s = regexp.MustCompile(`\s+`).ReplaceAllLiteralString(s, `\s*`)

	s = strings.Replace(s, `-`, ".", -1)
	s = strings.Replace(s, `#`, `([0-9LIl|J,B\]]+)`, -1)

	return matchNum(res, s)
}

func buildProfile(res []byte) profile.Profile {
	var p profile.Profile

	p.Nick = matchString(res, "([a-zA-Z0-9]+)[^\n]*\\s*[^\n]*LVL")
	p.Level = int(genMatchNum(res, "LVL #"))
	p.AP = genMatchNum(res, "# AP")

	p.UniquePortalsVisited = genMatchNum(res, "Unique Portals Visited #")
	p.PortalsDiscovered = genMatchNum(res, "Portals Discovered #")
	p.XMCollected = genMatchNum(res, "XM Collected # XM")

	p.Hacks = genMatchNum(res, "Hacks #")
	p.ResonatorsDeployed = genMatchNum(res, "Resonators Deployed #")
	p.LinksCreated = genMatchNum(res, "Links Created #")
	p.ControlFieldsCreated = genMatchNum(res, "Control Fields Created #")
	p.MindUnitsCaptured = genMatchNum(res, "Mind Units Captured # MUs")
	p.LongestLinkEverCreated = genMatchNum(res, "Longest Link Ever Created # km")
	p.LargestControlField = genMatchNum(res, "Largest Control Field # MUs")
	p.XMRecharged = genMatchNum(res, "XM Recharged # XM")
	p.PortalsCaptured = genMatchNum(res, "Portals Captured #")
	p.UniquePortalsCaptured = genMatchNum(res, "Unique Portals Captured #")

	p.ResonatorsDestroyed = genMatchNum(res, "Resonators Destroyed #")
	p.PortalsNeutralized = genMatchNum(res, "Portals Neutralized #")
	p.EnemyLinksDestroyed = genMatchNum(res, "Enemy Links Destroyed #")
	p.EnemyControlFieldsDestroyed = genMatchNum(res, "Enemy Control Fields Destroyed #")

	p.DistanceWalked = genMatchNum(res, "Distance Walked # km")

	p.MaxTimePortalHeld = genMatchNum(res, "Max Time Portal Held # days")
	p.MaxTimeLinkMaintained = genMatchNum(res, "Max Time Link Maintained # days")
	p.MaxLinkLengthXDays = genMatchNum(res, "Max Link Length x Days # km-days")
	p.MaxTimeFieldHeld = genMatchNum(res, "Max Time Field Held # days")
	p.LargestFieldMUsXDays = genMatchNum(res, "Largest Field MUs x Days # MU-days")

	p.UniqueMissionsCompleted = genMatchNum(res, "Unique Missions Completed #")

	return p
}

func handleInnovator(p *profile.Profile, data innovator) {
	if data.Good > 0 {
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
		log.Fatal("cv", err)
	}

	var innovator innovator
	decoder := json.NewDecoder(bytes.NewReader(res))
	err = decoder.Decode(&innovator)

	convert := exec.Command(conf.Config.ConvertBin, []string{
		cvFile,
		"-resize",
		"170%",
		"-level",
		"50%",
		"-colorspace",
		"gray",
		"+dither",
		"-colors",
		"2",
		"-negate",
		tmpFile,
	}...)

	err = convert.Run()
	if err != nil {
		log.Fatal("convert", err)
	}

	tesseract := exec.Command(conf.Config.TesseractBin, []string{
		"-psm",
		"4",
		tmpFile,
		"stdout",
		"ingress",
	}...)

	res, err = tesseract.Output()
	if err != nil {
		log.Fatal("tesseract", err)
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
