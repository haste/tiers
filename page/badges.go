package page

import (
	"log"
	"net/http"

	"tiers/user"
	"tiers/session"
)

func BadgesHandler(w http.ResponseWriter, r *http.Request) {
	s, _ := session.Get(r, "tiers")

	uid := s.Values["user"].(int)

	profiles := user.GetAllProfiles(uid)

	log.Printf("%+v", profiles)
}
