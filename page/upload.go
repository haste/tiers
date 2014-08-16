package page

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"tiers/conf"
	"tiers/profile"
	"tiers/session"
	"time"
)

var templates *template.Template

func sanitizeNum(n string) uint {
	n = strings.Replace(n, "l", "1", -1)
	n = strings.Replace(n, "o", "0", -1)
	n = strings.Replace(n, ",", "", -1)

	un, _ := strconv.ParseUint(n, 10, 0)
	return uint(un)
}

func matchString(res, pattern string) string {
	r := regexp.MustCompile(pattern)

	if r.MatchString(res) != true {
		return ""
	}

	return r.FindStringSubmatch(res)[1]
}

func matchNum(res, pattern string) uint {
	r := regexp.MustCompile(pattern)

	if r.MatchString(res) != true {
		return 0
	}

	return sanitizeNum(r.FindStringSubmatch(res)[1])
}

func buildProfile(res string) profile.Profile {
	var digit = `([0-9l,]+)`
	var p profile.Profile

	p.Nick = matchString(res, "([a-zA-Z0-9]+)[^\n]+\nLVL")
	p.Level = matchNum(res, "LVL ?"+digit)
	p.AP = matchNum(res, digit+" AP")

	p.UniquePortalsVisited = matchNum(res, "Unique Portals Visited "+digit)
	p.PortalsDiscovered = matchNum(res, "Portals Discovered "+digit)
	p.XMCollected = matchNum(res, "XM Collected "+digit+" XM")

	p.Hacks = matchNum(res, "Hacks "+digit)
	p.ResonatorsDeployed = matchNum(res, "Resonators Deployed "+digit)
	p.LinksCreated = matchNum(res, "Links Created "+digit)
	p.ControlFieldsCreated = matchNum(res, "Control Fields Created "+digit)
	p.MindUnitsCaptured = matchNum(res, "Mind Units Captured "+digit)
	p.LongestLinkEverCreated = matchNum(res, "Longest Link Ever Created "+digit+" km")
	p.LargestControlField = matchNum(res, "Largest Control Field "+digit+" MUs")
	p.XMRecharged = matchNum(res, "XM Recharged "+digit+" XM")
	p.PortalsCaptured = matchNum(res, "Portals Captured "+digit)
	p.UniquePortalsCaptured = matchNum(res, "Unique Portals Captured "+digit)

	p.ResonatorsDestroyed = matchNum(res, "Resonators Destroyed "+digit)
	p.PortalsNeutralized = matchNum(res, "Portals Neutralized "+digit)
	p.EnemyLinksDestroyed = matchNum(res, "Enemy Links Destroyed "+digit)
	p.EnemyControlFieldsDestroyed = matchNum(res, "Enemy Control Fields Destroyed "+digit)

	p.DistanceWalked = matchNum(res, "Distance Walked "+digit)

	p.MaxTimePortalHeld = matchNum(res, "Max Time Portal Held "+digit+" days")
	p.MaxTimeLinkMaintained = matchNum(res, "Max Time Link Maintained "+digit+" days")
	p.MaxLinkLengthXDays = matchNum(res, "Max Link Length x Days "+digit+" km-days")
	p.MaxTimeFieldHeld = matchNum(res, "Max Time Field Held "+digit+" days")

	p.LargestFieldMUsXDays = matchNum(res, "Largest Field MUs x Days "+digit+" MU-days")

	return p
}

func UploadViewHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if templates, err = template.New("").ParseFiles(
		"templates/header.html",
		"templates/footer.html",
		"templates/nav.html",
		"templates/upload.html",
	); err != nil {
		log.Fatal(err)
	}

	session, _ := session.Get(r, "tiers")
	userid, ok := session.Values["user"]

	if !ok {
		return
	}

	templates.ExecuteTemplate(w, "upload", userid)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var userid = 1

	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// XXX: Handle errors.
	var db, _ = sql.Open("mysql", conf.Config.Database)
	defer db.Close()

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		//if part.FileName() is empty, skip this iteration.
		if part.FileName() == "" {
			continue
		}

		// profile_20140815_135412_0.png
		r := regexp.MustCompile("^(ingress|profile)_(\\d+)_(\\d+)_\\d+\\.png$")
		if r.MatchString(part.FileName()) != true {
			// XXX: Should probably handle this..
			continue
		}

		m := r.FindStringSubmatch(part.FileName())
		t, _ := time.ParseInLocation("20060102150405", m[2]+m[3], time.Local)

		var fileName = fmt.Sprintf("%d_%s", userid, part.FileName())

		dst, err := os.Create(conf.Config.Cache + fileName)
		defer dst.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// XXX: Handle errors...
		db.Exec(`
		INSERT INTO tiers_queues(user_id, timestamp, file)
		VALUES(?, ?, ?)
		`, userid, t.Unix(), fileName)

	}
}
