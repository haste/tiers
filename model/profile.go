package model

import (
	"log"
	"tiers/profile"
)

func InsertProfile(user_id, timestamp int, p profile.Profile) int64 {
	// Handle errors
	res, err := db.Exec(`
			INSERT INTO tiers_profiles (user_id, timestamp, agent, level, ap, unique_portals_visited, portals_discovered,
			xm_collected, hacks, resonators_deployed, links_created, control_fields_created, mind_units_captured,
			longest_link_ever_created, largest_control_field, xm_recharged, portals_captured, unique_portals_captured,
			mods_deployed, resonators_destroyed, portals_neutralized, enemy_links_destroyed, enemy_control_fields_destroyed,
			distance_walked, max_time_portal_held, max_time_link_maintained, max_link_length_x_days, max_time_field_held,
			largest_field_mus_x_days, glyph_hack_points, unique_missions_completed, agents_successfully_recruited, innovator)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
			`,
		user_id, timestamp,
		p.Nick, p.Level, p.AP,
		p.UniquePortalsVisited, p.PortalsDiscovered, p.XMCollected,
		p.Hacks, p.ResonatorsDeployed, p.LinksCreated, p.ControlFieldsCreated, p.MindUnitsCaptured, p.LongestLinkEverCreated,
		p.LargestControlField, p.XMRecharged, p.PortalsCaptured, p.UniquePortalsCaptured, p.ModsDeployed,
		p.ResonatorsDestroyed, p.PortalsNeutralized, p.EnemyLinksDestroyed, p.EnemyControlFieldsDestroyed,
		p.DistanceWalked,
		p.MaxTimePortalHeld, p.MaxTimeLinkMaintained, p.MaxLinkLengthXDays, p.MaxTimeFieldHeld,
		p.LargestFieldMUsXDays,
		p.GlyphHackPoints,
		p.UniqueMissionsCompleted,
		p.AgentsSuccessfullyRecruited,
		p.InnovatorLevel,
	)

	if err != nil {
		log.Fatal(err)
	}

	insertId, _ := res.LastInsertId()

	return insertId
}
