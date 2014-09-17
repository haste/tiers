package model

import (
	"database/sql"
	"errors"
	"log"

	"code.google.com/p/go.crypto/bcrypt"

	"tiers/profile"
)

var ErrUserNotFound = errors.New("User not found.")
var ErrEmailAlreadyUsed = errors.New("E-mail already used.")

type User struct {
	Id          int
	Email       string
	Password    string
	Valid_email bool
}

func GetUserByMail(email string) (*User, error) {
	u := new(User)

	err := db.QueryRow(`
		SELECT id, email, password, valid_email FROM tiers_users WHERE email = ?`,
		email,
	).Scan(&u.Id, &u.Email, &u.Password, &u.Valid_email)

	switch {
	case err == sql.ErrNoRows:
		return nil, ErrUserNotFound
	case err != nil:
		log.Fatal(err)
	}

	return u, nil
}

func SignInUser(email, password string) (*User, error) {
	u, err := GetUserByMail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, ErrUserNotFound
	}

	return u, nil
}

func isEmailUsed(email string) bool {
	_, err := GetUserByMail(email)

	if err == ErrUserNotFound {
		return false
	}

	return true
}

func CreateUser(email, password string) (*User, error) {
	if isEmailUsed(email) {
		return nil, ErrEmailAlreadyUsed
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if nil != err {
		return nil, err
	}

	res, _ := db.Exec(`
		INSERT INTO tiers_users (email, password, valid_email)
		VALUES(?, ?, 0)
		`,
		email, hash,
	)

	id, _ := res.LastInsertId()

	return &User{int(id), email, string(hash), false}, nil
}

func GetAllProfiles(user_id int) []profile.Profile {
	// XXX: Handle errors.
	rows, _ := db.Query(`
		(SELECT id, user_id, timestamp, agent, level, ap,
		unique_portals_visited, portals_discovered, xm_collected,
		hacks, resonators_deployed, links_created, control_fields_created, mind_units_captured,
		longest_link_ever_created, largest_control_field, xm_recharged, portals_captured,
		unique_portals_captured,
		resonators_destroyed, portals_neutralized, enemy_links_destroyed, enemy_control_fields_destroyed,
		distance_walked,
		max_time_portal_held, max_time_link_maintained, max_link_length_x_days, max_time_field_held,
		largest_field_mus_x_days
		FROM tiers_profiles
		WHERE user_id = ?
		ORDER BY timestamp DESC
		LIMIT 50) ORDER BY timestamp ASC
		`, user_id)

	var profiles []profile.Profile
	for rows.Next() {
		var p profile.Profile

		rows.Scan(
			&p.Id, &p.UserId, &p.Timestamp,
			&p.Nick, &p.Level, &p.AP,
			&p.UniquePortalsVisited, &p.PortalsDiscovered, &p.XMCollected,
			&p.Hacks, &p.ResonatorsDeployed, &p.LinksCreated, &p.ControlFieldsCreated, &p.MindUnitsCaptured,
			&p.LongestLinkEverCreated, &p.LargestControlField, &p.XMRecharged, &p.PortalsCaptured,
			&p.UniquePortalsCaptured,
			&p.ResonatorsDestroyed, &p.PortalsNeutralized, &p.EnemyLinksDestroyed, &p.EnemyControlFieldsDestroyed,
			&p.DistanceWalked,
			&p.MaxTimePortalHeld, &p.MaxTimeLinkMaintained, &p.MaxLinkLengthXDays, &p.MaxTimeFieldHeld,
			&p.LargestFieldMUsXDays,
		)

		profiles = append(profiles, p)
	}

	return profiles
}

func GetNewestProfile(user_id int) profile.Profile {
	// XXX: Handle errors.
	row, _ := db.Query(`
		SELECT id, user_id, timestamp, agent, level, ap,
		unique_portals_visited, portals_discovered, xm_collected,
		hacks, resonators_deployed, links_created, control_fields_created, mind_units_captured,
		longest_link_ever_created, largest_control_field, xm_recharged, portals_captured,
		unique_portals_captured,
		resonators_destroyed, portals_neutralized, enemy_links_destroyed, enemy_control_fields_destroyed,
		distance_walked,
		max_time_portal_held, max_time_link_maintained, max_link_length_x_days, max_time_field_held,
		largest_field_mus_x_days
		FROM tiers_profiles
		WHERE user_id = ?
		ORDER BY timestamp DESC
		LIMIT 1
		`, user_id)

	var p profile.Profile

	row.Next()
	row.Scan(
		&p.Id, &p.UserId, &p.Timestamp,
		&p.Nick, &p.Level, &p.AP,
		&p.UniquePortalsVisited, &p.PortalsDiscovered, &p.XMCollected,
		&p.Hacks, &p.ResonatorsDeployed, &p.LinksCreated, &p.ControlFieldsCreated, &p.MindUnitsCaptured,
		&p.LongestLinkEverCreated, &p.LargestControlField, &p.XMRecharged, &p.PortalsCaptured,
		&p.UniquePortalsCaptured,
		&p.ResonatorsDestroyed, &p.PortalsNeutralized, &p.EnemyLinksDestroyed, &p.EnemyControlFieldsDestroyed,
		&p.DistanceWalked,
		&p.MaxTimePortalHeld, &p.MaxTimeLinkMaintained, &p.MaxLinkLengthXDays, &p.MaxTimeFieldHeld,
		&p.LargestFieldMUsXDays,
	)

	return p
}
