package model

import (
	"log"
	"tiers/profile"

	"github.com/lann/squirrel"
)

func InsertProfile(user_id, timestamp int, p profile.Profile) int64 {
	// Handle errors
	res, err := db.Exec(`
			INSERT INTO tiers_profiles (user_id, timestamp, agent, level, ap, unique_portals_visited, portals_discovered,
			xm_collected, hacks, resonators_deployed, links_created, control_fields_created, mind_units_captured,
			longest_link_ever_created, largest_control_field, xm_recharged, portals_captured, unique_portals_captured,
			mods_deployed, resonators_destroyed, portals_neutralized, enemy_links_destroyed, enemy_control_fields_destroyed,
			distance_walked, max_time_portal_held, max_time_link_maintained, max_link_length_x_days, max_time_field_held,
			largest_field_mus_x_days, glyph_hack_points, consecutive_days_hacking, unique_missions_completed,
			agents_successfully_recruited, innovator)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
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
		p.GlyphHackPoints, p.ConsecutiveDaysHacking,
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

func UpdateProfile(profileId int64, p profile.Profile) {
	sdb.Update("tiers_profiles").
		SetMap(
		squirrel.Eq{
			// General
			"agent": p.Nick,
			"level": p.Level,
			"ap":    p.AP,

			// Discovery
			"unique_portals_visited": p.UniquePortalsVisited,
			"portals_discovered":     p.PortalsDiscovered,
			"xm_collected":           p.XMCollected,

			// Health
			"distance_walked": p.DistanceWalked,

			// Building
			"resonators_deployed":       p.ResonatorsDeployed,
			"links_created":             p.LinksCreated,
			"control_fields_created":    p.ControlFieldsCreated,
			"mind_units_captured":       p.MindUnitsCaptured,
			"longest_link_ever_created": p.LongestLinkEverCreated,
			"largest_control_field":     p.LargestControlField,
			"xm_recharged":              p.XMRecharged,
			"portals_captured":          p.PortalsCaptured,
			"unique_portals_captured":   p.UniquePortalsCaptured,
			"mods_deployed":             p.ModsDeployed,

			// Combat
			"resonators_destroyed":           p.ResonatorsDestroyed,
			"portals_neutralized":            p.PortalsNeutralized,
			"enemy_links_destroyed":          p.EnemyLinksDestroyed,
			"enemy_control_fields_destroyed": p.EnemyControlFieldsDestroyed,

			// Defense
			"max_time_portal_held":     p.MaxTimePortalHeld,
			"max_time_link_maintained": p.MaxTimeLinkMaintained,
			"max_link_length_x_days":   p.MaxLinkLengthXDays,
			"max_time_field_held":      p.MaxTimeFieldHeld,
			"largest_field_mus_x_days": p.LargestFieldMUsXDays,

			// Missions
			"unique_missions_completed": p.UniqueMissionsCompleted,

			// Resource Gathering
			"hacks":                    p.Hacks,
			"glyph_hack_points":        p.GlyphHackPoints,
			"consecutive_days_hacking": p.ConsecutiveDaysHacking,

			// Mentoring
			"agents_successfully_recruited": p.AgentsSuccessfullyRecruited,

			// Badges
			"innovator": p.InnovatorLevel,
		}).
		Where("id = ?", profileId).
		Exec()
}
