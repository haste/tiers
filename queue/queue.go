package queue

import (
	"database/sql"
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
	n = strings.Replace(n, "B", "8", -1)
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
	var digit = `([0-9l,B]+)`
	var p profile.Profile

	p.Nick = matchString(res, "([a-zA-Z0-9]+)[^\n]+\n*LVL")
	p.Level = matchNum(res, "LVL ?"+digit)
	p.AP = matchNum(res, digit+" AP")

	p.UniquePortalsVisited = matchNum(res, "Unique ?Portals ?Visited "+digit)
	p.PortalsDiscovered = matchNum(res, "Portals ?Discovered "+digit)
	p.XMCollected = matchNum(res, "XM ?Collected "+digit+" XM")

	p.Hacks = matchNum(res, "Hacks "+digit)
	p.ResonatorsDeployed = matchNum(res, "Resonators ?Deployed "+digit)
	p.LinksCreated = matchNum(res, "Links ?Created "+digit)
	p.ControlFieldsCreated = matchNum(res, "Control ?Fields ?Created "+digit)
	p.MindUnitsCaptured = matchNum(res, "Mind ?Units ?Captured "+digit)
	p.LongestLinkEverCreated = matchNum(res, "Longest ?Link ?Ever ?Created "+digit+" km")
	p.LargestControlField = matchNum(res, "Largest ?Control ?Field "+digit+" MUs")
	p.XMRecharged = matchNum(res, "XM ?Recharged "+digit+" XM")
	p.PortalsCaptured = matchNum(res, "Portals ?Captured "+digit)
	p.UniquePortalsCaptured = matchNum(res, "Unique ?Portals ?Captured "+digit)

	p.ResonatorsDestroyed = matchNum(res, "Resonators ?Destroyed "+digit)
	p.PortalsNeutralized = matchNum(res, "Portals ?Neutralized "+digit)
	p.EnemyLinksDestroyed = matchNum(res, "Enemy ?Links ?Destroyed "+digit)
	p.EnemyControlFieldsDestroyed = matchNum(res, "Enemy ?Control ?Fields ?Destroyed "+digit)

	p.DistanceWalked = matchNum(res, "Distance ?Walked "+digit)

	p.MaxTimePortalHeld = matchNum(res, "Max ?Time ?Portal ?Held "+digit+" days")
	p.MaxTimeLinkMaintained = matchNum(res, "Max ?Time ?Link ?Maintained "+digit+" days")
	p.MaxLinkLengthXDays = matchNum(res, "Max ?Link ?Length ?x ?Days "+digit+" km-days")
	p.MaxTimeFieldHeld = matchNum(res, "Max ?Time ?Field ?Held "+digit+" days")

	p.LargestFieldMUsXDays = matchNum(res, "Largest ?Field ?MUs ?x ?Days "+digit+" MU-days")

	return p
}

func ProcessQueue() {
	for {
		<-Queue
		log.Println("Queue: Processing.")

		// XXX: Handle errors.
		var db, _ = sql.Open("mysql", conf.Config.Database)

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
			var tmpFile string
			var processed bool

			if err := rows.Scan(&id, &user_id, &timestamp, &file, &processed); err != nil {
				log.Fatal(err)
			}

			tmpFile = conf.Config.Cache + "tmp_" + file

			convert := exec.Command("convert", []string{
				conf.Config.Cache + file,
				"-level",
				"40%",
				tmpFile,
			}...)

			err = convert.Run()
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

			p := buildProfile(string(res))

			// Handle errors
			_, err = db.Exec(`
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

			processTime := time.Now().Sub(start).Nanoseconds() / 1e6

			// Handle errors
			_, err = db.Exec(`
			UPDATE tiers_queues
			SET processed = 1,
			processtime = ?
			WHERE id = ?
			`, processTime, id)

			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Queue: Entry processed in %dms: %s L%d %dAP", processTime, p.Nick, p.Level, p.AP)
		}

		log.Println("Queue: Done.")

		rows.Close()
		db.Close()
	}
}
