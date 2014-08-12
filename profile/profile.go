package profile

type Profile struct {
	Id int
	UserId int
	Timestamp uint

	Nick  string
	Level uint
	AP    uint

	NextLevel LevelRequirement

	Bronze   int
	Silver   int
	Gold     int
	Platinum int
	Onyx     int

	Badges Badges

	UniquePortalsVisited uint
	PortalsDiscovered    uint
	XMCollected          uint

	Hacks                  uint
	ResonatorsDeployed     uint
	LinksCreated           uint
	ControlFieldsCreated   uint
	MindUnitsCaptured      uint
	LongestLinkEverCreated uint
	LargestControlField    uint
	XMRecharged            uint
	PortalsCaptured        uint
	UniquePortalsCaptured  uint

	ResonatorsDestroyed         uint
	PortalsNeutralized          uint
	EnemyLinksDestroyed         uint
	EnemyControlFieldsDestroyed uint

	DistanceWalked uint

	MaxTimePortalHeld     uint
	MaxTimeLinkMaintained uint
	MaxLinkLengthXDays    uint
	MaxTimeFieldHeld      uint

	LargestFieldMUsXDays uint
}

type LevelRequirement struct {
	Level uint
	AP    uint

	Bronze   int
	Silver   int
	Gold     int
	Platinum int
	Onyx     int
}

var LevelRequirements = []LevelRequirement{
	{1, 0, 0, 0, 0, 0, 0},
	{2, 2500, 0, 0, 0, 0, 0},
	{3, 20000, 0, 0, 0, 0, 0},
	{4, 70000, 0, 0, 0, 0, 0},
	{5, 150000, 0, 0, 0, 0, 0},
	{6, 300000, 0, 0, 0, 0, 0},
	{7, 600000, 0, 0, 0, 0, 0},
	{8, 1200000, 0, 0, 0, 0, 0},
	{9, 2400000, 0, 4, 1, 0, 0},
	{10, 4000000, 0, 5, 2, 0, 0},
	{11, 6000000, 0, 6, 4, 0, 0},
	{12, 8400000,  0, 7, 6, 0, 0},
	{13, 12000000, 0, 0, 7, 1, 0},
	{14, 17000000, 0, 0, 7, 2, 0},
	{15, 24000000, 0, 0, 7, 3, 0},
	{16, 40000000, 0, 0, 7, 4, 2},
}

func findLevel(p *Profile) {
	for i := len(LevelRequirements) - 1; i >= 0; i-- {
		lr := LevelRequirements[i]
		if p.AP >= lr.AP &&
			p.Bronze >= lr.Bronze &&
			p.Silver >= lr.Silver &&
			p.Gold >= lr.Gold &&
			p.Platinum >= lr.Platinum &&
			p.Onyx >= p.Onyx {
			break
		}
	}
}

func nextLevel(p *Profile) {
	if p.Level < 16 {
		p.NextLevel = LevelRequirements[p.Level]
	}
}
