package profile

type Profile struct {
	Id int
	UserId int
	Timestamp uint

	Nick  string
	Level uint
	AP    uint

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
