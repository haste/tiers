package page

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/smtp"

	"github.com/gorilla/mux"

	"text/template"
	"tiers/conf"
	"tiers/model"
	"tiers/session"
)

var mailTemplate = `From: noreply@tie.rs
To: {{.To}}
Subject: {{.Subject}}
MIME-version: 1.0
Content-Type: text/html; charset="UTF-8"

<div>
	<strong>Reset password</strong>
	<p>Use the link below to reset your password. If you did not request a
	password reset, then simply ignore this message.</p>
	<p><a href="https://tie.rs/reset_password/{{.Token}}">Reset password</a></p>
	<p>This link is only valid for 24 hours.</p>
</div>
`

type TokenResetPage struct {
	User int
	Data interface{}
}

func ResetPassViewHandler(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]

	templates := loadTemplates(
		"header.html",
		"footer.html",
		"nav.html",
		"resetpass.html",
	)

	user_id := model.ValidateResetToken(token)
	if user_id == 0 {
		fmt.Fprintf(w, "Invalid or expired token...")
		return
	}

	templates.ExecuteTemplate(w, "resetpass", TokenResetPage{
		Data: token,
	})
}

func ResetPassHandler(w http.ResponseWriter, r *http.Request) {
	password := r.PostFormValue("password")
	if len(password) == 0 {
		fmt.Fprintf(w, "Password can't be empty.")
		return
	}

	token := mux.Vars(r)["token"]
	user_id := model.ValidateResetToken(token)
	if user_id == 0 {
		fmt.Fprintf(w, "Invalid or expired token...")
		return
	}

	model.DeleteResetToken(token)
	model.SetUserPassword(user_id, password)

	session.Set(w, r, user_id)

	http.Redirect(w, r, "/", 302)
}

func ResetPassMailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	user, err := model.GetUserByMail(email)
	if err != nil {
		return
	}

	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	ip := net.ParseIP(host)

	token := model.SetResetPassword(user.Id, ip)

	c, err := smtp.Dial(conf.Config.SMTP)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Mail("noreply@tie.rs"); err != nil {
		log.Fatal(err)
	}

	if err := c.Rcpt(user.Email); err != nil {
		log.Fatal(err)
	}

	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}

	param := struct {
		From    string
		To      string
		Subject string
		Token   string
	}{
		"noreply@ti.rs",
		user.Email,
		"Reset password",
		token,
	}

	var msg bytes.Buffer
	tmpl, _ := template.New("mail").Parse(mailTemplate)
	tmpl.Execute(&msg, &param)

	if _, err := wc.Write(msg.Bytes()); err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
