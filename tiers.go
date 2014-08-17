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

type IndexPage struct {
	User int
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	if templates, err = template.New("").ParseFiles(
		"templates/header.html",
		"templates/footer.html",
		"templates/nav.html",
		"templates/index-unauthed.html",
		"templates/index-authed.html",
	); err != nil {
		log.Fatal(err)
	}

	session, _ := session.Get(r, "tiers")

	if userid, ok := session.Values["user"]; ok {
		templates.ExecuteTemplate(w, "index-authed", IndexPage{
			User: userid.(int),
		})
	} else {
		templates.ExecuteTemplate(w, "index-unauthed", nil)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username, password := r.PostFormValue("email"), r.PostFormValue("password")

	fmt.Println(username, password)

	u := User{}

	var db, _ = sql.Open("mysql", conf.Config.Database)
	defer db.Close()

	row := db.QueryRow(`
		SELECT id, email, password, valid_email FROM tiers_users WHERE email = ?`,
		username)

	row.Scan(&u.Id, &u.Email, &u.Password, &u.Valid_email)

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		log.Print("Invalid password! ")
	}

	session.Set(w, r, u.Id)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username, password := r.FormValue("username"), r.FormValue("password")

	var db, _ = sql.Open("mysql", conf.Config.Database)
	defer db.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if nil != err {
		log.Fatal(err)
	}

	result, err := db.Exec(`
		INSERT INTO tiers_users (email, password, valid_email)
		VALUES(?, ?, 1)
		`,
		username, hash,
	)

	log.Printf("%s", hash)

	log.Print(result, err)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/register", RegisterHandler)
	r.HandleFunc("/badges", page.BadgesHandler)

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
