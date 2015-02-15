package page

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"tiers/model"
	"tiers/profile"
	"tiers/session"
)

// Periods:
// Daily: 24 hours
// Previous: Last two uploads
// Weekly: 7 days
// Monthly: Same day, previous month?

type ProfilePageData struct {
	Profile profile.Profile
	Diff    interface{}
	Int64   int64
	Queue   int
}

type ProfilePage struct {
	User int
	Data ProfilePageData
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
		period := mux.Vars(r)["period"]
		fmt.Println(period)

		var p profile.Profile
		var diff profile.Profile

		queue := model.GetNumQueuedProfiles(userid.(int))
		profiles := model.GetNewestProfiles(userid.(int), 2)

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
				Profile: p,
				Diff:    diff,
				Queue:   queue,
			},
		})
	} else {
		templates.ExecuteTemplate(w, "index-unauthed", nil)
	}
}
