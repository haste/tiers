package page

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"tiers/conf"
	"tiers/session"

	"code.google.com/p/go.crypto/bcrypt"
)

func RegisterViewHandler(w http.ResponseWriter, r *http.Request) {
	templates := LoadTemplates(
		"header.html",
		"footer.html",
		"nav.html",
		"register.html",
	)

	session, _ := session.Get(r, "tiers")

	if _, ok := session.Values["user"]; ok {
		http.Redirect(w, r, "/", 302)
		return
	} else {
		templates.ExecuteTemplate(w, "register", nil)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	email, password := r.PostFormValue("email"), r.PostFormValue("password")

	re := regexp.MustCompile("^.+@.+\\..+$")
	if re.MatchString(email) != true {
		fmt.Fprintf(w, "Invalid e-mail provided.")
		return
	}

	if len(password) == 0 {
		fmt.Fprintf(w, "Password can't be empty.")
		return
	}

	var db, _ = sql.Open("mysql", conf.Config.Database)
	defer db.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if nil != err {
		log.Fatal(err)
	}

	db.Exec(`
		INSERT INTO tiers_users (email, password, valid_email)
		VALUES(?, ?, 0)
		`,
		email, hash,
	)

	http.Redirect(w, r, "/", 302)
}
