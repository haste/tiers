package ocr

import (
	"testing"

	"tiers/conf"
	"tiers/profile"
)

var testData = map[string]profile.Profile{
	/*
		"": profile.Profile{
			Nick: "",
			Level:
			AP:

			UniquePortalsVisited:
			PortalsDiscovered:
			XMCollected:

			Hacks:
			ResonatorsDeployed:
			LinksCreated:
			ControlFieldsCreated:
			MindUnitsCaptured:
			LongestLinkEverCreated:
			LargestControlField:
			XMRecharged:
			PortalsCaptured:
			UniquePortalsCaptured:

			ResonatorsDestroyed:
			PortalsNeutralized:
			EnemyLinksDestroyed:
			EnemyControlFieldsDestroyed:

			DistanceWalked:

			MaxTimePortalHeld:
			MaxTimeLinkMaintained:
			MaxLinkLengthXDays:
			MaxTimeFieldHeld:
			LargestFieldMUsXDays:
		},
	*/
	"haste_v1620_nexus5.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    20463677,

		UniquePortalsVisited: 3414,
		PortalsDiscovered:    33,
		XMCollected:          93567145,

		Hacks:                  47171,
		ResonatorsDeployed:     45778,
		LinksCreated:           5046,
		ControlFieldsCreated:   2445,
		MindUnitsCaptured:      32598,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            24941649,
		PortalsCaptured:        5482,
		UniquePortalsCaptured:  1367,

		ResonatorsDestroyed:         39168,
		PortalsNeutralized:          5831,
		EnemyLinksDestroyed:         6945,
		EnemyControlFieldsDestroyed: 3115,

		DistanceWalked: 1824,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,
	},

	"wexp_v0_unknown.png": profile.Profile{
		Nick:  "wexp",
		Level: 8,
		AP:    1487296,

		UniquePortalsVisited: 259,
		PortalsDiscovered:    0,
		XMCollected:          7119691,

		Hacks:                  2755,
		ResonatorsDeployed:     2265,
		LinksCreated:           700,
		ControlFieldsCreated:   361,
		MindUnitsCaptured:      22959,
		LongestLinkEverCreated: 23,
		LargestControlField:    9875,
		XMRecharged:            3478701,
		PortalsCaptured:        194,
		UniquePortalsCaptured:  102,

		ResonatorsDestroyed:         1131,
		PortalsNeutralized:          170,
		EnemyLinksDestroyed:         277,
		EnemyControlFieldsDestroyed: 142,

		DistanceWalked: 157,

		MaxTimePortalHeld:     15,
		MaxTimeLinkMaintained: 9,
		MaxLinkLengthXDays:    38,
		MaxTimeFieldHeld:      9,
		LargestFieldMUsXDays:  9119,
	},
	"forferdet_v0_unknown.png": profile.Profile{
		Nick:  "forferdet",
		Level: 11,
		AP:    14322151,

		UniquePortalsVisited: 1441,
		PortalsDiscovered:    1,
		XMCollected:          61566079,

		Hacks:                  24298,
		ResonatorsDeployed:     25518,
		LinksCreated:           5351,
		ControlFieldsCreated:   2778,
		MindUnitsCaptured:      21328,
		LongestLinkEverCreated: 45,
		LargestControlField:    977,
		XMRecharged:            28106333,
		PortalsCaptured:        2694,
		UniquePortalsCaptured:  840,

		ResonatorsDestroyed:         19363,
		PortalsNeutralized:          2811,
		EnemyLinksDestroyed:         3707,
		EnemyControlFieldsDestroyed: 1746,

		DistanceWalked: 1008,

		MaxTimePortalHeld:     51,
		MaxTimeLinkMaintained: 55,
		MaxLinkLengthXDays:    93,
		MaxTimeFieldHeld:      52,
		LargestFieldMUsXDays:  2632,
	},

	"forferdet_v0_unknowntablet.png": profile.Profile{
		Nick:  "forferdet",
		Level: 11,
		AP:    12690292,

		UniquePortalsVisited: 1079,
		PortalsDiscovered:    0,
		XMCollected:          53584955,

		Hacks:                  21300,
		ResonatorsDeployed:     22618,
		LinksCreated:           5045,
		ControlFieldsCreated:   2626,
		MindUnitsCaptured:      18881,
		LongestLinkEverCreated: 45,
		LargestControlField:    692,
		XMRecharged:            24460830,
		PortalsCaptured:        2278,
		UniquePortalsCaptured:  706,

		ResonatorsDestroyed:         15806,
		PortalsNeutralized:          2297,
		EnemyLinksDestroyed:         3040,
		EnemyControlFieldsDestroyed: 1451,

		DistanceWalked: 925,

		MaxTimePortalHeld:     37,
		MaxTimeLinkMaintained: 55,
		MaxLinkLengthXDays:    93,
		MaxTimeFieldHeld:      52,
		LargestFieldMUsXDays:  540,
	},

	"erebwain_v1620_s4active.png": profile.Profile{
		Nick:  "erebwain",
		Level: 14,
		AP:    19617882,

		UniquePortalsVisited: 2378,
		PortalsDiscovered:    79,
		XMCollected:          70070561,

		Hacks:                  45889,
		ResonatorsDeployed:     40751,
		LinksCreated:           5259,
		ControlFieldsCreated:   2670,
		MindUnitsCaptured:      41709,
		LongestLinkEverCreated: 361,
		LargestControlField:    1227,
		XMRecharged:            16776081,
		PortalsCaptured:        4245,
		UniquePortalsCaptured:  1115,

		ResonatorsDestroyed:         31206,
		PortalsNeutralized:          3704,
		EnemyLinksDestroyed:         7496,
		EnemyControlFieldsDestroyed: 3855,

		DistanceWalked: 1730,

		MaxTimePortalHeld:     167,
		MaxTimeLinkMaintained: 28,
		MaxLinkLengthXDays:    442,
		MaxTimeFieldHeld:      27,
		LargestFieldMUsXDays:  1687,
	},

	"erebwain_v1630_unknown.png": profile.Profile{
		Nick:  "erebwain",
		Level: 14,
		AP:    20030145,

		UniquePortalsVisited: 2383,
		PortalsDiscovered:    79,
		XMCollected:          71673444,

		Hacks:                  46295,
		ResonatorsDeployed:     41287,
		LinksCreated:           5297,
		ControlFieldsCreated:   2696,
		MindUnitsCaptured:      41972,
		LongestLinkEverCreated: 361,
		LargestControlField:    1227,
		XMRecharged:            17363089,
		PortalsCaptured:        4309,
		UniquePortalsCaptured:  1118,

		ResonatorsDestroyed:         31645,
		PortalsNeutralized:          3760,
		EnemyLinksDestroyed:         7627,
		EnemyControlFieldsDestroyed: 3920,

		DistanceWalked: 1758,

		MaxTimePortalHeld:     167,
		MaxTimeLinkMaintained: 28,
		MaxLinkLengthXDays:    442,
		MaxTimeFieldHeld:      27,
		LargestFieldMUsXDays:  1687,
	},

	"zyp_v1630_unknown.png": profile.Profile{
		Nick:  "zyp",
		Level: 9,
		AP:    3092754,

		UniquePortalsVisited: 1146,
		PortalsDiscovered:    0,
		XMCollected:          8480812,

		Hacks:                  3818,
		ResonatorsDeployed:     4304,
		LinksCreated:           1271,
		ControlFieldsCreated:   605,
		MindUnitsCaptured:      79529,
		LongestLinkEverCreated: 172,
		LargestControlField:    24772,
		XMRecharged:            4137505,
		PortalsCaptured:        597,
		UniquePortalsCaptured:  413,

		ResonatorsDestroyed:         3719,
		PortalsNeutralized:          435,
		EnemyLinksDestroyed:         1030,
		EnemyControlFieldsDestroyed: 455,

		DistanceWalked: 179,

		MaxTimePortalHeld:     32,
		MaxTimeLinkMaintained: 25,
		MaxLinkLengthXDays:    180,
		MaxTimeFieldHeld:      15,
		LargestFieldMUsXDays:  3117,
	},

	"haste_v1630_nexus5.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    22186081,

		UniquePortalsVisited: 3423,
		PortalsDiscovered:    34,
		XMCollected:          97206086,

		Hacks:                  48188,
		ResonatorsDeployed:     47369,
		LinksCreated:           5082,
		ControlFieldsCreated:   2457,
		MindUnitsCaptured:      32617,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            25349179,
		PortalsCaptured:        5683,
		UniquePortalsCaptured:  1373,

		ResonatorsDestroyed:         41394,
		PortalsNeutralized:          6133,
		EnemyLinksDestroyed:         7590,
		EnemyControlFieldsDestroyed: 3480,

		DistanceWalked: 1865,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,
	},

	"haste_v1630_nexus5-1.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    22335897,

		UniquePortalsVisited: 3425,
		PortalsDiscovered:    34,
		XMCollected:          97603673,

		Hacks:                  48376,
		ResonatorsDeployed:     47538,
		LinksCreated:           5087,
		ControlFieldsCreated:   2458,
		MindUnitsCaptured:      32618,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            25377891,
		PortalsCaptured:        5690,
		UniquePortalsCaptured:  1373,

		ResonatorsDestroyed:         41496,
		PortalsNeutralized:          6145,
		EnemyLinksDestroyed:         7623,
		EnemyControlFieldsDestroyed: 3503,

		DistanceWalked: 1869,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,
	},
	"oteckeh_v1630_nexus4.png": profile.Profile{
		Nick:  "Oteckeh",
		Level: 10,
		AP:    5854511,

		UniquePortalsVisited: 2054,
		PortalsDiscovered:    5,
		XMCollected:          37956247,

		Hacks:                  15960,
		ResonatorsDeployed:     16929,
		LinksCreated:           1744,
		ControlFieldsCreated:   472,
		MindUnitsCaptured:      5369,
		LongestLinkEverCreated: 166,
		LargestControlField:    816,
		XMRecharged:            14595505,
		PortalsCaptured:        1201,
		UniquePortalsCaptured:  649,

		ResonatorsDestroyed:         10152,
		PortalsNeutralized:          1261,
		EnemyLinksDestroyed:         2197,
		EnemyControlFieldsDestroyed: 1026,

		DistanceWalked: 622,

		MaxTimePortalHeld:     202,
		MaxTimeLinkMaintained: 24,
		MaxLinkLengthXDays:    98,
		MaxTimeFieldHeld:      21,
		LargestFieldMUsXDays:  662,
	},

	"mvnch_v1640_iphone5.jpeg": profile.Profile{
		Nick:  "mvnch",
		Level: 8,
		AP:    1200134,

		UniquePortalsVisited: 137,
		PortalsDiscovered:    0,
		XMCollected:          1673726,

		Hacks:                  864,
		ResonatorsDeployed:     1208,
		LinksCreated:           392,
		ControlFieldsCreated:   230,
		MindUnitsCaptured:      701,
		LongestLinkEverCreated: 4,
		LargestControlField:    53,
		XMRecharged:            494542,
		PortalsCaptured:        114,
		UniquePortalsCaptured:  66,

		ResonatorsDestroyed:         534,
		PortalsNeutralized:          71,
		EnemyLinksDestroyed:         159,
		EnemyControlFieldsDestroyed: 80,

		DistanceWalked: 90,

		MaxTimePortalHeld:     20,
		MaxTimeLinkMaintained: 13,
		MaxLinkLengthXDays:    35,
		MaxTimeFieldHeld:      12,
		LargestFieldMUsXDays:  531,
	},

	"mvnch_v1640_iphone5-1.jpeg": profile.Profile{
		Nick:  "mvnch",
		Level: 8,
		AP:    1313096,

		UniquePortalsVisited: 138,
		PortalsDiscovered:    0,
		XMCollected:          2039828,

		Hacks:                  980,
		ResonatorsDeployed:     1361,
		LinksCreated:           450,
		ControlFieldsCreated:   267,
		MindUnitsCaptured:      749,
		LongestLinkEverCreated: 4,
		LargestControlField:    53,
		XMRecharged:            627128,
		PortalsCaptured:        127,
		UniquePortalsCaptured:  69,

		ResonatorsDestroyed:         640,
		PortalsNeutralized:          84,
		EnemyLinksDestroyed:         173,
		EnemyControlFieldsDestroyed: 85,

		DistanceWalked: 104,

		MaxTimePortalHeld:     20,
		MaxTimeLinkMaintained: 13,
		MaxLinkLengthXDays:    35,
		MaxTimeFieldHeld:      12,
		LargestFieldMUsXDays:  531,
	},
}

