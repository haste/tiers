package page

import (
	"net/http"

	"tiers/model"
	"tiers/session"
)

type ProfilePage struct {
	User int
	Data interface{}
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

	// Extra:
	// Average MU per field.

	if ok {
		p := model.GetNewestProfile(userid.(int))

		templates.ExecuteTemplate(w, "profile", ProfilePage{
			User: userid.(int),
			Data: p,
		})
	} else {
		templates.ExecuteTemplate(w, "index-unauthed", nil)
	}
}
