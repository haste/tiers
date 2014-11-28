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
	n := string(input)
	n = strings.Replace(n, "l", "1", -1)
	n = strings.Replace(n, "L", "1", -1)
	n = strings.Replace(n, "|", "1", -1)
	n = strings.Replace(n, "]", "1", -1)
	n = strings.Replace(n, "J", "1", -1)
	n = strings.Replace(n, "I", "1", -1)
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

func buildProfile(res []byte) profile.Profile {
	var digit = `([0-9LIl|J,B\]]+)`
	var p profile.Profile

	p.Nick = matchString(res, "([a-zA-Z0-9]+)[^\n]*\\s*[^\n]*LVL")
	p.Level = int(matchNum(res, "LVL\\s*"+digit))
	p.AP = matchNum(res, digit+"\\s*A[Pp]")

	p.UniquePortalsVisited = matchNum(res, "Uniqu[ae]\\s*Port[ae]ls\\s*Visit[ae]d\\s*"+digit)
	p.PortalsDiscovered = matchNum(res, "Port[ae]ls\\s*Discov[ae]r[ae]d\\s*"+digit)
	p.XMCollected = matchNum(res, "XM\\s*Coll[ae]ct[ae]d\\s*"+digit+"\\s*XM")

	p.Hacks = matchNum(res, "H[ae]cks\\s*"+digit)
	p.ResonatorsDeployed = matchNum(res, "R[ae]son[ae]tors\\s*D[ae]ploy[ae]d\\s*"+digit)
	p.LinksCreated = matchNum(res, "Links\\s*Cr[ae][ae]t[ae]d\\s*"+digit)
	p.ControlFieldsCreated = matchNum(res, "Control\\s*Fi[ae]lds\\s*Cr[ae][ae]t[ae]d\\s*"+digit)
	p.MindUnitsCaptured = matchNum(res, "Mind\\s*Units\\s*C[ae]ptur[ae]d\\s*"+digit)
	p.LongestLinkEverCreated = matchNum(res, "Long[ae]st\\s*Link\\s*Ev[ae]r\\s*Cr[ae][ae]t[ae]d\\s*"+digit+"\\s*km")
	p.LargestControlField = matchNum(res, "L[ae]rg[ae]st\\s*Control\\s*Fi[ae]ld\\s*"+digit+"\\s*MUs")
	p.XMRecharged = matchNum(res, "XM\\s*R[ae]ch[ae]rg[ae]d "+digit+"\\s*XM")
	p.PortalsCaptured = matchNum(res, "Port[ae]ls\\s*C[ae]ptur[ae]d\\s*"+digit)
	p.UniquePortalsCaptured = matchNum(res, "Uniqu[ae]\\s*Port[ae]ls\\s*C[ae]ptur[ae]d\\s*"+digit)

	p.ResonatorsDestroyed = matchNum(res, "R[ae]son[ae]tors\\s*D[ae]stroy[ae]d\\s*"+digit)
	p.PortalsNeutralized = matchNum(res, "Port[ae]ls\\s*N[ae]utr[ae]liz[ae]d\\s*"+digit)
	p.EnemyLinksDestroyed = matchNum(res, "En[ae]my\\s*Links\\s*D[ae]stroy[ae]d\\s*"+digit)
	p.EnemyControlFieldsDestroyed = matchNum(res, "En[ae]my\\s*Control\\s*Fi[ae]lds\\s*D[ae]stroy[ae]d\\s*"+digit)

	p.DistanceWalked = matchNum(res, "Dist[ae]nc[ae]\\s*W[ae]lk[ae]d\\s*"+digit)

	p.MaxTimePortalHeld = matchNum(res, "M[ae]x\\s*Tim[ae]\\s*Port[ae]l\\s*H[ae]ld\\s*"+digit+"\\s*d[ae]ys")
	p.MaxTimeLinkMaintained = matchNum(res, "M[ae]x\\s*Tim[ae]\\s*Link\\s*M[ae]int[ae]in[ae]d\\s*"+digit+"\\s*d[ae]ys")
	p.MaxLinkLengthXDays = matchNum(res, "M[ae]x\\s*Link\\s*L[ae]ngth\\s*x\\s*D[ae]ys\\s*"+digit+"\\s*km.d[ae]ys")
	p.MaxTimeFieldHeld = matchNum(res, "M[ae]x\\s*Tim[ae]\\s*Fi[ae]ld\\s*H[ae]ld\\s*"+digit+"\\s*d[ae]ys")
	p.LargestFieldMUsXDays = matchNum(res, "L[ae]rg[ae]st\\s*Fi[ae]ld\\s*MUs\\s*x\\s*D[ae]ys\\s*"+digit+"\\s*MU.d[ae]ys")

	p.UniqueMissionsCompleted = matchNum(res, "Uniqu[ae]\\s*Missions\\s*Compl[ae]t[ae]d\\s*"+digit)

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

	cv := exec.Command(conf.Config.UtilsDir+"innovator-crop/crop.py", []string{
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
