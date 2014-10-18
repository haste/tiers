package model

import (
	"database/sql"
	"log"
)

func GetPendingQueues() *sql.Rows {
	// XXX: Handle errors.
	rows, err := db.Query(`
		SELECT id, user_id, timestamp, file, processed
		FROM tiers_queues
		WHERE processed = 0
		`)

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func SetQueueProcessed(id int, processTime, profileId int64) {
	// Handle errors
	_, err := db.Exec(`
			UPDATE tiers_queues
			SET processed = 1,
			processtime = ?,
			profile_id = ?
			WHERE id = ?
			`, processTime, profileId, id)

	if err != nil {
		log.Fatal(err)
	}
}