func init() {
	conf.Load("../config.json")
	conf.Config.Cache = "testdata/"
}

func validateOCR(t *testing.T, file string, p profile.Profile) {
	e := testData[file]

	if p.Nick != e.Nick {
		t.Errorf("%s: .Nick: Got %v Expected %v", file, p.Nick, e.Nick)
	}

	if p.Level != e.Level {
		t.Errorf("%s: .Level: Got %v Expected %v", file, p.Level, e.Level)
	}

	if p.AP != e.AP {
		t.Errorf("%s: .AP: Got %v Expected %v", file, p.AP, e.AP)
	}

	if p.UniquePortalsVisited != e.UniquePortalsVisited {
		t.Errorf("%s: .UniquePortalsVisited: Got %v Expected %v", file, p.UniquePortalsVisited, e.UniquePortalsVisited)
	}

	if p.PortalsDiscovered != e.PortalsDiscovered {
		t.Errorf("%s: .PortalsDiscovered: Got %v Expected %v", file, p.PortalsDiscovered, e.PortalsDiscovered)
	}

	if p.XMCollected != e.XMCollected {
		t.Errorf("%s: .XMCollected: Got %v Expected %v", file, p.XMCollected, e.XMCollected)
	}

	if p.Hacks != e.Hacks {
		t.Errorf("%s: .Hacks: Got %v Expected %v", file, p.Hacks, e.Hacks)
	}

	if p.ResonatorsDeployed != e.ResonatorsDeployed {
		t.Errorf("%s: .ResonatorsDeployed: Got %v Expected %v", file, p.ResonatorsDeployed, e.ResonatorsDeployed)
	}

	if p.LinksCreated != e.LinksCreated {
		t.Errorf("%s: .LinksCreated: Got %v Expected %v", file, p.LinksCreated, e.LinksCreated)
	}

	if p.ControlFieldsCreated != e.ControlFieldsCreated {
		t.Errorf("%s: .ControlFieldsCreated: Got %v Expected %v", file, p.ControlFieldsCreated, e.ControlFieldsCreated)
	}

	if p.MindUnitsCaptured != e.MindUnitsCaptured {
		t.Errorf("%s: .MindUnitsCaptured: Got %v Expected %v", file, p.MindUnitsCaptured, e.MindUnitsCaptured)
	}

	if p.LongestLinkEverCreated != e.LongestLinkEverCreated {
		t.Errorf("%s: .LongestLinkEverCreated: Got %v Expected %v", file, p.LongestLinkEverCreated, e.LongestLinkEverCreated)
	}

	if p.LargestControlField != e.LargestControlField {
		t.Errorf("%s: .LargestControlField: Got %v Expected %v", file, p.LargestControlField, e.LargestControlField)
	}

	if p.XMRecharged != e.XMRecharged {
		t.Errorf("%s: .XMRecharged: Got %v Expected %v", file, p.XMRecharged, e.XMRecharged)
	}

	if p.PortalsCaptured != e.PortalsCaptured {
		t.Errorf("%s: .PortalsCaptured: Got %v Expected %v", file, p.PortalsCaptured, e.PortalsCaptured)
	}

	if p.UniquePortalsCaptured != e.UniquePortalsCaptured {
		t.Errorf("%s: .UniquePortalsCaptured: Got %v Expected %v", file, p.UniquePortalsCaptured, e.UniquePortalsCaptured)
	}

	if p.ResonatorsDestroyed != e.ResonatorsDestroyed {
		t.Errorf("%s: .ResonatorsDestroyed: Got %v Expected %v", file, p.ResonatorsDestroyed, e.ResonatorsDestroyed)
	}

	if p.PortalsNeutralized != e.PortalsNeutralized {
		t.Errorf("%s: .PortalsNeutralized: Got %v Expected %v", file, p.PortalsNeutralized, e.PortalsNeutralized)
	}

	if p.EnemyLinksDestroyed != e.EnemyLinksDestroyed {
		t.Errorf("%s: .EnemyLinksDestroyed: Got %v Expected %v", file, p.EnemyLinksDestroyed, e.EnemyLinksDestroyed)
	}

	if p.EnemyControlFieldsDestroyed != e.EnemyControlFieldsDestroyed {
		t.Errorf("%s: .EnemyControlFieldsDestroyed: Got %v Expected %v", file, p.EnemyControlFieldsDestroyed, e.EnemyControlFieldsDestroyed)
	}

	if p.DistanceWalked != e.DistanceWalked {
		t.Errorf("%s: .DistanceWalked: Got %v Expected %v", file, p.DistanceWalked, e.DistanceWalked)
	}

	if p.MaxTimePortalHeld != e.MaxTimePortalHeld {
		t.Errorf("%s: .MaxTimePortalHeld: Got %v Expected %v", file, p.MaxTimePortalHeld, e.MaxTimePortalHeld)
	}

	if p.MaxTimeLinkMaintained != e.MaxTimeLinkMaintained {
		t.Errorf("%s: .MaxTimeLinkMaintained: Got %v Expected %v", file, p.MaxTimeLinkMaintained, e.MaxTimeLinkMaintained)
	}

	if p.MaxLinkLengthXDays != e.MaxLinkLengthXDays {
		t.Errorf("%s: .MaxLinkLengthXDays: Got %v Expected %v", file, p.MaxLinkLengthXDays, e.MaxLinkLengthXDays)
	}

	if p.MaxTimeFieldHeld != e.MaxTimeFieldHeld {
		t.Errorf("%s: .MaxTimeFieldHeld: Got %v Expected %v", file, p.MaxTimeFieldHeld, e.MaxTimeFieldHeld)
	}

	if p.LargestFieldMUsXDays != e.LargestFieldMUsXDays {
		t.Errorf("%s: .LargestFieldMUsXDays: Got %v Expected %v", file, p.LargestFieldMUsXDays, e.LargestFieldMUsXDays)
	}
}

