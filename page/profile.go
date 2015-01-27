package page

import (
	"net/http"

	"tiers/model"
	"tiers/profile"
	"tiers/session"
)

type ProfilePageData struct {
	Profile profile.Profile
	Diff    interface{}
	Int64   int64
	Queue   int
}

type ProfilePage struct {
	User int
	Data ProfilePageData
}

func createDiff(a, b profile.Profile) profile.Profile {
	var diff profile.Profile

	diff.Level = b.Level - a.Level
	diff.AP = b.AP - a.AP

	diff.Bronze = b.Bronze - a.Bronze
	diff.Silver = b.Silver - a.Silver
	diff.Gold = b.Gold - a.Gold
	diff.Platinum = b.Platinum - a.Platinum
	diff.Onyx = b.Onyx - a.Onyx

	diff.UniquePortalsVisited = b.UniquePortalsVisited - a.UniquePortalsVisited
	diff.PortalsDiscovered = b.PortalsDiscovered - a.PortalsDiscovered
	diff.XMCollected = b.XMCollected - a.XMCollected

	diff.Hacks = b.Hacks - a.Hacks
	diff.ResonatorsDeployed = b.ResonatorsDeployed - a.ResonatorsDeployed
	diff.LinksCreated = b.LinksCreated - a.LinksCreated
	diff.ControlFieldsCreated = b.ControlFieldsCreated - a.ControlFieldsCreated
	diff.MindUnitsCaptured = b.MindUnitsCaptured - a.MindUnitsCaptured
	diff.LongestLinkEverCreated = b.LongestLinkEverCreated - a.LongestLinkEverCreated
	diff.LargestControlField = b.LargestControlField - a.LargestControlField
	diff.XMRecharged = b.XMRecharged - a.XMRecharged
	diff.PortalsCaptured = b.PortalsCaptured - a.PortalsCaptured
	diff.UniquePortalsCaptured = b.UniquePortalsCaptured - a.UniquePortalsCaptured
	diff.ModsDeployed = b.ModsDeployed - a.ModsDeployed

	diff.ResonatorsDestroyed = b.ResonatorsDestroyed - a.ResonatorsDestroyed
	diff.PortalsNeutralized = b.PortalsNeutralized - a.PortalsNeutralized
	diff.EnemyLinksDestroyed = b.EnemyLinksDestroyed - a.EnemyLinksDestroyed
	diff.EnemyControlFieldsDestroyed = b.EnemyControlFieldsDestroyed - a.EnemyControlFieldsDestroyed

	diff.DistanceWalked = b.DistanceWalked - a.DistanceWalked

	diff.MaxTimePortalHeld = b.MaxTimePortalHeld - a.MaxTimePortalHeld
	diff.MaxTimeLinkMaintained = b.MaxTimeLinkMaintained - a.MaxTimeLinkMaintained
	diff.MaxLinkLengthXDays = b.MaxLinkLengthXDays - a.MaxLinkLengthXDays
	diff.MaxTimeFieldHeld = b.MaxTimeFieldHeld - a.MaxTimeFieldHeld
	diff.LargestFieldMUsXDays = b.LargestFieldMUsXDays - a.LargestFieldMUsXDays

	diff.UniqueMissionsCompleted = b.UniqueMissionsCompleted - a.UniqueMissionsCompleted

	diff.AgentsSuccessfullyRecruited = b.AgentsSuccessfullyRecruited - a.AgentsSuccessfullyRecruited

	return diff
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Get(r, "tiers")
	userid, ok := session.Values["user"]

	templates := loadTemplates(
		"header.html",
		"footer.html",
		"nav.html",
		"index-unauthed.html",
		"profile.html",
	)

	if ok {
		var p profile.Profile
		var diff profile.Profile

		queue := model.GetNumQueuedProfiles(userid.(int))
		profiles := model.GetNewestProfiles(userid.(int), 2)

		switch len(profiles) {
		case 1:
			p = profiles[0]
		case 2:
			diff = createDiff(profiles[1], profiles[0])
			p = profiles[0]
		}

		templates.ExecuteTemplate(w, "profile", ProfilePage{
			User: userid.(int),
			Data: ProfilePageData{
				Profile: p,
				Diff:    diff,
				Queue:   queue,
			},
		})
	} else {
		templates.ExecuteTemplate(w, "index-unauthed", nil)
	}
}
