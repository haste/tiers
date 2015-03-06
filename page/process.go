package page

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"tiers/model"
	"tiers/ocr"
	"tiers/profile"
	"tiers/session"

	"github.com/gorilla/mux"
)

type userQueue struct {
	Id     int
	Latest int64
	Count  int64
}

type userQueueEntry struct {
	Id          int
	File        string
	OCRProfile  int
	Processed   int
	Processtime int
	HasProblem  bool
	Previous    int
	Profile     profile.Profile
}

func comparePrevious(a1, a2 profile.Profile) bool {
	v1 := reflect.ValueOf(a1)
	v2 := reflect.ValueOf(a2)

	for i, n := 0, v1.NumField(); i < n; i++ {
		t1f := reflect.TypeOf(a1).Field(i)
		if t1f.Name == "Id" {
			continue
		}

		v1f := v1.Field(i)
		v2f := v2.Field(i)
		switch v1f.Kind() {
		case reflect.String:
			if v1f.String() != v2f.String() {
				log.Printf("%s: %s vs %s", t1f.Name, v1f.String(), v2f.String())
				return true
			}
		case reflect.Int, reflect.Int64:
			if v1f.Int() > v2f.Int() {
				log.Printf("%s: %d vs %d", t1f.Name, v1f.Int(), v2f.Int())
				return true
			}
		}
	}

	return false
}

func rowToQueueProfile(rows *sql.Rows) userQueueEntry {
	var e userQueueEntry
	err := rows.Scan(
		&e.Id, &e.File, &e.OCRProfile, &e.Processed, &e.Processtime,

		&e.Profile.Id, &e.Profile.Timestamp, &e.Profile.Nick, &e.Profile.Level,
		&e.Profile.AP,

		&e.Profile.UniquePortalsVisited, &e.Profile.PortalsDiscovered,
		&e.Profile.XMCollected,

		&e.Profile.DistanceWalked,

		&e.Profile.ResonatorsDeployed, &e.Profile.LinksCreated,
		&e.Profile.ControlFieldsCreated, &e.Profile.MindUnitsCaptured,
		&e.Profile.LongestLinkEverCreated, &e.Profile.LargestControlField,
		&e.Profile.XMRecharged, &e.Profile.PortalsCaptured,
		&e.Profile.UniquePortalsCaptured, &e.Profile.ModsDeployed,

		&e.Profile.ResonatorsDestroyed, &e.Profile.PortalsNeutralized,
		&e.Profile.EnemyLinksDestroyed, &e.Profile.EnemyControlFieldsDestroyed,

		&e.Profile.MaxTimePortalHeld, &e.Profile.MaxTimeLinkMaintained,
		&e.Profile.MaxLinkLengthXDays, &e.Profile.MaxTimeFieldHeld,
		&e.Profile.LargestFieldMUsXDays,

		&e.Profile.UniqueMissionsCompleted,

		&e.Profile.Hacks, &e.Profile.GlyphHackPoints,
		&e.Profile.ConsecutiveDaysHacking,

		&e.Profile.AgentsSuccessfullyRecruited,

		&e.Profile.InnovatorLevel,
	)

	if err != nil {
		fmt.Println(err)
	}

	return e
}

func ProcessRunHandler(w http.ResponseWriter, r *http.Request) {
	faulty := mux.Vars(r)["faulty"]
	faultyInt, _ := strconv.Atoi(faulty)
	row := model.GetQueueWithProfileById(faultyInt)
	defer row.Close()

	owner := mux.Vars(r)["owner"]
	previous := mux.Vars(r)["previous"]

	row.Next()
	faultyProfile := rowToQueueProfile(row)

	var start = time.Now()
	o := ocr.New(faultyProfile.File, 0)

	o.Split()
	o.Process()
	o.CleanUp()

	model.UpdateProfile(faultyProfile.Profile.Id, o.Profile)

	processTime := time.Now().Sub(start).Nanoseconds() / 1e6
	model.SetQueueProcessed(faultyInt, processTime, faultyProfile.Profile.Id)

	http.Redirect(
		w, r,
		fmt.Sprintf("/admin/process/%s/%s/%s", owner, faulty, previous),
		302,
	)
}

func ProcessQueueHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Get(r, "tiers")
	userid, _ := session.Values["user"]

	owner := mux.Vars(r)["owner"]
	ownerInt, _ := strconv.Atoi(owner)

	faulty := mux.Vars(r)["faulty"]
	faultyInt, _ := strconv.Atoi(faulty)
	row := model.GetQueueWithProfileById(faultyInt)
	defer row.Close()

	row.Next()
	faultyProfile := rowToQueueProfile(row)

	previous := mux.Vars(r)["previous"]
	previousInt, _ := strconv.Atoi(previous)
	row = model.GetQueueWithProfileById(previousInt)
	defer row.Close()

	row.Next()
	previousProfile := rowToQueueProfile(row)

	o := ocr.New(faultyProfile.File, 0)

	o.Split()
	o.Process()
	o.CleanUp()

	templates := loadTemplates(
		"header.html",
		"footer.html",
		"nav.html",
		"admin/processQueue.html",
	)

	templates.ExecuteTemplate(w, "processQueue", struct {
		User     int
		Owner    int
		File     string
		Faulty   userQueueEntry
		Previous userQueueEntry
		Dry      profile.Profile
	}{
		User:     userid.(int),
		Owner:    ownerInt,
		Faulty:   faultyProfile,
		Previous: previousProfile,
		Dry:      o.Profile,
	})
}

func ProcessListHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Get(r, "tiers")
	userid, _ := session.Values["user"]

	owner := mux.Vars(r)["owner"]
	ownerInt, _ := strconv.Atoi(owner)
	rows := model.GetAllQueuesWithProfileByUser(ownerInt)
	defer rows.Close()

	var list []userQueueEntry
	for rows.Next() {
		e := rowToQueueProfile(rows)
		list = append(list, e)
	}

	// First entry is always seen as good.
	list[0].HasProblem = false
	for i := 1; i < len(list)-1; i++ {
		list[i].Previous = list[i-1].Id
		list[i].HasProblem = comparePrevious(list[i-1].Profile, list[i].Profile)
	}

	templates := loadTemplates(
		"header.html",
		"footer.html",
		"nav.html",
		"admin/processList.html",
	)

	templates.ExecuteTemplate(w, "processList", struct {
		User  int
		Owner int
		List  []userQueueEntry
	}{
		User:  userid.(int),
		Owner: ownerInt,
		List:  list,
	})
}

func ProcessIndexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Get(r, "tiers")
	userid, _ := session.Values["user"]

	rows := model.GetAllUserQueues()
	defer rows.Close()

	var list []userQueue
	for rows.Next() {
		var entry userQueue
		rows.Scan(&entry.Id, &entry.Latest, &entry.Count)

		list = append(list, entry)
	}

	templates := loadTemplates(
		"header.html",
		"footer.html",
		"nav.html",
		"admin/processIndex.html",
	)

	templates.ExecuteTemplate(w, "processIndex", struct {
		User int
		List []userQueue
	}{
		User: userid.(int),
		List: list,
	})
}
