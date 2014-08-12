package page

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"tiers/user"
	"tiers/profile"
	"tiers/session"
)

func BadgesHandler(w http.ResponseWriter, r *http.Request) {
	s, _ := session.Get(r, "tiers")

	uid := s.Values["user"].(int)

	log.Println(uid)

	p := user.GetNewestProfile(uid)

	profile.HandleBadges(&p)

	bp := profile.BuildBadgeProgress(p)
	v, err := json.Marshal(bp)

	// XXX: Handle err
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s",  v)
}
