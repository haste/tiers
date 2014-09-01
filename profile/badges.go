package profile

type Badge struct {
	Rank    int
	Ranks   []uint
	Current uint
	Next    uint
}

type Badges struct {
	Connector      Badge
	Builder        Badge
	Explorer       Badge
	Guardian       Badge
	Hacker         Badge
	MindController Badge
	Purifier       Badge
	Seer           Badge
	Liberator      Badge
	Pioneer        Badge
	Recharger      Badge
}

type BadgeProgress struct {
	Title    string `json:"title"`
	Ranges   []uint `json:"ranges"`
	Measures []uint `json:"measures"`
}

var BadgeRanks = map[string][]uint{
	"Connector":      {50, 1000, 5000, 25000, 100000},
	"Builder":        {2000, 10000, 30000, 100000, 200000},
	"Explorer":       {100, 1000, 2000, 10000, 30000},
	"Guardian":       {3, 10, 20, 90, 150},
	"Hacker":         {2000, 10000, 30000, 100000, 200000},
	"MindController": {100, 500, 2000, 10000, 40000},
	"Purifier":       {2000, 10000, 30000, 100000, 300000},
	"Seer":           {10, 50, 200, 500, 5000},
	"Liberator":      {200, 2000, 5000, 15000, 40000},
	"Pioneer":        {20, 200, 1000, 5000, 20000},
	"Recharger":      {100000, 1000000, 3000000, 10000000, 25000000},
}

// The order is defined by the agent profile.
var BadgeOrder = []string{
	"Explorer",
	"Seer",
	"Hacker",
	"Builder",
	"Connector",
	"Mind Controller",
	"Recharger",
	"Liberator",
	"Pioneer",
	"Purifier",
	"Guardian",
}

func incBadgeRank(p *Profile, b *Badge, current uint, reqs []uint) {
	for i := 0; i < len(reqs); i++ {
		req := reqs[i]
		b.Ranks = append(b.Ranks, req)

		if current >= req {
			switch i {
			case 0:
				p.Bronze++
			case 1:
				p.Silver++
			case 2:
				p.Gold++
			case 3:
				p.Platinum++
			case 4:
				p.Onyx++
			}
		} else {
			b.Rank = i
			b.Current = current
			b.Next = req

			break
		}
	}
}

func HandleBadges(p *Profile) {
	for k, v := range BadgeRanks {
		switch k {
		case "Connector":
			incBadgeRank(p, &p.Badges.Connector, p.LinksCreated, v)
		case "Builder":
			incBadgeRank(p, &p.Badges.Builder, p.ResonatorsDeployed, v)
		case "Explorer":
			incBadgeRank(p, &p.Badges.Explorer, p.UniquePortalsVisited, v)
		case "Guardian":
			incBadgeRank(p, &p.Badges.Guardian, p.MaxTimePortalHeld, v)
		case "Hacker":
			incBadgeRank(p, &p.Badges.Hacker, p.Hacks, v)
		case "MindController":
			incBadgeRank(p, &p.Badges.MindController, p.ControlFieldsCreated, v)
		case "Purifier":
			incBadgeRank(p, &p.Badges.Purifier, p.ResonatorsDestroyed, v)
		case "Seer":
			incBadgeRank(p, &p.Badges.Seer, p.PortalsDiscovered, v)
		case "Liberator":
			incBadgeRank(p, &p.Badges.Liberator, p.PortalsCaptured, v)
		case "Pioneer":
			incBadgeRank(p, &p.Badges.Pioneer, p.UniquePortalsCaptured, v)
		case "Recharger":
			incBadgeRank(p, &p.Badges.Recharger, p.XMRecharged, v)
		}
	}
}

func BuildBadgeProgress(p Profile) []BadgeProgress {
	var bp []BadgeProgress

	for _, bn := range BadgeOrder {
		var current Badge
		switch bn {
		case "Connector":
			current = p.Badges.Connector
		case "Builder":
			current = p.Badges.Builder
		case "Explorer":
			current = p.Badges.Explorer
		case "Guardian":
			current = p.Badges.Guardian
		case "Hacker":
			current = p.Badges.Hacker
		case "Mind Controller":
			current = p.Badges.MindController
		case "Purifier":
			current = p.Badges.Purifier
		case "Seer":
			current = p.Badges.Seer
		case "Liberator":
			current = p.Badges.Liberator
		case "Pioneer":
			current = p.Badges.Pioneer
		case "Recharger":
			current = p.Badges.Recharger
		}

		bp = append(bp, BadgeProgress{
			Title:    bn,
			Ranges:   current.Ranks,
			Measures: []uint{current.Current},
		})
	}

	return bp

}
