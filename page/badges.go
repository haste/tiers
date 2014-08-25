package page

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"tiers/profile"
	"tiers/session"
	"tiers/user"
)

type BadgePage struct {
	User         int
	Measurements template.JS
}

func BadgesHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Get(r, "tiers")
	userid, ok := session.Values["user"]

	if !ok {
		return
	}

	templates := LoadTemplates(
		"header.html",
		"footer.html",
		"nav.html",
		"badges.html",
	)

	p := user.GetNewestProfile(userid.(int))

	profile.HandleBadges(&p)

	bp := profile.BuildBadgeProgress(p)
	v, err := json.Marshal(bp)

	// XXX: Handle err
	if err != nil {
		log.Fatal(err)
	}

	page := &BadgePage{
		User:         userid.(int),
		Measurements: template.JS(v),
	}

	templates.ExecuteTemplate(w, "badges", page)
}
