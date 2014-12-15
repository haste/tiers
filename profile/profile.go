package profile

type Profile struct {
	Id        int
	UserId    int
	Timestamp int64

	Nick  string
	Level int
	AP    int64

	NextLevel LevelRequirement

	Bronze   int
	Silver   int
	Gold     int
	Platinum int
	Onyx     int

	Badges Badges

	UniquePortalsVisited int64
	PortalsDiscovered    int64
	XMCollected          int64

	Hacks                  int64
	ResonatorsDeployed     int64
	LinksCreated           int64
	ControlFieldsCreated   int64
	MindUnitsCaptured      int64
	LongestLinkEverCreated int64
	LargestControlField    int64
	XMRecharged            int64
	PortalsCaptured        int64
	UniquePortalsCaptured  int64
	ModsDeployed           int64

	ResonatorsDestroyed         int64
	PortalsNeutralized          int64
	EnemyLinksDestroyed         int64
	EnemyControlFieldsDestroyed int64

	DistanceWalked int64

	MaxTimePortalHeld     int64
	MaxTimeLinkMaintained int64
	MaxLinkLengthXDays    int64
	MaxTimeFieldHeld      int64
	LargestFieldMUsXDays  int64

	UniqueMissionsCompleted int64

	InnovatorLevel int64
}

type LevelRequirement struct {
	Level uint
	AP    int64

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
	{12, 8400000, 0, 7, 6, 0, 0},
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

func HandleProfile(p *Profile) {
	HandleBadges(p)
	nextLevel(p)
}
