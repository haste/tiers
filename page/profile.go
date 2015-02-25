package page

import (
	"net/http"

	"github.com/gorilla/mux"

	"tiers/model"
	"tiers/profile"
	"tiers/session"
	"time"
)

// Periods:
// Daily: 24 hours
// Previous: Last two uploads
// Weekly: 7 days
// Monthly: Same day, previous month?

type ProfilePageData struct {
	Period  string
	Profile profile.Profile
	Diff    interface{}
	Int64   int64
	Queue   int
}

type ProfilePage struct {
	User int
	Data ProfilePageData
}

func getOffset(period string) int64 {
	var (
		t   time.Time
		now = time.Now()
	)

	switch period {
	case "weekly":
		t = now.AddDate(0, 0, -7)
	case "monthly":
		t = now.AddDate(0, -1, 0)
	}

	return t.Unix()
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Get(r, "tiers")
	userid, ok := session.Values["user"]

	templates := loadTemplates(
		"header.html",
		"footer.html",
		"nav.html",
		"index-unauthed.html",
		"profile.html",
	)

	if ok {
		var profiles []profile.Profile

		period := mux.Vars(r)["period"]

		var p profile.Profile
		var diff profile.Profile

		queue := model.GetNumQueuedProfiles(userid.(int))

		switch period {
		case "daily":
			period = "Daily"
			p1 := model.GetNewestProfile(userid.(int))
			p2 := model.GetProfileByTimestamp(userid.(int), time.Now().AddDate(0, 0, -1).Unix())

			profiles = append(profiles, p1, p2)
		case "weekly":
			period = "Weekly"
			p1 := model.GetNewestProfile(userid.(int))
			p2 := model.GetProfileByTimestamp(userid.(int), time.Now().AddDate(0, 0, -7).Unix())

			profiles = append(profiles, p1, p2)
		case "monthly":
			period = "Monthly"
			p1 := model.GetNewestProfile(userid.(int))
			p2 := model.GetProfileByTimestamp(userid.(int), time.Now().AddDate(0, -1, 0).Unix())

			profiles = append(profiles, p1, p2)
		default:
			period = "Previous"
			profiles = model.GetNewestProfiles(userid.(int), 2)
		}

		switch len(profiles) {
		case 1:
			p = profiles[0]
		case 2:
			diff = profile.Diff(profiles[1], profiles[0])
			p = profiles[0]
		}

		templates.ExecuteTemplate(w, "profile", ProfilePage{
			User: userid.(int),
			Data: ProfilePageData{
				Period:  period,
				Profile: p,
				Diff:    diff,
				Queue:   queue,
			},
		})
	} else {
		templates.ExecuteTemplate(w, "index-unauthed", nil)
	}
}
