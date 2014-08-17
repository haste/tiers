package page

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"tiers/conf"
	"tiers/queue"
	"tiers/session"
	"time"
)

type UploadPage struct {
	User int
}

var templates *template.Template

func UploadViewHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if templates, err = template.New("").ParseFiles(
		"templates/header.html",
		"templates/footer.html",
		"templates/nav.html",
		"templates/upload.html",
	); err != nil {
		log.Fatal(err)
	}

	session, _ := session.Get(r, "tiers")
	userid, ok := session.Values["user"]

	if !ok {
		return
	}

	templates.ExecuteTemplate(w, "upload", UploadPage{
		User: userid.(int),
	})
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Get(r, "tiers")
	userid, ok := session.Values["user"]

	if !ok {
		return
	}

	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// XXX: Handle errors.
	var db, _ = sql.Open("mysql", conf.Config.Database)
	defer db.Close()

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		//if part.FileName() is empty, skip this iteration.
		if part.FileName() == "" {
			continue
		}

		// profile_20140815_135412_0.png
		r := regexp.MustCompile("^(ingress|profile)_(\\d+)_(\\d+)_\\d+\\.png$")
		if r.MatchString(part.FileName()) != true {
			// XXX: Should probably handle this..
			continue
		}

		m := r.FindStringSubmatch(part.FileName())
		t, _ := time.ParseInLocation("20060102150405", m[2]+m[3], time.Local)

		var fileName = fmt.Sprintf("%d_%s", userid, part.FileName())

		dst, err := os.Create(conf.Config.Cache + fileName)
		defer dst.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// XXX: Handle errors...
		db.Exec(`
		INSERT INTO tiers_queues(user_id, timestamp, file)
		VALUES(?, ?, ?)
		`, userid, t.Unix(), fileName)
	}

	queue.Queue <- true
}
