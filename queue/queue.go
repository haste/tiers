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
	p.Level = matchNum(res, "LVL\\s*"+digit)
	p.AP = matchNum(res, digit+"\\s*AP")

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
