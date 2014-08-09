package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
)

type configFormat struct {
	Database     string `json:"database"`
	CookieSecret string `json:"cookie-secret"`
}

var BadgeRank = []string{
	"Bronze",
	"Silver",
	"Gold",
	"Platinum",
	"Onyx",
}

var Badges = map[string][]uint{
	"Connector":       {50, 1000, 5000, 25000, 100000},
	"Builder":         {2000, 10000, 30000, 100000, 200000},
	"Explorer":        {100, 1000, 2000, 10000, 30000},
	"Guardian":        {3, 10, 20, 90, 150},
	"Hacker":          {2000, 10000, 30000, 100000, 200000},
	"Mind Controller": {100, 500, 2000, 10000, 40000},
	"Purifier":        {2000, 10000, 30000, 100000, 300000},
	"Seer":            {10, 50, 200, 500, 5000},
	"Liberator":       {200, 2000, 8000, 15000, 40000},
	"Pioneer":         {20, 200, 1000, 5000, 20000},
	"Recharger":       {100000, 1000000, 3000000, 10000000, 25000000},
}

type Badge struct {
	Rank     int
	Current  uint
	Required uint
}

type LevelRequirement struct {
	Level uint
	AP    uint

	Bronze   int
	Silver   int
	Gold     int
	Platinum int
	Onyx     int
}

var LevelRequirements = []LevelRequirement{
	{1, 0, 0, 0, 0, 0, 0},
	{2, 2500, 0, 0, 0, 0, 0},
	{3, 20000, 0, 0, 0, 0, 0},
	{4, 70000, 0, 0, 0, 0, 0},
	{5, 150000, 0, 0, 0, 0, 0},
	{6, 300000, 0, 0, 0, 0, 0},
	{7, 600000, 0, 0, 0, 0, 0},
	{8, 1200000, 0, 0, 0, 0, 0},
	{9, 2400000, 0, 4, 1, 0, 0},
	{10, 4000000, 0, 5, 2, 0, 0},
	{11, 6000000, 0, 6, 4, 0, 0},
	{12, 8400000, 0, 7, 6, 0, 0},
	{13, 12000000, 0, 0, 7, 1, 0},
	{14, 17000000, 0, 0, 7, 2, 0},
	{15, 24000000, 0, 0, 7, 3, 0},
	{16, 40000000, 0, 0, 7, 4, 2},
}

type Profile struct {
	Nick  string
	Level uint
	AP    uint

	NextLevel LevelRequirement

	Badges struct {
		Connector      Badge
		Builder        Badge
		Explorer       Badge
		Guardian       Badge
		Hacker         Badge
		MindController Badge
		Purifier       Badge
		Seer           Badge
		Liberator      Badge
		Pioneer        Badge
		Recharger      Badge
	}

	Bronze   int
	Silver   int
	Gold     int
	Platinum int
	Onyx     int

	UniquePortalsVisited uint
	PortalsDiscovered    uint
	XMCollected          uint

	Hacks                  uint
	ResonatorsDeployed     uint
	LinksCreated           uint
	ControlFieldsCreated   uint
	MindUnitsCaptured      uint
	LongestLinkEverCreated uint
	LargestControlField    uint
	XMRecharged            uint
	PortalsCaptured        uint
	UniquePortalsCaptured  uint

	ResonatorsDestroyed         uint
	PortalsNeutralized          uint
	EnemyLinksDestroyed         uint
	EnemyControlFieldsDestroyed uint

	DistanceWalked uint

	MaxTimePortalHeld     uint
	MaxTimeLinkMaintained uint
	MaxLinkLengthXDays    uint
	MaxTimeFieldHeld      uint

	LargestFieldMUsXDays uint
}

func findLevel(p *Profile) {
	for i := len(LevelRequirements) - 1; i >= 0; i-- {
		lr := LevelRequirements[i]
		if p.AP >= lr.AP &&
			p.Bronze >= lr.Bronze &&
			p.Silver >= lr.Silver &&
			p.Gold >= lr.Gold &&
			p.Platinum >= lr.Platinum &&
			p.Onyx >= p.Onyx {
			break
		}
	}
}

