package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"tiers/conf"

	"code.google.com/p/go.crypto/bcrypt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var templates *template.Template
var CookieStore *sessions.CookieStore

type User struct {
	Id          int
	Email       string
	Password    string
	Valid_email bool
}

func init() {
	var err error

	if templates, err = template.New("").ParseFiles(
		"templates/base.html",
	); err != nil {
		log.Fatal(err)
	}
}

func setSession(w http.ResponseWriter, r *http.Request, userId int) {
	session, _ := CookieStore.Get(r, "tiers")
	session.Values["user"] = userId
	session.Save(r, w)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := CookieStore.Get(r, "tiers")

	if userid, ok := session.Values["user"]; ok {
		log.Println(userid)
	} else {
		log.Println("No cookie")
	}

	templates.ExecuteTemplate(w, "base.html", "")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username, password := r.FormValue("username"), r.FormValue("password")

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

	setSession(w, r, u.Id)
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
	CookieStore = sessions.NewCookieStore([]byte(conf.Config.CookieSecret))

	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/register", RegisterHandler)

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css/"))))
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts/"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js/"))))
	r.PathPrefix("/vendor/").Handler(http.StripPrefix("/vendor/", http.FileServer(http.Dir("static/vendor/"))))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
