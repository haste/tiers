package profile

type Badge struct {
	Rank     int
	Current  uint
	Required uint
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

var BadgeRanks = map[string][]uint{
	"Connector":       {50, 1000, 5000, 25000, 100000},
	"Builder":         {2000, 10000, 30000, 100000, 200000},
	"Explorer":        {100, 1000, 2000, 10000, 30000},
	"Guardian":        {3, 10, 20, 90, 150},
	"Hacker":          {2000, 10000, 30000, 100000, 200000},
	"Mind Controller": {100, 500, 2000, 10000, 40000},
	"Purifier":        {2000, 10000, 30000, 100000, 300000},
	"Seer":            {10, 50, 200, 500, 5000},
	"Liberator":       {200, 2000, 8000, 15000, 40000},
	"Pioneer":         {20, 200, 1000, 5000, 20000},
	"Recharger":       {100000, 1000000, 3000000, 10000000, 25000000},
}

func incBadgeRank(b *Badge, current uint, reqs []uint) {
	for i := 0; i < len(reqs); i++ {
		req := reqs[i]
		if current <= req {
			b.Rank = i
			b.Current = current
			b.Required = req

			break
		}
	}
}

func countBadges(p Profile, b *Badges) {
	for k, v := range BadgeRanks {
		switch k {
		case "Connector":
			incBadgeRank(&b.Connector, p.LinksCreated, v)
		case "Builder":
			incBadgeRank(&b.Builder, p.ResonatorsDeployed, v)
		case "Explorer":
			incBadgeRank(&b.Explorer, p.UniquePortalsVisited, v)
		case "Guardian":
			incBadgeRank(&b.Guardian, p.MaxTimePortalHeld, v)
		case "Hacker":
			incBadgeRank(&b.Hacker, p.Hacks, v)
		case "Mind Controller":
			incBadgeRank(&b.MindController, p.ControlFieldsCreated, v)
		case "Purifier":
			incBadgeRank(&b.Purifier, p.ResonatorsDestroyed, v)
		case "Seer":
			incBadgeRank(&b.Seer, p.PortalsDiscovered, v)
		case "Liberator":
			incBadgeRank(&b.Liberator, p.PortalsCaptured, v)
		case "Pioneer":
			incBadgeRank(&b.Pioneer, p.UniquePortalsCaptured, v)
		case "Recharger":
			incBadgeRank(&b.Recharger, p.XMRecharged, v)
		}
	}
}

