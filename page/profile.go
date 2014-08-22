package page

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"tiers/session"
	"tiers/user"
)

type ProfilePage struct {
	User int
	Data interface{}
}

func comma(n uint) string {
	var h []byte
	var s = strconv.Itoa(int(n))

	for i := len(s) - 1; i >= 0; i-- {
		o := len(s) - 1 - i
		if o%3 == 0 && o != 0 {
			h = append(h, ',')
		}

		h = append(h, s[i])
	}

	for i, j := 0, len(h)-1; i < j; i, j = i+1, j-1 {
		h[i], h[j] = h[j], h[i]
	}

	return string(h)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Get(r, "tiers")
	userid, ok := session.Values["user"]

	var err error

	if templates, err = template.New("").Funcs(template.FuncMap{
		"comma": comma,
	}).ParseFiles(
		"templates/header.html",
		"templates/footer.html",
		"templates/nav.html",
		"templates/index-unauthed.html",
		"templates/profile.html",
	); err != nil {
		log.Fatal(err)
	}

	if ok {
		p := user.GetNewestProfile(userid.(int))

		templates.ExecuteTemplate(w, "profile", ProfilePage{
			User: userid.(int),
			Data: p,
		})
	} else {
		templates.ExecuteTemplate(w, "index-unauthed", nil)
	}
}
