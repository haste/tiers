package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"tiers/conf"
	"tiers/page"
	"tiers/queue"
	"tiers/session"

	"code.google.com/p/go.crypto/bcrypt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
)

var templates *template.Template

type User struct {
	Id          int
	Email       string
	Password    string
	Valid_email bool
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username, password := r.PostFormValue("email"), r.PostFormValue("password")

	u := User{}

	var db, _ = sql.Open("mysql", conf.Config.Database)
	defer db.Close()

	row := db.QueryRow(`
		SELECT id, email, password, valid_email FROM tiers_users WHERE email = ?`,
		username)

	row.Scan(&u.Id, &u.Email, &u.Password, &u.Valid_email)

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		fmt.Fprintln(w, "Invalid password or username!")
		return
	}

	session.Set(w, r, u.Id)
	http.Redirect(w, r, "/", 302)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", page.ProfileHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/badges", page.BadgesHandler)

	r.HandleFunc("/register", page.RegisterViewHandler).Methods("GET")
	r.HandleFunc("/register", page.RegisterHandler).Methods("POST")

	r.HandleFunc("/upload", page.UploadViewHandler).Methods("GET")
	r.HandleFunc("/upload", page.UploadHandler).Methods("POST")

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css/"))))
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts/"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js/"))))
	r.PathPrefix("/vendor/").Handler(http.StripPrefix("/vendor/", http.FileServer(http.Dir("static/vendor/"))))

	http.Handle("/", r)

	go queue.ProcessQueue()
	queue.Queue <- true

	log.Fatal(http.ListenAndServeTLS("localhost:45633", conf.Config.Cert, conf.Config.Key, nil))
}
