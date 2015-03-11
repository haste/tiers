package model

import (
	"database/sql"
	"log"

	"github.com/lann/squirrel"
)

func getQueueWithProfile() squirrel.SelectBuilder {
	return sdb.
		Select(`
			q.id AS queue_id, q.file, q.ocr_profile, q.processed, q.processtime,

			p.id AS profile_id, p.timestamp, p.agent, p.level, p.ap,

			p.unique_portals_visited, p.portals_discovered, p.xm_collected,

			p.distance_walked,

			p.resonators_deployed, p.links_created, p.control_fields_created,
			p.mind_units_captured, p.longest_link_ever_created,
			p.largest_control_field, p.xm_recharged, p.portals_captured,
			p.unique_portals_captured, p.mods_deployed,

			p.resonators_destroyed, p.portals_neutralized, p.enemy_links_destroyed,
			p.enemy_control_fields_destroyed,

			p.max_time_portal_held, p.max_time_link_maintained,
			p.max_link_length_x_days, p.max_time_field_held,
			p.largest_field_mus_x_days,

			p.unique_missions_completed,

			p.hacks,
			p.glyph_hack_points,
			p.consecutive_days_hacking,

			p.agents_successfully_recruited,

			p.innovator`).
		From("tiers_queues q").
		Join("tiers_profiles p ON p.id = q.profile_id").
		OrderBy("p.timestamp ASC")
}

func GetPendingQueues() *sql.Rows {
	// XXX: Handle errors.
	rows, err := db.Query(`
		SELECT id, user_id, timestamp, file, processed, ocr_profile
		FROM tiers_queues
		WHERE processed = 0
		`)

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func GetAllUserQueues() *sql.Rows {
	query := squirrel.
		Select("q.user_id, MAX(q.timestamp) AS latest, COUNT(q.id) AS count").
		From("tiers_queues q").
		Join("tiers_users u ON u.id = q.user_id").
		GroupBy("q.user_id")

	// XXX: Handle errors.
	rows, err := query.RunWith(db).Query()

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func GetAllUserQueuesWithProfiles() *sql.Rows {
	query := getQueueWithProfile()

	// XXX: Handle errors.
	rows, err := query.Query()

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func GetAllQueuesWithProfileByUser(userId int) *sql.Rows {
	query := getQueueWithProfile().
		Where(squirrel.Eq{"p.user_id": userId})

	// XXX: Handle errors.
	rows, err := query.Query()

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func GetQueueWithProfileById(queueId int) *sql.Rows {
	query := getQueueWithProfile().
		Where(squirrel.Eq{"q.id": queueId})

	// XXX: Handle errors.
	row, err := query.Query()

	if err != nil {
		log.Fatal(err)
	}

	return row
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
