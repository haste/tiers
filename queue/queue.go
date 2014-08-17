package queue

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"tiers/conf"
	"tiers/profile"
	"time"
)

var Queue = make(chan bool, 1)

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

func generateTmpFile() (fname string, e error) {
	myTmpDir := "" // TODO: enable to choose optionally
	f, e := ioutil.TempFile(myTmpDir, "gosseract")
	if e != nil {
		return
	}
	fname = f.Name()
	return
}

func ProcessQueue() {
	for {
		<-Queue
		log.Println("Queue: Processing.")

		// XXX: Handle errors.
		var db, _ = sql.Open("mysql", conf.Config.Database)
		defer db.Close()

		// XXX: Handle errors.
		rows, err := db.Query(`
		SELECT id, user_id, timestamp, file, processed
		FROM tiers_queues
		WHERE processed = 0
		`)

		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var start = time.Now()

			var id, user_id, timestamp int
			var file string
			var processed bool

			if err := rows.Scan(&id, &user_id, &timestamp, &file, &processed); err != nil {
				log.Fatal(err)
			}

			var args = []string{
				"-psm",
				"4",
				conf.Config.Cache + file,
			}

			tmpFile, _ := generateTmpFile()
			args = append(args, tmpFile)

			cmd := exec.Command("tesseract", args...)
			cmd.Run()

			fpath := tmpFile + ".txt"
			out, _ := os.OpenFile(fpath, 1, 1)
			buffer, _ := ioutil.ReadFile(out.Name())
			res := string(buffer)

			p := buildProfile(res)

			// Handle errors
			_, err := db.Exec(`
			INSERT INTO tiers_profiles (user_id, timestamp, agent, level, ap, unique_portals_visited, portals_discovered,
			xm_collected, hacks, resonators_deployed, links_created, control_fields_created, mind_units_captured,
			longest_link_ever_created, largest_control_field, xm_recharged, portals_captured, unique_portals_captured,
			resonators_destroyed, portals_neutralized, enemy_links_destroyed, enemy_control_fields_destroyed,
			distance_walked, max_time_portal_held, max_time_link_maintained, max_link_length_x_days, max_time_field_held,
			largest_field_mus_x_days)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
			`,
				user_id, timestamp,
				p.Nick, p.Level, p.AP,
				p.UniquePortalsVisited, p.PortalsDiscovered, p.XMCollected,
				p.Hacks, p.ResonatorsDeployed, p.LinksCreated, p.ControlFieldsCreated, p.MindUnitsCaptured, p.LongestLinkEverCreated,
				p.LargestControlField, p.XMRecharged, p.PortalsCaptured, p.UniquePortalsCaptured,
				p.ResonatorsDestroyed, p.PortalsNeutralized, p.EnemyLinksDestroyed, p.EnemyControlFieldsDestroyed,
				p.DistanceWalked,
				p.MaxTimePortalHeld, p.MaxTimeLinkMaintained, p.MaxLinkLengthXDays, p.MaxTimeFieldHeld,
				p.LargestFieldMUsXDays,
			)

			if err != nil {
				log.Fatal(err)
			}

			// Handle errors
			_, err = db.Exec(`
			UPDATE tiers_queues
			SET processed = 1
			WHERE id = ?
			`, id)

			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Queue: Entry processed in %dms: %s L%d %dAP", time.Now().Sub(start).Nanoseconds()/1e6, p.Nick, p.Level, p.AP)
		}

		log.Println("Queue: Done.")
	}
}
