package page

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
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

func UploadViewHandler(w http.ResponseWriter, r *http.Request) {
	templates := loadTemplates(
		"header.html",
		"footer.html",
		"nav.html",
		"upload.html",
	)

	session, _ := session.Get(r, "tiers")
	userid, ok := session.Values["user"]

	if !ok {
		http.Redirect(w, r, "/", 302)
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
		if err != nil {
			break
		}

		// if part.FileName() is empty, skip this iteration.
		if part.FileName() == "" {
			continue
		}

		var fileName string
		var t time.Time
		// profile_20140815_135412_0.png
		r := regexp.MustCompile("^(ingress|profile)_(\\d+)_(\\d+)_\\d+\\.png$")
		if r.MatchString(part.FileName()) {
			m := r.FindStringSubmatch(part.FileName())
			t, _ = time.ParseInLocation("20060102150405", m[2]+m[3], time.Local)

			fileName = fmt.Sprintf("%d_%s", userid, part.FileName())
		} else {
			t = time.Now()
			fileName = fmt.Sprintf("%d_%s.png", userid, t.Format("20060102_150405"))
		}

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

	var numQueue float32
	db.QueryRow(`SELECT count(id) FROM tiers_queues WHERE processed = 0`).Scan(&numQueue)

	queueText := fmt.Sprintf(
		"Your file has been added to the queue and should be processed within %.1f seconds.",
		numQueue*6.6,
	)

	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		data, _ := json.Marshal(map[string]interface{}{
			"success": true,
			"message": queueText,
		})
		fmt.Fprintf(w, "%s", data)
	} else {
		http.Redirect(w, r, "/", 302)
	}

	queue.Queue <- true
}
