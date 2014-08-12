package user

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"

	"tiers/conf"
	"tiers/profile"
)

func GetAllProfiles(user_id int) []profile.Profile {
	// XXX: Handle errors.
	var db, _ = sql.Open("mysql", conf.Config.Database)
	defer db.Close()

	// XXX: Handle errors.
	rows, _ := db.Query(`
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
	var db, _ = sql.Open("mysql", conf.Config.Database)
	defer db.Close()

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