func nextLevel(p *Profile) {
	if p.Level < 16 {
		p.NextLevel = LevelRequirements[p.Level]
	}
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

func incBadgeRank(p *Profile, b *Badge, current uint, reqs []uint) {
	for i := 0; i < len(reqs); i++ {
		req := reqs[i]
		if current >= req {
			switch i {
			case 0:
				p.Bronze++
			case 1:
				p.Silver++
			case 2:
				p.Gold++
			case 3:
				p.Platinum++
			case 4:
				p.Onyx++
			}
		} else {
			b.Rank = i
			b.Current = current
			b.Required = req

			break
		}
	}
}

func countBadges(p *Profile) {
	for k, v := range Badges {
		switch k {
		case "Connector":
			incBadgeRank(p, &p.Badges.Connector, p.LinksCreated, v)
		case "Builder":
			incBadgeRank(p, &p.Badges.Builder, p.ResonatorsDeployed, v)
		case "Explorer":
			incBadgeRank(p, &p.Badges.Explorer, p.UniquePortalsVisited, v)
		case "Guardian":
			incBadgeRank(p, &p.Badges.Guardian, p.MaxTimePortalHeld, v)
		case "Hacker":
			incBadgeRank(p, &p.Badges.Hacker, p.Hacks, v)
		case "Mind Controller":
			incBadgeRank(p, &p.Badges.MindController, p.ControlFieldsCreated, v)
		case "Purifier":
			incBadgeRank(p, &p.Badges.Purifier, p.ResonatorsDestroyed, v)
		case "Seer":
			incBadgeRank(p, &p.Badges.Seer, p.PortalsDiscovered, v)
		case "Liberator":
			incBadgeRank(p, &p.Badges.Liberator, p.PortalsCaptured, v)
		case "Pioneer":
			incBadgeRank(p, &p.Badges.Pioneer, p.UniquePortalsCaptured, v)
		case "Recharger":
			incBadgeRank(p, &p.Badges.Recharger, p.XMRecharged, v)
		}
	}
}

func buildProfile(res string) Profile {
	var digit = `([0-9l,]+)`
	var p Profile

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

func main() {
	var config configFormat

	configBody, _ := ioutil.ReadFile("config.json")
	err := json.Unmarshal(configBody, &config)
	if err != nil {
		log.Fatal("Config error: %s\n", err)
	}


	var db, _ = sql.Open("mysql", config.Database)
	defer db.Close()

	var args = []string{
		"-psm",
		"4",
		os.Args[1],
	}

	file, _ := generateTmpFile()
	args = append(args, file)

	cmd := exec.Command("tesseract", args...)
	cmd.Run()

	fpath := file + ".txt"
	out, _ := os.OpenFile(fpath, 1, 1)
	buffer, _ := ioutil.ReadFile(out.Name())
	res := string(buffer)

	log.Println(res)

	os.Remove(out.Name())

	p := buildProfile(res)

	countBadges(&p)
	nextLevel(&p)

	result, err := db.Exec(`
		INSERT INTO tiers_profiles (user_id, timestamp, agent, level, ap, unique_portals_visited, portals_discovered,
		xm_collected, hacks, resonators_deployed, links_created, control_fields_created, mind_units_captured,
		longest_link_ever_created, largest_control_field, xm_recharged, portals_captured, unique_portals_captured,
		resonators_destroyed, portals_neutralized, enemy_links_destroyed, enemy_control_fields_destroyed,
		distance_walked, max_time_portal_held, max_time_link_maintained, max_link_length_x_days, max_time_field_held,
		largest_field_mus_x_days)
		VALUES(1, UNIX_TIMESTAMP(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		p.Nick, p.Level, p.AP,
		p.UniquePortalsVisited, p.PortalsDiscovered, p.XMCollected,
		p.Hacks, p.ResonatorsDeployed, p.LinksCreated, p.ControlFieldsCreated, p.MindUnitsCaptured, p.LongestLinkEverCreated,
		p.LargestControlField, p.XMRecharged, p.PortalsCaptured, p.UniquePortalsCaptured,
		p.ResonatorsDestroyed, p.PortalsNeutralized, p.EnemyLinksDestroyed, p.EnemyControlFieldsDestroyed,
		p.DistanceWalked,
		p.MaxTimePortalHeld, p.MaxTimeLinkMaintained, p.MaxLinkLengthXDays, p.MaxTimeFieldHeld,
		p.LargestFieldMUsXDays,
	)

	log.Println(result, err)

	log.Printf("%+v\n", p)
}
