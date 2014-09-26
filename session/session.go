package session

import (
	"net/http"

	"tiers/conf"

	"github.com/gorilla/sessions"
)

var CookieStore *sessions.CookieStore

func init() {
	CookieStore = sessions.NewCookieStore(
		conf.Config.CookieHashKey,
		conf.Config.CookieBlockKey,
	)
}

func Set(w http.ResponseWriter, r *http.Request, userId int) {
	session, _ := CookieStore.Get(r, "tiers")
	session.Values["user"] = userId
	session.Save(r, w)
}

func Get(r *http.Request, field string) (*sessions.Session, error) {
	return CookieStore.Get(r, field)
}
