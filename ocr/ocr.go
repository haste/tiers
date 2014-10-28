package ocr

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"tiers/conf"
	"tiers/profile"
)

func sanitizeNum(input []byte) uint {
	n := string(input)
	n = strings.Replace(n, "l", "1", -1)
	n = strings.Replace(n, "L", "1", -1)
	n = strings.Replace(n, "|", "1", -1)
	n = strings.Replace(n, "]", "1", -1)
	n = strings.Replace(n, "o", "0", -1)
	n = strings.Replace(n, "B", "8", -1)
	n = strings.Replace(n, ",", "", -1)

	un, _ := strconv.ParseUint(n, 10, 0)
	return uint(un)
}

func matchString(res []byte, pattern string) string {
	r := regexp.MustCompile(pattern)

	if r.Match(res) != true {
		return ""
	}

	return string(r.FindSubmatch(res)[1])
}

func matchNum(res []byte, pattern string) uint {
	r := regexp.MustCompile(pattern)

	if r.Match(res) != true {
		return 0
	}

	return sanitizeNum(r.FindSubmatch(res)[1])
}

func buildProfile(res []byte) profile.Profile {
	var digit = `([0-9Ll|,B\]]+)`
	var p profile.Profile

	// Remove the menu.
	res = res[bytes.Index(res, []byte("\n")):]

	p.Nick = matchString(res, "([a-zA-Z0-9]+)[^\n]*\n*[^\n]*LVL")
	p.Level = matchNum(res, "LVL\\s*"+digit)
	p.AP = matchNum(res, digit+"\\s*A[Pp]")

	p.UniquePortalsVisited = matchNum(res, "Unique\\s*Portals\\s*Visited\\s*"+digit)
	p.PortalsDiscovered = matchNum(res, "Portals\\s*Discovered\\s*"+digit)
	p.XMCollected = matchNum(res, "XM\\s*Collected\\s*"+digit+"\\s*XM")

	p.Hacks = matchNum(res, "Hacks\\s*"+digit)
	p.ResonatorsDeployed = matchNum(res, "Resonators\\s*Deployed\\s*"+digit)
	p.LinksCreated = matchNum(res, "Links\\s*Created\\s*"+digit)
	p.ControlFieldsCreated = matchNum(res, "Control\\s*Fields\\s*Created\\s*"+digit)
	p.MindUnitsCaptured = matchNum(res, "Mind\\s*Units\\s*Captured\\s*"+digit)
	p.LongestLinkEverCreated = matchNum(res, "Longest\\s*Link\\s*Ever\\s*Created\\s*"+digit+"\\s*km")
	p.LargestControlField = matchNum(res, "Largest\\s*Control\\s*Field\\s*"+digit+"\\s*MUs")
	p.XMRecharged = matchNum(res, "XM\\s*Recharged "+digit+"\\s*XM")
	p.PortalsCaptured = matchNum(res, "Portals\\s*Captured\\s*"+digit)
	p.UniquePortalsCaptured = matchNum(res, "Unique\\s*Portals\\s*Captured\\s*"+digit)

	p.ResonatorsDestroyed = matchNum(res, "Resonators\\s*Destroyed\\s*"+digit)
	p.PortalsNeutralized = matchNum(res, "Portals\\s*Neutralized\\s*"+digit)
	p.EnemyLinksDestroyed = matchNum(res, "Enemy\\s*Links\\s*Destroyed\\s*"+digit)
	p.EnemyControlFieldsDestroyed = matchNum(res, "Enemy\\s*Control\\s*Fields\\s*Destroyed\\s*"+digit)

	p.DistanceWalked = matchNum(res, "Distance\\s*Walked\\s*"+digit)

	p.MaxTimePortalHeld = matchNum(res, "Max\\s*Time\\s*Portal\\s*Held\\s*"+digit+"\\s*days")
	p.MaxTimeLinkMaintained = matchNum(res, "Max\\s*Time\\s*Link\\s*Maintained\\s*"+digit+"\\s*days")
	p.MaxLinkLengthXDays = matchNum(res, "Max\\s*Link\\s*Length\\s*x\\s*Days\\s*"+digit+"\\s*km.days")
	p.MaxTimeFieldHeld = matchNum(res, "Max\\s*Time\\s*Field\\s*Held\\s*"+digit+"\\s*days")

	p.LargestFieldMUsXDays = matchNum(res, "Largest\\s*Field\\s*MUs\\s*x\\s*Days\\s*"+digit+"\\s*MU.days")

	return p
}

func runOCR(fileName string) []byte {
	tmpFile := conf.Config.Cache + "tmp_" + fileName

	convert := exec.Command("convert", []string{
		conf.Config.Cache + fileName,
		"-resize",
		"175%",
		"-level",
		"65%",
		"-colorspace",
		"gray",
		"+dither",
		"-colors",
		"2",
		"-normalize",
		tmpFile,
	}...)

	err := convert.Run()
	if err != nil {
		log.Fatal(err)
	}

	tesseract := exec.Command("tesseract", []string{
		"-psm",
		"4",
		tmpFile,
		"stdout",
		"ingress",
	}...)

	res, err := tesseract.Output()
	if err != nil {
		log.Fatal(err)
	}

	os.Remove(tmpFile)

	return res
}

// XXX: Should probably return an error as well
func OCR(fileName string) profile.Profile {
	res := runOCR(fileName)
	return buildProfile(res)
}