func TestOCRhaste_v1620_nexus5(t *testing.T) {
	file := "haste_v1620_nexus5.png"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}

func TestOCRwexp_v0_unknown(t *testing.T) {
	file := "wexp_v0_unknown.png"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}

func TestOCRforferdet_v0_unknown(t *testing.T) {
	file := "forferdet_v0_unknown.png"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}

func TestOCRforferdet_v0_unknowntablet(t *testing.T) {
	file := "forferdet_v0_unknowntablet.png"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}

func TestOCRerebwain_v1620_s4active(t *testing.T) {
	file := "erebwain_v1620_s4active.png"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}

func TestOCRerebwain_v1630_unknown(t *testing.T) {
	file := "erebwain_v1630_unknown.png"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}

func TestOCRzyp_v1630_unknown(t *testing.T) {
	file := "zyp_v1630_unknown.png"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}

func TestOCRhaste_v1630_nexus5(t *testing.T) {
	file := "haste_v1630_nexus5.png"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}

func TestOCRhaste_v1630_nexus5_1(t *testing.T) {
	file := "haste_v1630_nexus5-1.png"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}

func TestOCRoteckeh_v1630_nexus4(t *testing.T) {
	file := "oteckeh_v1630_nexus4.png"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}

func TestOCRmvnch_v1640_iphone5(t *testing.T) {
	file := "mvnch_v1640_iphone5.jpeg"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}

func TestOCRmvnch_v1640_iphone5_1(t *testing.T) {
	file := "mvnch_v1640_iphone5-1.jpeg"
	res := runOCR(file)
	p := buildProfile(res)

	validateOCR(t, file, p)
}
