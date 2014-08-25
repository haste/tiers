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
	http.Redirect(w, r, "/", 302)
}

func LogoutHandle(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Get(r, "tiers")
	delete(session.Values, "user")
	session.Save(r, w)

	http.Redirect(w, r, "/", 302)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", page.ProfileHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/logout", LogoutHandle)
	r.HandleFunc("/badges", page.BadgesHandler)

	r.HandleFunc("/register", page.RegisterViewHandler).Methods("GET")
	r.HandleFunc("/register", page.RegisterHandler).Methods("POST")

	r.HandleFunc("/upload", page.UploadViewHandler).Methods("GET")
	r.HandleFunc("/upload", page.UploadHandler).Methods("POST")

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(rice.MustFindBox("static/css/").HTTPBox())))
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(rice.MustFindBox("static/fonts/").HTTPBox())))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(rice.MustFindBox("static/js/").HTTPBox())))
	r.PathPrefix("/vendor/").Handler(http.StripPrefix("/vendor/", http.FileServer(rice.MustFindBox("static/vendor/").HTTPBox())))

	http.Handle("/", r)

	go queue.ProcessQueue()
	queue.Queue <- true

	log.Fatal(http.ListenAndServeTLS("localhost:45633", conf.Config.Cert, conf.Config.Key, nil))
}
