package model

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"net"

	"github.com/gorilla/securecookie"

	"crypto/sha256"

	"code.google.com/p/go.crypto/bcrypt"

	"tiers/profile"
	"time"
)

var ErrUserNotFound = errors.New("User not found.")
var ErrEmailAlreadyUsed = errors.New("E-mail already used.")

type User struct {
	Id          int
	Email       string
	GPlusId     string
	Password    string
	AccessToken string
	Valid_email bool
}

func GetUserByMail(email string) (*User, error) {
	u := new(User)

	err := db.QueryRow(`
		SELECT id, email, gplus_id, password, access_token, valid_email FROM tiers_users WHERE email = ?`,
		email,
	).Scan(&u.Id, &u.Email, &u.GPlusId, &u.Password, &u.AccessToken, &u.Valid_email)

	switch {
	case err == sql.ErrNoRows:
		return nil, ErrUserNotFound
	case err != nil:
		log.Fatal(err)
	}

	return u, nil
}

func GetUserById(id int) (*User, error) {
	u := new(User)

	err := db.QueryRow(`
		SELECT id, email, gplus_id, password, access_token, valid_email FROM tiers_users WHERE id = ?`,
		id,
	).Scan(&u.Id, &u.Email, &u.GPlusId, &u.Password, &u.AccessToken, &u.Valid_email)

	switch {
	case err == sql.ErrNoRows:
		return nil, ErrUserNotFound
	case err != nil:
		log.Fatal(err)
	}

	return u, nil
}

func GetUserByGPlusId(gplusId string) (*User, error) {
	u := new(User)

	err := db.QueryRow(`
		SELECT id, email, gplus_id, password, access_token, valid_email FROM tiers_users WHERE gplus_id = ?`,
		gplusId,
	).Scan(&u.Id, &u.Email, &u.GPlusId, &u.Password, &u.AccessToken, &u.Valid_email)

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

	return &User{Id: int(id), Email: email, Password: string(hash), Valid_email: false}, nil
}

func CreateGPlusUser(gplusId, accessToken string) *User {
	res, _ := db.Exec(`
		INSERT INTO tiers_users (gplus_id, access_token)
		VALUES(?, ?)
		`,
		gplusId, accessToken,
	)

	id, _ := res.LastInsertId()

	return &User{Id: int(id), GPlusId: gplusId, AccessToken: accessToken}
}

func SetResetPassword(user_id int, ip net.IP) string {
	key := securecookie.GenerateRandomKey(32)

	hash := sha256.New()
	hash.Write(key)
	token := hex.EncodeToString(hash.Sum(nil))
	expires := time.Now().Add(time.Hour * 24).Unix()

	db.Exec(`
	INSERT INTO tiers_reset_password(user_id, expires, ip, token)
	VALUES(?, ?, INET6_ATON(?), ?)
	`,
		user_id, expires, ip.String(), token,
	)

	return token
}

func ValidateResetToken(token string) int {
	var user_id int
	db.QueryRow(`
		SELECT user_id FROM tiers_reset_password
		WHERE token = ?
	`, token).Scan(&user_id)

	return user_id
}

func DeleteResetToken(token string) {
	db.Exec("DELETE FROM tiers_reset_password WHERE token = ?", token)
}

func SetUserPassword(user_id int, password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	db.Exec(`
		UPDATE tiers_users
		SET password = ?
		WHERE id = ?
	`, hash, user_id)
}

func UpdateGPlusToken(id int, accessToken string) {
	db.Exec(`
		UPDATE tiers_users
		SET access_token = ?
		WHERE id = ?
	`, accessToken, id)
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
	defer rows.Close()

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
	defer row.Close()

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

func GetNewestProfiles(user_id, limit int) []profile.Profile {
	var profiles []profile.Profile

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
		LIMIT ?
		`, user_id, limit)
	defer row.Close()

	for row.Next() {
		var p profile.Profile
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

		profiles = append(profiles, p)
	}

	return profiles
}
