package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"tiers/conf"
	"tiers/model"
	"tiers/page"
	"tiers/queue"
	"tiers/session"

	"github.com/GeertJohan/go.rice"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
)

var templates *template.Template

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email, password := r.PostFormValue("email"), r.PostFormValue("password")

	u, err := model.SignInUser(email, password)
	switch {
	case err == model.ErrUserNotFound:
		fmt.Fprintln(w, "Invalid username or password.")
		return
	}

	session.Set(w, r, u.Id)

	numProfiles := model.GetNumProfiles(u.Id)
	numQueued := model.GetNumQueuedProfiles(u.Id)
	switch {
	case numProfiles == 0 && numQueued == 0:
		http.Redirect(w, r, "/upload", 302)
	default:
		http.Redirect(w, r, "/", 302)
	}
}

func LogoutHandle(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Get(r, "tiers")
	userid, ok := session.Values["user"]

	if !ok {
		http.Redirect(w, r, "/", 302)
		return
	}

	u, _ := model.GetUserById(userid.(int))
	if u.GPlusId != "" {
		// Execute HTTP GET request to revoke current token
		url := "https://accounts.google.com/o/oauth2/revoke?token=" + u.AccessToken
		resp, _ := http.Get(url)
		defer resp.Body.Close()
	}

	delete(session.Values, "user")
	session.Save(r, w)

	http.Redirect(w, r, "/", 302)
}

func main() {
	r := mux.NewRouter()

	rice.MustFindBox("templates")

	r.HandleFunc("/", page.ProfileHandler)
	r.HandleFunc("/badges", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/progress", 302)
	})
	r.HandleFunc("/progress", page.ProgressHandler)

	r.HandleFunc("/signin", LoginHandler)
	r.HandleFunc("/logout", LogoutHandle)
	r.HandleFunc("/signup", page.SignupHandler).Methods("POST")
	r.HandleFunc("/reset_password/{token:[a-f0-9]+}", page.ResetPassViewHandler).Methods("GET")
	r.HandleFunc("/reset_password/{token:[a-f0-9]+}", page.ResetPassHandler).Methods("POST")
	r.HandleFunc("/reset_password", page.ResetPassMailHandler).Methods("POST")
	r.HandleFunc("/gplus", page.GPlusHandler)

	r.HandleFunc("/upload", page.UploadViewHandler).Methods("GET")
	r.HandleFunc("/upload", page.UploadHandler).Methods("POST")

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(rice.MustFindBox("static/css").HTTPBox())))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(rice.MustFindBox("static/images").HTTPBox())))
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(rice.MustFindBox("static/fonts").HTTPBox())))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(rice.MustFindBox("static/js").HTTPBox())))
	r.PathPrefix("/vendor/").Handler(http.StripPrefix("/vendor/", http.FileServer(rice.MustFindBox("static/vendor").HTTPBox())))

	http.Handle("/", r)

	go queue.ProcessQueue()
	queue.Queue <- true

	log.Fatal(http.ListenAndServeTLS(conf.Config.Address, conf.Config.Cert, conf.Config.Key, nil))
}
