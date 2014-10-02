package page

import (
	"fmt"
	"net/http"
	"regexp"

	"tiers/model"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
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

	_, err := model.CreateUser(email, password)
	switch {
	case err == model.ErrEmailAlreadyUsed:
		fmt.Fprintf(w, "An user is already registered on that mail.")
		return
	}

	http.Redirect(w, r, "/", 302)
}
