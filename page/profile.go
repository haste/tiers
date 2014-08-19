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

	for c := 0; n > 0; c++ {
		i := n % 10

		n -= i
		n /= 10

		if c == 3 {
			c = 0
			h = append(h, ',')
		}

		h = strconv.AppendUint(h, uint64(i), 10)
	}

	for i, j := 0, len(h)-1; i < j; i, j = i+1, j-1 {
		h[i], h[j] = h[j], h[i]
	}

	return string(h)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Get(r, "tiers")
	userid, ok := session.Values["user"]

	if !ok {
		return
	}

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
