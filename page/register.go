package page

import (
	"fmt"
	"net/http"
	"regexp"
	"tiers/session"
)

func RegisterViewHandler(w http.ResponseWriter, r *http.Request) {
	templates := loadTemplates(
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

	http.Redirect(w, r, "/", 302)
}
