package ocr

import (
	"database/sql"
	"flag"
	"fmt"
	"testing"

	"tiers/conf"
	"tiers/model"
	"tiers/profile"
)

type userQueueEntry struct {
	Id          int
	File        string
	OCRProfile  int
	Processed   int
	Processtime int
	Profile     profile.Profile
}

var testData = map[string]profile.Profile{
	/*
		"": profile.Profile{
			Nick: "",
			Level:
			AP:

			UniquePortalsVisited:
			PortalsDiscovered:
			XMCollected:

			DistanceWalked:

			ResonatorsDeployed:
			LinksCreated:
			ControlFieldsCreated:
			MindUnitsCaptured:
			LongestLinkEverCreated:
			LargestControlField:
			XMRecharged:
			PortalsCaptured:
			UniquePortalsCaptured:
			ModsDeployed:

			ResonatorsDestroyed:
			PortalsNeutralized:
			EnemyLinksDestroyed:
			EnemyControlFieldsDestroyed:

			MaxTimePortalHeld:
			MaxTimeLinkMaintained:
			MaxLinkLengthXDays:
			MaxTimeFieldHeld:
			LargestFieldMUsXDays:

			UniqueMissionsCompleted:

			Hacks:
			GlyphHackPoints:
			ConsecutiveDaysHacking:

			AgentsSuccessfullyRecruited:

			InnovatorLevel:
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

		InnovatorLevel: 3,
	},

	"mrwolfe_v1640_unknown.png": profile.Profile{
		Nick:  "MrWolfe",
		Level: 14,
		AP:    17983405,

		UniquePortalsVisited: 2002,
		PortalsDiscovered:    167,
		XMCollected:          67535405,

		Hacks:                  24443,
		ResonatorsDeployed:     28948,
		LinksCreated:           5383,
		ControlFieldsCreated:   2393,
		MindUnitsCaptured:      1122111,
		LongestLinkEverCreated: 435,
		LargestControlField:    119540,
		XMRecharged:            28111419,
		PortalsCaptured:        3518,
		UniquePortalsCaptured:  1022,

		ResonatorsDestroyed:         32967,
		PortalsNeutralized:          3892,
		EnemyLinksDestroyed:         7824,
		EnemyControlFieldsDestroyed: 3958,

		DistanceWalked: 826,

		MaxTimePortalHeld:     75,
		MaxTimeLinkMaintained: 64,
		MaxLinkLengthXDays:    1853,
		MaxTimeFieldHeld:      64,
		LargestFieldMUsXDays:  56215,

		UniqueMissionsCompleted: 1,

		InnovatorLevel: 13,
	},

	"haste_v1660_nexus5.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    24617282,

		UniquePortalsVisited: 3490,
		PortalsDiscovered:    36,
		XMCollected:          103004717,

		Hacks:                  50428,
		ResonatorsDeployed:     51628,
		LinksCreated:           5172,
		ControlFieldsCreated:   2502,
		MindUnitsCaptured:      32748,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26279823,
		PortalsCaptured:        6086,
		UniquePortalsCaptured:  1405,

		ResonatorsDestroyed:         44043,
		PortalsNeutralized:          6521,
		EnemyLinksDestroyed:         8201,
		EnemyControlFieldsDestroyed: 3813,

		DistanceWalked: 1965,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 0,

		InnovatorLevel: 13,
	},

	"haste_v1660_nexus5-1.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    25042616,

		UniquePortalsVisited: 3493,
		PortalsDiscovered:    36,
		XMCollected:          104041230,

		DistanceWalked: 1993,

		ResonatorsDeployed:     52542,
		LinksCreated:           5216,
		ControlFieldsCreated:   2527,
		MindUnitsCaptured:      32777,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26357201,
		PortalsCaptured:        6168,
		UniquePortalsCaptured:  1406,
		ModsDeployed:           179,

		ResonatorsDestroyed:         44459,
		PortalsNeutralized:          6587,
		EnemyLinksDestroyed:         8314,
		EnemyControlFieldsDestroyed: 3888,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 51113,

		UniqueMissionsCompleted: 0,

		InnovatorLevel: 13,
	},

	"zyp_v0_unknown.png": profile.Profile{
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
		ModsDeployed:           0,

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

		UniqueMissionsCompleted: 0,

		InnovatorLevel: 0,
	},
	"viddy_v1640_nexus7_2013.png": profile.Profile{
		Nick:  "viddy",
		Level: 14,
		AP:    20592365,

		UniquePortalsVisited: 3033,
		PortalsDiscovered:    15,
		XMCollected:          103966957,

		Hacks:                  39544,
		ResonatorsDeployed:     43736,
		LinksCreated:           3531,
		ControlFieldsCreated:   1506,
		MindUnitsCaptured:      15251,
		LongestLinkEverCreated: 166,
		LargestControlField:    1250,
		XMRecharged:            31730665,
		PortalsCaptured:        6460,
		UniquePortalsCaptured:  1482,
		ModsDeployed:           0,

		ResonatorsDestroyed:         42377,
		PortalsNeutralized:          5973,
		EnemyLinksDestroyed:         7120,
		EnemyControlFieldsDestroyed: 3426,

		DistanceWalked: 1418,

		MaxTimePortalHeld:     218,
		MaxTimeLinkMaintained: 67,
		MaxLinkLengthXDays:    98,
		MaxTimeFieldHeld:      64,
		LargestFieldMUsXDays:  5153,

		UniqueMissionsCompleted: 0,

		InnovatorLevel: 13,
	},

	"sockerdricka_v0_unknown.png": profile.Profile{
		Nick:  "sockerdricka",
		Level: 10,
		AP:    6470507,

		UniquePortalsVisited: 704,
		PortalsDiscovered:    26,
		XMCollected:          27964186,

		Hacks:                  10174,
		ResonatorsDeployed:     10403,
		LinksCreated:           2276,
		ControlFieldsCreated:   1160,
		MindUnitsCaptured:      35381,
		LongestLinkEverCreated: 198,
		LargestControlField:    12191,
		XMRecharged:            12265545,
		PortalsCaptured:        1008,
		UniquePortalsCaptured:  335,
		ModsDeployed:           0,

		ResonatorsDestroyed:         6530,
		PortalsNeutralized:          849,
		EnemyLinksDestroyed:         1693,
		EnemyControlFieldsDestroyed: 912,

		DistanceWalked: 472,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 77,
		MaxLinkLengthXDays:    11298,
		MaxTimeFieldHeld:      54,
		LargestFieldMUsXDays:  659702,

		UniqueMissionsCompleted: 0,

		InnovatorLevel: 0,
	},

	"scissorhill_v0_iphone.jpeg": profile.Profile{
		Nick:  "Scissorhill",
		Level: 10,
		AP:    7980258,

		UniquePortalsVisited: 714,
		PortalsDiscovered:    0,
		XMCollected:          22655119,

		DistanceWalked: 729,

		ResonatorsDeployed:     9838,
		LinksCreated:           2760,
		ControlFieldsCreated:   1540,
		MindUnitsCaptured:      179031,
		LongestLinkEverCreated: 14,
		LargestControlField:    10243,
		XMRecharged:            9493188,
		PortalsCaptured:        1047,
		UniquePortalsCaptured:  328,
		ModsDeployed:           187,

		ResonatorsDestroyed:         7128,
		PortalsNeutralized:          919,
		EnemyLinksDestroyed:         1561,
		EnemyControlFieldsDestroyed: 716,

		MaxTimePortalHeld:     51,
		MaxTimeLinkMaintained: 37,
		MaxLinkLengthXDays:    43,
		MaxTimeFieldHeld:      25,
		LargestFieldMUsXDays:  17837,

		UniqueMissionsCompleted: 1,

		Hacks: 6796,

		InnovatorLevel: 9,
	},

	"tufte_v1660_iphone.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    4001715,

		UniquePortalsVisited: 970,
		PortalsDiscovered:    0,
		XMCollected:          13809076,

		DistanceWalked: 340,

		ResonatorsDeployed:     5323,
		LinksCreated:           968,
		ControlFieldsCreated:   576,
		MindUnitsCaptured:      6704,
		LongestLinkEverCreated: 6,
		LargestControlField:    838,
		XMRecharged:            6035130,
		PortalsCaptured:        467,
		UniquePortalsCaptured:  254,
		ModsDeployed:           249,

		ResonatorsDestroyed:         3387,
		PortalsNeutralized:          405,
		EnemyLinksDestroyed:         774,
		EnemyControlFieldsDestroyed: 413,

		MaxTimePortalHeld:     22,
		MaxTimeLinkMaintained: 22,
		MaxLinkLengthXDays:    7,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  518,

		Hacks: 4312,

		UniqueMissionsCompleted: 0,

		InnovatorLevel: 3,
	},

	"haste_v1670_nexus5.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    25554981,

		UniquePortalsVisited: 3515,
		PortalsDiscovered:    36,
		XMCollected:          106750398,

		DistanceWalked: 2021,

		ResonatorsDeployed:     53809,
		LinksCreated:           5264,
		ControlFieldsCreated:   2546,
		MindUnitsCaptured:      32799,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26504289,
		PortalsCaptured:        6247,
		UniquePortalsCaptured:  1408,
		ModsDeployed:           399,

		ResonatorsDestroyed:         45071,
		PortalsNeutralized:          6671,
		EnemyLinksDestroyed:         8490,
		EnemyControlFieldsDestroyed: 3992,

		MaxTimePortalHeld:     113,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 52367,

		UniqueMissionsCompleted: 0,

		InnovatorLevel: 13,
	},

	"tufte_v1670_iphone.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    4347932,

		UniquePortalsVisited: 1067,
		PortalsDiscovered:    0,
		XMCollected:          16119070,

		DistanceWalked: 379,

		ResonatorsDeployed:     6105,
		LinksCreated:           1059,
		ControlFieldsCreated:   647,
		MindUnitsCaptured:      7440,
		LongestLinkEverCreated: 6,
		LargestControlField:    838,
		XMRecharged:            7252265,
		PortalsCaptured:        496,
		UniquePortalsCaptured:  264,
		ModsDeployed:           384,

		ResonatorsDestroyed:         3723,
		PortalsNeutralized:          443,
		EnemyLinksDestroyed:         861,
		EnemyControlFieldsDestroyed: 456,

		MaxTimePortalHeld:     27,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    7,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  518,

		Hacks: 4915,

		UniqueMissionsCompleted: 0,

		InnovatorLevel: 3,
	},

	"haste_v1680_nexus6.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    25964012,

		UniquePortalsVisited: 3525,
		PortalsDiscovered:    37,
		XMCollected:          108520560,

		DistanceWalked: 2044,

		ResonatorsDeployed:     54622,
		LinksCreated:           5318,
		ControlFieldsCreated:   2568,
		MindUnitsCaptured:      32856,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26524437,
		PortalsCaptured:        6308,
		UniquePortalsCaptured:  1415,
		ModsDeployed:           594,

		ResonatorsDestroyed:         45557,
		PortalsNeutralized:          6739,
		EnemyLinksDestroyed:         8707,
		EnemyControlFieldsDestroyed: 4150,

		MaxTimePortalHeld:     122,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 52997,

		UniqueMissionsCompleted: 0,

		InnovatorLevel: 13,
	},

	"haste_v1680_nexus6-1.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    25965522,

		UniquePortalsVisited: 3541,
		PortalsDiscovered:    37,
		XMCollected:          108540094,

		DistanceWalked: 2045,

		ResonatorsDeployed:     54626,
		LinksCreated:           5318,
		ControlFieldsCreated:   2568,
		MindUnitsCaptured:      32856,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26527164,
		PortalsCaptured:        6308,
		UniquePortalsCaptured:  1415,
		ModsDeployed:           596,

		ResonatorsDestroyed:         45557,
		PortalsNeutralized:          6739,
		EnemyLinksDestroyed:         8707,
		EnemyControlFieldsDestroyed: 4150,

		MaxTimePortalHeld:     123,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 53033,

		InnovatorLevel: 13,
	},

	"forferdet_v1680_oneplusone.png": profile.Profile{
		Nick:  "forferdet",
		Level: 14,
		AP:    25137484,

		UniquePortalsVisited: 1746,
		PortalsDiscovered:    13,
		XMCollected:          97771074,

		DistanceWalked: 1372,

		ResonatorsDeployed:     41586,
		LinksCreated:           7113,
		ControlFieldsCreated:   3459,
		MindUnitsCaptured:      24251,
		LongestLinkEverCreated: 156,
		LargestControlField:    977,
		XMRecharged:            33691418,
		PortalsCaptured:        4414,
		UniquePortalsCaptured:  1049,
		ModsDeployed:           1333,

		ResonatorsDestroyed:         35730,
		PortalsNeutralized:          4935,
		EnemyLinksDestroyed:         7296,
		EnemyControlFieldsDestroyed: 3627,

		MaxTimePortalHeld:     53,
		MaxTimeLinkMaintained: 55,
		MaxLinkLengthXDays:    2543,
		MaxTimeFieldHeld:      52,
		LargestFieldMUsXDays:  2632,

		Hacks: 37532,

		UniqueMissionsCompleted: 0,

		InnovatorLevel: 11,
	},

	"madder79_v1680_unknown.png": profile.Profile{
		Nick:  "madder79",
		Level: 10,
		AP:    7334883,

		UniquePortalsVisited: 1102,
		PortalsDiscovered:    4,
		XMCollected:          25527106,

		DistanceWalked: 418,

		ResonatorsDeployed:     10419,
		LinksCreated:           2137,
		ControlFieldsCreated:   1081,
		MindUnitsCaptured:      18187,
		LongestLinkEverCreated: 37,
		LargestControlField:    803,
		XMRecharged:            6500245,
		PortalsCaptured:        887,
		UniquePortalsCaptured:  351,
		ModsDeployed:           481,

		ResonatorsDestroyed:         10412,
		PortalsNeutralized:          1345,
		EnemyLinksDestroyed:         2169,
		EnemyControlFieldsDestroyed: 1132,

		MaxTimePortalHeld:     20,
		MaxTimeLinkMaintained: 21,
		MaxLinkLengthXDays:    199,
		MaxTimeFieldHeld:      13,
		LargestFieldMUsXDays:  2624,

		UniqueMissionsCompleted: 3,

		Hacks: 9533,

		AgentsSuccessfullyRecruited: 1,

		InnovatorLevel: 9,
	},

	"haste_v1680_nexus6-2.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    26526758,

		UniquePortalsVisited: 3564,
		PortalsDiscovered:    37,
		XMCollected:          110765815,

		DistanceWalked: 2084,

		ResonatorsDeployed:     56074,
		LinksCreated:           5362,
		ControlFieldsCreated:   2586,
		MindUnitsCaptured:      32880,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26547006,
		PortalsCaptured:        6459,
		UniquePortalsCaptured:  1432,
		ModsDeployed:           761,

		ResonatorsDestroyed:         46454,
		PortalsNeutralized:          6886,
		EnemyLinksDestroyed:         8944,
		EnemyControlFieldsDestroyed: 4296,

		MaxTimePortalHeld:     134,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 0,

		Hacks: 53963,

		AgentsSuccessfullyRecruited: 0,

		InnovatorLevel: 13,
	},

	"haste_v1690_nexus6.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    26601022,

		UniquePortalsVisited: 3567,
		PortalsDiscovered:    39,
		XMCollected:          111451742,

		DistanceWalked: 2093,

		ResonatorsDeployed:     56352,
		LinksCreated:           5374,
		ControlFieldsCreated:   2593,
		MindUnitsCaptured:      32903,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26907756,
		PortalsCaptured:        6471,
		UniquePortalsCaptured:  1433,
		ModsDeployed:           810,

		ResonatorsDestroyed:         46508,
		PortalsNeutralized:          6893,
		EnemyLinksDestroyed:         8947,
		EnemyControlFieldsDestroyed: 4296,

		MaxTimePortalHeld:     136,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 54204,

		InnovatorLevel: 13,
	},

	"haste_v1690_nexus6-1.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    26706948,

		UniquePortalsVisited: 3569,
		PortalsDiscovered:    39,
		XMCollected:          111894104,

		DistanceWalked: 2097,

		ResonatorsDeployed:     56630,
		LinksCreated:           5374,
		ControlFieldsCreated:   2593,
		MindUnitsCaptured:      32903,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            27039952,
		PortalsCaptured:        6498,
		UniquePortalsCaptured:  1435,
		ModsDeployed:           819,

		ResonatorsDestroyed:         46737,
		PortalsNeutralized:          6921,
		EnemyLinksDestroyed:         9005,
		EnemyControlFieldsDestroyed: 4324,

		MaxTimePortalHeld:     137,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks:           54287,
		GlyphHackPoints: 9,

		InnovatorLevel: 13,
	},

	"zyp_v1690_unknown.png": profile.Profile{
		Nick:  "zyp",
		Level: 9,
		AP:    3625188,

		UniquePortalsVisited: 1649,
		PortalsDiscovered:    0,
		XMCollected:          10863728,

		DistanceWalked: 223,

		ResonatorsDeployed:     5111,
		LinksCreated:           1302,
		ControlFieldsCreated:   617,
		MindUnitsCaptured:      79600,
		LongestLinkEverCreated: 172,
		LargestControlField:    24772,
		XMRecharged:            5129754,
		PortalsCaptured:        820,
		UniquePortalsCaptured:  608,
		ModsDeployed:           151,

		ResonatorsDestroyed:         4949,
		PortalsNeutralized:          650,
		EnemyLinksDestroyed:         1303,
		EnemyControlFieldsDestroyed: 589,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 25,
		MaxLinkLengthXDays:    180,
		MaxTimeFieldHeld:      15,
		LargestFieldMUsXDays:  3117,

		UniqueMissionsCompleted: 5,

		Hacks:           4488,
		GlyphHackPoints: 8,

		InnovatorLevel: 9,
	},

	"zyp_v1690_unknown-1.png": profile.Profile{
		Nick:  "zyp",
		Level: 9,
		AP:    3671555,

		UniquePortalsVisited: 1810,
		XMCollected:          10965815,

		DistanceWalked: 234,

		ResonatorsDeployed:     5188,
		LinksCreated:           1314,
		ControlFieldsCreated:   626,
		MindUnitsCaptured:      79830,
		LongestLinkEverCreated: 172,
		LargestControlField:    24772,
		XMRecharged:            5130754,
		PortalsCaptured:        828,
		UniquePortalsCaptured:  616,
		ModsDeployed:           156,

		ResonatorsDestroyed:         5009,
		PortalsNeutralized:          656,
		EnemyLinksDestroyed:         1316,
		EnemyControlFieldsDestroyed: 595,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 25,
		MaxLinkLengthXDays:    180,
		MaxTimeFieldHeld:      15,
		LargestFieldMUsXDays:  3117,

		UniqueMissionsCompleted: 5,

		Hacks:           4660,
		GlyphHackPoints: 8,

		InnovatorLevel: 9,
	},

	"oteckeh_v1690_nexus4.png": profile.Profile{
		Nick:  "Oteckeh",
		Level: 11,
		AP:    7288498,

		UniquePortalsVisited: 2190,
		PortalsDiscovered:    5,
		XMCollected:          42660034,

		DistanceWalked: 690,

		ResonatorsDeployed:     19354,
		LinksCreated:           1857,
		ControlFieldsCreated:   542,
		MindUnitsCaptured:      5620,
		LongestLinkEverCreated: 166,
		LargestControlField:    816,
		XMRecharged:            14968211,
		PortalsCaptured:        1494,
		UniquePortalsCaptured:  700,
		ModsDeployed:           195,

		ResonatorsDestroyed:         13038,
		PortalsNeutralized:          1652,
		EnemyLinksDestroyed:         2979,
		EnemyControlFieldsDestroyed: 1478,

		MaxTimePortalHeld:     220,
		MaxTimeLinkMaintained: 24,
		MaxLinkLengthXDays:    98,
		MaxTimeFieldHeld:      21,
		LargestFieldMUsXDays:  662,

		Hacks: 17775,

		InnovatorLevel: 9,
	},

	"forferdet_v1700_oneplusone.png": profile.Profile{
		Nick:  "forferdet",
		Level: 14,
		AP:    27199137,

		UniquePortalsVisited: 1759,
		PortalsDiscovered:    13,
		XMCollected:          107292028,

		DistanceWalked: 1471,

		ResonatorsDeployed:     45165,
		LinksCreated:           7569,
		ControlFieldsCreated:   3648,
		MindUnitsCaptured:      25072,
		LongestLinkEverCreated: 156,
		LargestControlField:    977,
		XMRecharged:            35797326,
		PortalsCaptured:        4810,
		UniquePortalsCaptured:  1062,
		ModsDeployed:           2000,

		ResonatorsDestroyed:         39521,
		PortalsNeutralized:          5475,
		EnemyLinksDestroyed:         8171,
		EnemyControlFieldsDestroyed: 4070,

		MaxTimePortalHeld:     53,
		MaxTimeLinkMaintained: 55,
		MaxLinkLengthXDays:    2543,
		MaxTimeFieldHeld:      52,
		LargestFieldMUsXDays:  2632,

		Hacks:           40723,
		GlyphHackPoints: 11472,

		InnovatorLevel: 11,
	},
	"scissorhill_v0_ipad.jpeg": profile.Profile{
		Nick:  "Scissorhill",
		Level: 10,
		AP:    7980258,

		UniquePortalsVisited: 714,
		XMCollected:          22655119,

		DistanceWalked: 729,

		ResonatorsDeployed:     9838,
		LinksCreated:           2760,
		ControlFieldsCreated:   1540,
		MindUnitsCaptured:      179031,
		LongestLinkEverCreated: 14,
		LargestControlField:    10243,
		XMRecharged:            9493188,
		PortalsCaptured:        1047,
		UniquePortalsCaptured:  328,
		ModsDeployed:           187,

		ResonatorsDestroyed:         7128,
		PortalsNeutralized:          919,
		EnemyLinksDestroyed:         1561,
		EnemyControlFieldsDestroyed: 716,

		MaxTimePortalHeld:     51,
		MaxTimeLinkMaintained: 37,
		MaxLinkLengthXDays:    43,
		MaxTimeFieldHeld:      25,
		LargestFieldMUsXDays:  17837,

		UniqueMissionsCompleted: 1,

		Hacks: 6796,

		InnovatorLevel: 9,
	},

	"tufte_v1700_iphone.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    7024572,

		UniquePortalsVisited: 1461,
		XMCollected:          32848892,

		DistanceWalked: 586,

		ResonatorsDeployed:     11020,
		LinksCreated:           1693,
		ControlFieldsCreated:   1024,
		MindUnitsCaptured:      623080,
		LongestLinkEverCreated: 304,
		LargestControlField:    610877,
		XMRecharged:            15618577,
		PortalsCaptured:        932,
		UniquePortalsCaptured:  470,
		ModsDeployed:           1099,

		ResonatorsDestroyed:         7667,
		PortalsNeutralized:          916,
		EnemyLinksDestroyed:         1840,
		EnemyControlFieldsDestroyed: 1025,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    274,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  354224,

		UniqueMissionsCompleted: 2,

		Hacks:           8335,
		GlyphHackPoints: 713,

		InnovatorLevel: 3,
	},

	"tufte_v1700_iphone-1.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    7093540,

		UniquePortalsVisited: 1461,
		XMCollected:          33088719,

		DistanceWalked: 588,

		ResonatorsDeployed:     11135,
		LinksCreated:           1717,
		ControlFieldsCreated:   1036,
		MindUnitsCaptured:      623112,
		LongestLinkEverCreated: 304,
		LargestControlField:    610877,
		XMRecharged:            15695063,
		PortalsCaptured:        945,
		UniquePortalsCaptured:  471,
		ModsDeployed:           1116,

		ResonatorsDestroyed:         7750,
		PortalsNeutralized:          926,
		EnemyLinksDestroyed:         1863,
		EnemyControlFieldsDestroyed: 1034,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    274,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  354224,

		UniqueMissionsCompleted: 2,

		Hacks:           8400,
		GlyphHackPoints: 721,

		InnovatorLevel: 3,
	},

	"zyp_v1700_unknown.png": profile.Profile{
		Nick:  "zyp",
		Level: 9,
		AP:    3931449,

		UniquePortalsVisited: 2016,
		XMCollected:          11476134,

		DistanceWalked: 247,

		ResonatorsDeployed:     5631,
		LinksCreated:           1350,
		ControlFieldsCreated:   640,
		MindUnitsCaptured:      80098,
		LongestLinkEverCreated: 172,
		LargestControlField:    24772,
		XMRecharged:            5216613,
		PortalsCaptured:        894,
		UniquePortalsCaptured:  664,
		ModsDeployed:           207,

		ResonatorsDestroyed:         5540,
		PortalsNeutralized:          722,
		EnemyLinksDestroyed:         1448,
		EnemyControlFieldsDestroyed: 654,

		MaxTimePortalHeld:     52,
		MaxTimeLinkMaintained: 25,
		MaxLinkLengthXDays:    180,
		MaxTimeFieldHeld:      15,
		LargestFieldMUsXDays:  3117,

		UniqueMissionsCompleted: 9,

		Hacks:           4919,
		GlyphHackPoints: 44,

		InnovatorLevel: 9,
	},

	"haste_v1711_nexus6.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    29629843,

		UniquePortalsVisited: 3805,
		PortalsDiscovered:    40,
		XMCollected:          123400576,

		DistanceWalked: 2258,

		ResonatorsDeployed:     63128,
		LinksCreated:           5723,
		ControlFieldsCreated:   2780,
		MindUnitsCaptured:      33453,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            27698152,
		PortalsCaptured:        7092,
		UniquePortalsCaptured:  1523,
		ModsDeployed:           1502,

		ResonatorsDestroyed:         51499,
		PortalsNeutralized:          7566,
		EnemyLinksDestroyed:         10347,
		EnemyControlFieldsDestroyed: 5087,

		MaxTimePortalHeld:     163,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 16,

		Hacks:           57908,
		GlyphHackPoints: 8637,
	},

	"zyp_v1711_unknown.png": profile.Profile{
		Nick:  "zyp",
		Level: 10,
		AP:    4012418,

		UniquePortalsVisited: 2022,
		XMCollected:          11632244,

		DistanceWalked: 250,

		ResonatorsDeployed:     5751,
		LinksCreated:           1362,
		ControlFieldsCreated:   646,
		MindUnitsCaptured:      80104,
		LongestLinkEverCreated: 172,
		LargestControlField:    24772,
		XMRecharged:            5232639,
		PortalsCaptured:        909,
		UniquePortalsCaptured:  669,
		ModsDeployed:           208,

		ResonatorsDestroyed:         5629,
		PortalsNeutralized:          733,
		EnemyLinksDestroyed:         1497,
		EnemyControlFieldsDestroyed: 689,

		MaxTimePortalHeld:     55,
		MaxTimeLinkMaintained: 25,
		MaxLinkLengthXDays:    180,
		MaxTimeFieldHeld:      15,
		LargestFieldMUsXDays:  3117,

		UniqueMissionsCompleted: 9,

		Hacks:           4948,
		GlyphHackPoints: 48,

		InnovatorLevel: 9,
	},

	"leiasqz_v1711_unknown.png": profile.Profile{
		Nick:  "leiasqz",
		Level: 8,
		AP:    1705304,

		UniquePortalsVisited: 295,
		XMCollected:          1593907,

		DistanceWalked: 24,

		ResonatorsDeployed:     4245,
		LinksCreated:           1841,
		ControlFieldsCreated:   229,
		MindUnitsCaptured:      15415,
		LongestLinkEverCreated: 19,
		LargestControlField:    3705,
		XMRecharged:            465584,
		PortalsCaptured:        364,
		UniquePortalsCaptured:  138,
		ModsDeployed:           114,

		ResonatorsDestroyed:         3321,
		PortalsNeutralized:          110,
		EnemyLinksDestroyed:         779,
		EnemyControlFieldsDestroyed: 192,

		MaxTimePortalHeld:     5,
		MaxTimeLinkMaintained: 4,
		MaxLinkLengthXDays:    1,
		MaxTimeFieldHeld:      4,
		LargestFieldMUsXDays:  77,

		Hacks:           4034,
		GlyphHackPoints: 62,

		InnovatorLevel: 3,
	},

	"mvnch_vunknown_iphone5.jpeg": profile.Profile{
		Nick:  "mvnch",
		Level: 5,
		AP:    262843,

		UniquePortalsVisited: 123,
		XMCollected:          410381,

		DistanceWalked: 35,

		ResonatorsDeployed:     336,
		LinksCreated:           153,
		ControlFieldsCreated:   105,
		MindUnitsCaptured:      408,
		LongestLinkEverCreated: 4,
		LargestControlField:    53,
		XMRecharged:            181304,
		PortalsCaptured:        28,
		UniquePortalsCaptured:  25,

		ResonatorsDestroyed:         65,
		PortalsNeutralized:          8,
		EnemyLinksDestroyed:         12,
		EnemyControlFieldsDestroyed: 2,

		MaxTimePortalHeld:     5,
		MaxTimeLinkMaintained: 5,
		MaxLinkLengthXDays:    11,
		MaxTimeFieldHeld:      5,
		LargestFieldMUsXDays:  36,

		Hacks: 365,
	},

	"mvnch_vunknown_iphone5-1.jpeg": profile.Profile{
		Nick:  "mvnch",
		Level: 3,
		AP:    22187,

		UniquePortalsVisited: 27,
		XMCollected:          28348,

		DistanceWalked: 6,

		ResonatorsDeployed:     29,
		LinksCreated:           14,
		ControlFieldsCreated:   9,
		MindUnitsCaptured:      17,
		LongestLinkEverCreated: 1,
		LargestControlField:    4,
		XMRecharged:            2000,
		PortalsCaptured:        1,
		UniquePortalsCaptured:  1,

		MaxTimePortalHeld:     1,
		MaxTimeLinkMaintained: 1,
		MaxLinkLengthXDays:    2,
		MaxTimeFieldHeld:      1,
		LargestFieldMUsXDays:  7,

		Hacks: 45,
	},

	"steino_v0_unknown.png": profile.Profile{
		Nick:  "steino",
		Level: 8,
		AP:    1391847,

		UniquePortalsVisited: 586,
		XMCollected:          3247773,

		DistanceWalked: 145,

		ResonatorsDeployed:     2836,
		LinksCreated:           638,
		ControlFieldsCreated:   331,
		MindUnitsCaptured:      10650,
		LongestLinkEverCreated: 3,
		LargestControlField:    378,
		XMRecharged:            315312,
		PortalsCaptured:        240,
		UniquePortalsCaptured:  71,

		ResonatorsDestroyed:         1367,
		PortalsNeutralized:          226,
		EnemyLinksDestroyed:         207,
		EnemyControlFieldsDestroyed: 96,

		MaxTimePortalHeld:     11,
		MaxTimeLinkMaintained: 11,
		MaxLinkLengthXDays:    5,
		MaxTimeFieldHeld:      6,
		LargestFieldMUsXDays:  760,

		Hacks: 2444,
	},

	"wexp_v0_unknown-1.png": profile.Profile{
		Nick:  "wexp",
		Level: 8,
		AP:    1206693,

		UniquePortalsVisited: 176,
		XMCollected:          4884617,

		DistanceWalked: 123,

		ResonatorsDeployed:     1899,
		LinksCreated:           608,
		ControlFieldsCreated:   305,
		MindUnitsCaptured:      511,
		LongestLinkEverCreated: 20,
		LargestControlField:    59,
		XMRecharged:            1879663,
		PortalsCaptured:        162,
		UniquePortalsCaptured:  77,

		ResonatorsDestroyed:         869,
		PortalsNeutralized:          124,
		EnemyLinksDestroyed:         187,
		EnemyControlFieldsDestroyed: 90,

		MaxTimePortalHeld:     8,
		MaxTimeLinkMaintained: 5,
		MaxLinkLengthXDays:    38,
		MaxTimeFieldHeld:      4,
		LargestFieldMUsXDays:  35,

		Hacks: 2268,
	},

	"zyp_v0_unknown-1.png": profile.Profile{
		Nick:  "zyp",
		Level: 9,
		AP:    2538655,

		UniquePortalsVisited: 855,
		XMCollected:          7097102,

		DistanceWalked: 148,

		ResonatorsDeployed:     3778,
		LinksCreated:           1169,
		ControlFieldsCreated:   556,
		MindUnitsCaptured:      79252,
		LongestLinkEverCreated: 167,
		LargestControlField:    24772,
		XMRecharged:            3462885,
		PortalsCaptured:        526,
		UniquePortalsCaptured:  365,

		ResonatorsDestroyed:         3361,
		PortalsNeutralized:          385,
		EnemyLinksDestroyed:         916,
		EnemyControlFieldsDestroyed: 393,

		MaxTimePortalHeld:     32,
		MaxTimeLinkMaintained: 25,
		MaxLinkLengthXDays:    180,
		MaxTimeFieldHeld:      15,
		LargestFieldMUsXDays:  3117,

		Hacks: 3247,
	},

	"steino_v0_unknown-1.png": profile.Profile{
		Nick:  "steino",
		Level: 8,
		AP:    1496166,

		UniquePortalsVisited: 643,
		XMCollected:          3443256,

		DistanceWalked: 149,

		ResonatorsDeployed:     3064,
		LinksCreated:           652,
		ControlFieldsCreated:   337,
		MindUnitsCaptured:      10669,
		LongestLinkEverCreated: 3,
		LargestControlField:    378,
		XMRecharged:            327766,
		PortalsCaptured:        259,
		UniquePortalsCaptured:  89,

		ResonatorsDestroyed:         1613,
		PortalsNeutralized:          252,
		EnemyLinksDestroyed:         258,
		EnemyControlFieldsDestroyed: 123,

		MaxTimePortalHeld:     11,
		MaxTimeLinkMaintained: 11,
		MaxLinkLengthXDays:    5,
		MaxTimeFieldHeld:      6,
		LargestFieldMUsXDays:  760,

		Hacks: 2521,
	},

	"tufte_v0_iphone.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    6939195,

		UniquePortalsVisited: 1460,
		XMCollected:          32500971,

		DistanceWalked: 581,

		ResonatorsDeployed:     10897,
		LinksCreated:           1656,
		ControlFieldsCreated:   1005,
		MindUnitsCaptured:      622776,
		LongestLinkEverCreated: 304,
		LargestControlField:    610877,
		XMRecharged:            15524531,
		PortalsCaptured:        918,
		UniquePortalsCaptured:  466,
		ModsDeployed:           1071,

		ResonatorsDestroyed:         7571,
		PortalsNeutralized:          903,
		EnemyLinksDestroyed:         1817,
		EnemyControlFieldsDestroyed: 1015,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    274,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  354224,

		UniqueMissionsCompleted: 2,

		Hacks:           8285,
		GlyphHackPoints: 705,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-1.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 11,
		AP:    7854616,

		UniquePortalsVisited: 1662,
		XMCollected:          37340314,

		DistanceWalked: 642,

		ResonatorsDeployed:     12452,
		LinksCreated:           1871,
		ControlFieldsCreated:   1106,
		MindUnitsCaptured:      623487,
		LongestLinkEverCreated: 304,
		LargestControlField:    610877,
		XMRecharged:            17258148,
		PortalsCaptured:        1093,
		UniquePortalsCaptured:  555,
		ModsDeployed:           1243,

		ResonatorsDestroyed:         9184,
		PortalsNeutralized:          1092,
		EnemyLinksDestroyed:         2184,
		EnemyControlFieldsDestroyed: 1214,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    274,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  354224,

		UniqueMissionsCompleted: 8,

		Hacks:           9260,
		GlyphHackPoints: 1112,

		AgentsSuccessfullyRecruited: 1,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-2.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 11,
		AP:    7898767,

		UniquePortalsVisited: 1730,
		XMCollected:          37508101,

		DistanceWalked: 646,

		ResonatorsDeployed:     12554,
		LinksCreated:           1891,
		ControlFieldsCreated:   1110,
		MindUnitsCaptured:      623524,
		LongestLinkEverCreated: 304,
		LargestControlField:    610877,
		XMRecharged:            17329909,
		PortalsCaptured:        1107,
		UniquePortalsCaptured:  569,
		ModsDeployed:           1250,

		ResonatorsDestroyed:         9236,
		PortalsNeutralized:          1101,
		EnemyLinksDestroyed:         2187,
		EnemyControlFieldsDestroyed: 1216,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    274,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  354224,

		UniqueMissionsCompleted: 8,

		Hacks:           9382,
		GlyphHackPoints: 1156,

		AgentsSuccessfullyRecruited: 1,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-3.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    4705187,

		UniquePortalsVisited: 1121,
		XMCollected:          18711907,

		DistanceWalked: 411,

		ResonatorsDeployed:     6802,
		LinksCreated:           1140,
		ControlFieldsCreated:   707,
		MindUnitsCaptured:      7944,
		LongestLinkEverCreated: 6,
		LargestControlField:    838,
		XMRecharged:            8997171,
		PortalsCaptured:        555,
		UniquePortalsCaptured:  295,
		ModsDeployed:           490,

		ResonatorsDestroyed:         4147,
		PortalsNeutralized:          498,
		EnemyLinksDestroyed:         965,
		EnemyControlFieldsDestroyed: 519,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    7,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  518,

		UniqueMissionsCompleted: 1,

		Hacks: 5501,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-4.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    4347932,

		UniquePortalsVisited: 1067,
		XMCollected:          16119070,

		DistanceWalked: 379,

		ResonatorsDeployed:     6105,
		LinksCreated:           1059,
		ControlFieldsCreated:   647,
		MindUnitsCaptured:      7440,
		LongestLinkEverCreated: 6,
		LargestControlField:    838,
		XMRecharged:            7252265,
		PortalsCaptured:        496,
		UniquePortalsCaptured:  264,
		ModsDeployed:           384,

		ResonatorsDestroyed:         3723,
		PortalsNeutralized:          443,
		EnemyLinksDestroyed:         861,
		EnemyControlFieldsDestroyed: 456,

		MaxTimePortalHeld:     27,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    7,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  518,

		Hacks: 4915,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-5.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    4268726,

		UniquePortalsVisited: 1053,
		XMCollected:          15378942,

		DistanceWalked: 368,

		ResonatorsDeployed:     5895,
		LinksCreated:           1037,
		ControlFieldsCreated:   632,
		MindUnitsCaptured:      7018,
		LongestLinkEverCreated: 6,
		LargestControlField:    838,
		XMRecharged:            6829037,
		PortalsCaptured:        491,
		UniquePortalsCaptured:  262,
		ModsDeployed:           337,

		ResonatorsDestroyed:         3616,
		PortalsNeutralized:          429,
		EnemyLinksDestroyed:         841,
		EnemyControlFieldsDestroyed: 449,

		MaxTimePortalHeld:     26,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    7,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  518,

		Hacks: 4808,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-6.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 8,
		AP:    3325704,

		UniquePortalsVisited: 854,
		XMCollected:          9734486,

		DistanceWalked: 279,

		ResonatorsDeployed:     4188,
		LinksCreated:           798,
		ControlFieldsCreated:   468,
		MindUnitsCaptured:      6091,
		LongestLinkEverCreated: 6,
		LargestControlField:    838,
		XMRecharged:            3851703,
		PortalsCaptured:        379,
		UniquePortalsCaptured:  208,

		ResonatorsDestroyed:         2645,
		PortalsNeutralized:          316,
		EnemyLinksDestroyed:         585,
		EnemyControlFieldsDestroyed: 288,

		MaxTimePortalHeld:     13,
		MaxTimeLinkMaintained: 9,
		MaxLinkLengthXDays:    6,
		MaxTimeFieldHeld:      7,
		LargestFieldMUsXDays:  417,

		Hacks: 3400,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-7.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 9,
		AP:    3579919,

		UniquePortalsVisited: 895,
		XMCollected:          10510155,

		DistanceWalked: 294,

		ResonatorsDeployed:     4543,
		LinksCreated:           877,
		ControlFieldsCreated:   510,
		MindUnitsCaptured:      6236,
		LongestLinkEverCreated: 6,
		LargestControlField:    838,
		XMRecharged:            4094199,
		PortalsCaptured:        409,
		UniquePortalsCaptured:  232,

		ResonatorsDestroyed:         2898,
		PortalsNeutralized:          348,
		EnemyLinksDestroyed:         657,
		EnemyControlFieldsDestroyed: 336,

		MaxTimePortalHeld:     13,
		MaxTimeLinkMaintained: 9,
		MaxLinkLengthXDays:    6,
		MaxTimeFieldHeld:      9,
		LargestFieldMUsXDays:  417,

		Hacks: 3591,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-8.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    5247462,

		UniquePortalsVisited: 1254,
		XMCollected:          22217022,

		DistanceWalked: 458,

		ResonatorsDeployed:     7722,
		LinksCreated:           1237,
		ControlFieldsCreated:   767,
		MindUnitsCaptured:      8365,
		LongestLinkEverCreated: 11,
		LargestControlField:    838,
		XMRecharged:            10598977,
		PortalsCaptured:        642,
		UniquePortalsCaptured:  335,
		ModsDeployed:           681,

		ResonatorsDestroyed:         4909,
		PortalsNeutralized:          591,
		EnemyLinksDestroyed:         1185,
		EnemyControlFieldsDestroyed: 661,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    7,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  518,

		UniqueMissionsCompleted: 1,

		Hacks: 6473,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-9.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    7544424,

		UniquePortalsVisited: 1506,
		XMCollected:          35235674,

		DistanceWalked: 616,

		ResonatorsDeployed:     11813,
		LinksCreated:           1805,
		ControlFieldsCreated:   1081,
		MindUnitsCaptured:      623364,
		LongestLinkEverCreated: 304,
		LargestControlField:    610877,
		XMRecharged:            16281832,
		PortalsCaptured:        1020,
		UniquePortalsCaptured:  503,
		ModsDeployed:           1173,

		ResonatorsDestroyed:         8588,
		PortalsNeutralized:          1015,
		EnemyLinksDestroyed:         2087,
		EnemyControlFieldsDestroyed: 1168,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    274,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  354224,

		UniqueMissionsCompleted: 4,

		Hacks:           8788,
		GlyphHackPoints: 1009,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-10.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 11,
		AP:    8370496,

		UniquePortalsVisited: 1972,
		XMCollected:          39880838,

		DistanceWalked: 674,

		ResonatorsDeployed:     13326,
		LinksCreated:           1980,
		ControlFieldsCreated:   1158,
		MindUnitsCaptured:      623572,
		LongestLinkEverCreated: 304,
		LargestControlField:    610877,
		XMRecharged:            18896676,
		PortalsCaptured:        1287,
		UniquePortalsCaptured:  694,
		ModsDeployed:           1290,

		ResonatorsDestroyed:         10156,
		PortalsNeutralized:          1252,
		EnemyLinksDestroyed:         2313,
		EnemyControlFieldsDestroyed: 1266,

		MaxTimePortalHeld:     39,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    274,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  354224,

		UniqueMissionsCompleted: 11,

		Hacks:                  10030,
		GlyphHackPoints:        1304,
		ConsecutiveDaysHacking: 1,

		AgentsSuccessfullyRecruited: 1,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-11.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    5926107,

		UniquePortalsVisited: 1311,
		XMCollected:          26234325,

		DistanceWalked: 506,

		ResonatorsDeployed:     9027,
		LinksCreated:           1442,
		ControlFieldsCreated:   875,
		MindUnitsCaptured:      9194,
		LongestLinkEverCreated: 11,
		LargestControlField:    838,
		XMRecharged:            12595563,
		PortalsCaptured:        747,
		UniquePortalsCaptured:  392,
		ModsDeployed:           874,

		ResonatorsDestroyed:         5905,
		PortalsNeutralized:          706,
		EnemyLinksDestroyed:         1415,
		EnemyControlFieldsDestroyed: 780,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    7,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  518,

		UniqueMissionsCompleted: 1,

		Hacks:           7147,
		GlyphHackPoints: 106,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-12.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 11,
		AP:    8017423,

		UniquePortalsVisited: 1895,
		XMCollected:          38457336,

		DistanceWalked: 655,

		ResonatorsDeployed:     12686,
		LinksCreated:           1891,
		ControlFieldsCreated:   1110,
		MindUnitsCaptured:      623524,
		LongestLinkEverCreated: 304,
		LargestControlField:    610877,
		XMRecharged:            17960793,
		PortalsCaptured:        1173,
		UniquePortalsCaptured:  635,
		ModsDeployed:           1257,

		ResonatorsDestroyed:         9566,
		PortalsNeutralized:          1161,
		EnemyLinksDestroyed:         2230,
		EnemyControlFieldsDestroyed: 1232,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    274,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  354224,

		UniqueMissionsCompleted: 9,

		Hacks:           9673,
		GlyphHackPoints: 1220,

		AgentsSuccessfullyRecruited: 1,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-13.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    5185518,

		UniquePortalsVisited: 1196,
		XMCollected:          21881822,

		DistanceWalked: 446,

		ResonatorsDeployed:     7574,
		LinksCreated:           1231,
		ControlFieldsCreated:   760,
		MindUnitsCaptured:      8313,
		LongestLinkEverCreated: 11,
		LargestControlField:    838,
		XMRecharged:            10442711,
		PortalsCaptured:        626,
		UniquePortalsCaptured:  325,
		ModsDeployed:           657,

		ResonatorsDestroyed:         4817,
		PortalsNeutralized:          574,
		EnemyLinksDestroyed:         1172,
		EnemyControlFieldsDestroyed: 655,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    7,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  518,

		UniqueMissionsCompleted: 1,

		Hacks: 6346,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-14.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 11,
		AP:    8271973,

		UniquePortalsVisited: 1972,
		XMCollected:          39653322,

		DistanceWalked: 670,

		ResonatorsDeployed:     13108,
		LinksCreated:           1955,
		ControlFieldsCreated:   1148,
		MindUnitsCaptured:      623562,
		LongestLinkEverCreated: 304,
		LargestControlField:    610877,
		XMRecharged:            18820163,
		PortalsCaptured:        1254,
		UniquePortalsCaptured:  694,
		ModsDeployed:           1284,

		ResonatorsDestroyed:         9921,
		PortalsNeutralized:          1216,
		EnemyLinksDestroyed:         2284,
		EnemyControlFieldsDestroyed: 1255,

		MaxTimePortalHeld:     37,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    274,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  354224,

		UniqueMissionsCompleted: 10,

		Hacks:           9956,
		GlyphHackPoints: 1252,

		AgentsSuccessfullyRecruited: 1,

		InnovatorLevel: 0,
	},

	"tufte_v0_iphone-15.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 11,
		AP:    8161319,

		UniquePortalsVisited: 1970,
		XMCollected:          39293283,

		DistanceWalked: 665,

		ResonatorsDeployed:     12936,
		LinksCreated:           1915,
		ControlFieldsCreated:   1125,
		MindUnitsCaptured:      623539,
		LongestLinkEverCreated: 304,
		LargestControlField:    610877,
		XMRecharged:            18559324,
		PortalsCaptured:        1234,
		UniquePortalsCaptured:  694,
		ModsDeployed:           1267,

		ResonatorsDestroyed:         9795,
		PortalsNeutralized:          1200,
		EnemyLinksDestroyed:         2262,
		EnemyControlFieldsDestroyed: 1246,

		MaxTimePortalHeld:     36,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    274,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  354224,

		UniqueMissionsCompleted: 10,

		Hacks:           9900,
		GlyphHackPoints: 1236,

		AgentsSuccessfullyRecruited: 1,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-16.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 10,
		AP:    5703653,

		UniquePortalsVisited: 1292,
		XMCollected:          24439953,

		DistanceWalked: 487,

		ResonatorsDeployed:     8594,
		LinksCreated:           1373,
		ControlFieldsCreated:   836,
		MindUnitsCaptured:      8910,
		LongestLinkEverCreated: 11,
		LargestControlField:    838,
		XMRecharged:            11419550,
		PortalsCaptured:        711,
		UniquePortalsCaptured:  379,
		ModsDeployed:           823,

		ResonatorsDestroyed:         5571,
		PortalsNeutralized:          665,
		EnemyLinksDestroyed:         1344,
		EnemyControlFieldsDestroyed: 746,

		MaxTimePortalHeld:     34,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    7,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  518,

		UniqueMissionsCompleted: 1,

		Hacks: 6904,

		InnovatorLevel: 3,
	},

	"viddy_v1711_oneplusone.png": profile.Profile{
		Nick:  "viddy",
		Level: 15,
		AP:    26164065,

		UniquePortalsVisited: 3169,
		PortalsDiscovered:    16,
		XMCollected:          122692805,

		DistanceWalked: 1606,

		ResonatorsDeployed:     56207,
		LinksCreated:           3786,
		ControlFieldsCreated:   1667,
		MindUnitsCaptured:      15602,
		LongestLinkEverCreated: 166,
		LargestControlField:    1250,
		XMRecharged:            34247088,
		PortalsCaptured:        8446,
		UniquePortalsCaptured:  1550,
		ModsDeployed:           1046,

		ResonatorsDestroyed:         53601,
		PortalsNeutralized:          7753,
		EnemyLinksDestroyed:         9663,
		EnemyControlFieldsDestroyed: 4743,

		MaxTimePortalHeld:     220,
		MaxTimeLinkMaintained: 67,
		MaxLinkLengthXDays:    98,
		MaxTimeFieldHeld:      64,
		LargestFieldMUsXDays:  5153,

		UniqueMissionsCompleted: 3,

		Hacks: 47195,
		ConsecutiveDaysHacking: 0,

		InnovatorLevel: 13,
	},

	"tufte_v0_iphone-17.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 9,
		AP:    3664876,

		UniquePortalsVisited: 902,
		XMCollected:          11199749,

		DistanceWalked: 301,

		ResonatorsDeployed:     4674,
		LinksCreated:           886,
		ControlFieldsCreated:   514,
		MindUnitsCaptured:      6275,
		LongestLinkEverCreated: 6,
		LargestControlField:    838,
		XMRecharged:            4494688,
		PortalsCaptured:        423,
		UniquePortalsCaptured:  240,
		ModsDeployed:           112,

		ResonatorsDestroyed:         3025,
		PortalsNeutralized:          363,
		EnemyLinksDestroyed:         695,
		EnemyControlFieldsDestroyed: 364,

		MaxTimePortalHeld:     14,
		MaxTimeLinkMaintained: 14,
		MaxLinkLengthXDays:    7,
		MaxTimeFieldHeld:      14,
		LargestFieldMUsXDays:  518,

		Hacks: 3732,

		InnovatorLevel: 3,
	},

	"mvnch_v0_iphone-1.jpeg": profile.Profile{
		Nick:  "mvnch",
		Level: 7,
		AP:    606000,

		UniquePortalsVisited: 134,
		XMCollected:          1076570,

		DistanceWalked: 61,

		ResonatorsDeployed:     707,
		LinksCreated:           272,
		ControlFieldsCreated:   168,
		MindUnitsCaptured:      613,
		LongestLinkEverCreated: 4,
		LargestControlField:    53,
		XMRecharged:            351923,
		PortalsCaptured:        67,
		UniquePortalsCaptured:  39,

		ResonatorsDestroyed:         272,
		PortalsNeutralized:          35,
		EnemyLinksDestroyed:         73,
		EnemyControlFieldsDestroyed: 32,

		MaxTimePortalHeld:     20,
		MaxTimeLinkMaintained: 13,
		MaxLinkLengthXDays:    35,
		MaxTimeFieldHeld:      12,
		LargestFieldMUsXDays:  531,

		Hacks: 647,
	},

	"mvnch_v0_iphone-2.jpeg": profile.Profile{
		Nick:  "mvnch",
		Level: 5,
		AP:    174137,

		UniquePortalsVisited: 108,
		XMCollected:          189816,

		DistanceWalked: 28,

		ResonatorsDeployed:     249,
		LinksCreated:           109,
		ControlFieldsCreated:   65,
		MindUnitsCaptured:      232,
		LongestLinkEverCreated: 4,
		LargestControlField:    53,
		XMRecharged:            28996,
		PortalsCaptured:        19,
		UniquePortalsCaptured:  16,

		ResonatorsDestroyed:         30,
		PortalsNeutralized:          4,
		EnemyLinksDestroyed:         5,
		EnemyControlFieldsDestroyed: 1,

		MaxTimePortalHeld:     5,
		MaxTimeLinkMaintained: 5,
		MaxLinkLengthXDays:    4,
		MaxTimeFieldHeld:      5,
		LargestFieldMUsXDays:  36,

		Hacks: 266,
	},

	"haste_v0_nexus6.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    26336070,

		UniquePortalsVisited: 3564,
		PortalsDiscovered:    37,
		XMCollected:          109992991,

		DistanceWalked: 2075,

		ResonatorsDeployed:     55692,
		LinksCreated:           5353,
		ControlFieldsCreated:   2580,
		MindUnitsCaptured:      32872,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26544040,
		PortalsCaptured:        6397,
		UniquePortalsCaptured:  1432,
		ModsDeployed:           687,

		ResonatorsDestroyed:         46022,
		PortalsNeutralized:          6816,
		EnemyLinksDestroyed:         8850,
		EnemyControlFieldsDestroyed: 4244,

		MaxTimePortalHeld:     132,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 53791,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus6-1.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    30120647,

		UniquePortalsVisited: 3815,
		PortalsDiscovered:    40,
		XMCollected:          125280853,

		DistanceWalked: 2283,

		ResonatorsDeployed:     64049,
		LinksCreated:           5763,
		ControlFieldsCreated:   2802,
		MindUnitsCaptured:      33477,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            27725336,
		PortalsCaptured:        7203,
		UniquePortalsCaptured:  1530,
		ModsDeployed:           1622,

		ResonatorsDestroyed:         52267,
		PortalsNeutralized:          7683,
		EnemyLinksDestroyed:         10615,
		EnemyControlFieldsDestroyed: 5242,

		MaxTimePortalHeld:     170,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 18,

		Hacks:           58408,
		GlyphHackPoints: 9630,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus6-2.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    29053389,

		UniquePortalsVisited: 3711,
		PortalsDiscovered:    40,
		XMCollected:          120915318,

		DistanceWalked: 2221,

		ResonatorsDeployed:     61867,
		LinksCreated:           5674,
		ControlFieldsCreated:   2763,
		MindUnitsCaptured:      33367,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            27631913,
		PortalsCaptured:        6946,
		UniquePortalsCaptured:  1476,
		ModsDeployed:           1333,

		ResonatorsDestroyed:         50484,
		PortalsNeutralized:          7424,
		EnemyLinksDestroyed:         10096,
		EnemyControlFieldsDestroyed: 4949,

		MaxTimePortalHeld:     157,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 11,

		Hacks:           56899,
		GlyphHackPoints: 6162,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus6-3.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    28630454,

		UniquePortalsVisited: 3689,
		PortalsDiscovered:    40,
		XMCollected:          119639743,

		DistanceWalked: 2196,

		ResonatorsDeployed:     61125,
		LinksCreated:           5641,
		ControlFieldsCreated:   2737,
		MindUnitsCaptured:      33229,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            27611293,
		PortalsCaptured:        6865,
		UniquePortalsCaptured:  1475,
		ModsDeployed:           1276,

		ResonatorsDestroyed:         49655,
		PortalsNeutralized:          7319,
		EnemyLinksDestroyed:         9858,
		EnemyControlFieldsDestroyed: 4801,

		MaxTimePortalHeld:     155,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 11,

		Hacks:           56525,
		GlyphHackPoints: 4854,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus5.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    18698368,

		UniquePortalsVisited: 3088,
		PortalsDiscovered:    32,
		XMCollected:          85919804,

		DistanceWalked: 1702,

		ResonatorsDeployed:     42035,
		LinksCreated:           4630,
		ControlFieldsCreated:   2249,
		MindUnitsCaptured:      32197,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            22796779,
		PortalsCaptured:        5124,
		UniquePortalsCaptured:  1333,

		ResonatorsDestroyed:         36527,
		PortalsNeutralized:          5384,
		EnemyLinksDestroyed:         6362,
		EnemyControlFieldsDestroyed: 2823,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 43819,
	},

	"haste_v0_nexus5-1.png": profile.Profile{
		Nick:  "haste",
		Level: 13,
		AP:    15398687,

		UniquePortalsVisited: 2993,
		PortalsDiscovered:    20,
		XMCollected:          73273719,

		DistanceWalked: 1457,

		ResonatorsDeployed:     36334,
		LinksCreated:           3683,
		ControlFieldsCreated:   1700,
		MindUnitsCaptured:      30674,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            19206430,
		PortalsCaptured:        4389,
		UniquePortalsCaptured:  1250,

		ResonatorsDestroyed:         31621,
		PortalsNeutralized:          4599,
		EnemyLinksDestroyed:         5269,
		EnemyControlFieldsDestroyed: 2260,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 38504,
	},

	"haste_v0_nexus5-2.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    22520735,

		UniquePortalsVisited: 3427,
		PortalsDiscovered:    34,
		XMCollected:          97966376,

		DistanceWalked: 1872,

		ResonatorsDeployed:     47855,
		LinksCreated:           5088,
		ControlFieldsCreated:   2459,
		MindUnitsCaptured:      32619,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            25459736,
		PortalsCaptured:        5722,
		UniquePortalsCaptured:  1377,

		ResonatorsDestroyed:         41723,
		PortalsNeutralized:          6181,
		EnemyLinksDestroyed:         7654,
		EnemyControlFieldsDestroyed: 3522,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 48486,
	},

	"haste_v0_nexus5-3.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    20269836,

		UniquePortalsVisited: 3248,
		PortalsDiscovered:    32,
		XMCollected:          92640775,

		DistanceWalked: 1810,

		ResonatorsDeployed:     45267,
		LinksCreated:           5041,
		ControlFieldsCreated:   2445,
		MindUnitsCaptured:      32598,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            24807345,
		PortalsCaptured:        5446,
		UniquePortalsCaptured:  1352,

		ResonatorsDestroyed:         38911,
		PortalsNeutralized:          5785,
		EnemyLinksDestroyed:         6893,
		EnemyControlFieldsDestroyed: 3090,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 46596,
	},

	"haste_v0_nexus5-4.png": profile.Profile{
		Nick:  "haste",
		Level: 13,
		AP:    15213400,

		UniquePortalsVisited: 2915,
		PortalsDiscovered:    19,
		XMCollected:          72073572,

		DistanceWalked: 1440,

		ResonatorsDeployed:     35836,
		LinksCreated:           3656,
		ControlFieldsCreated:   1686,
		MindUnitsCaptured:      30634,
		LongestLinkEverCreated: 78,
		LargestControlField:    19445,
		XMRecharged:            18738511,
		PortalsCaptured:        4346,
		UniquePortalsCaptured:  1248,

		ResonatorsDestroyed:         31162,
		PortalsNeutralized:          4538,
		EnemyLinksDestroyed:         5196,
		EnemyControlFieldsDestroyed: 2232,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 38020,
	},

	"haste_v0_nexus5-5.png": profile.Profile{
		Nick:  "haste",
		Level: 13,
		AP:    14906187,

		UniquePortalsVisited: 2913,
		PortalsDiscovered:    19,
		XMCollected:          70852194,

		DistanceWalked: 1412,

		ResonatorsDeployed:     35195,
		LinksCreated:           3585,
		ControlFieldsCreated:   1641,
		MindUnitsCaptured:      30538,
		LongestLinkEverCreated: 78,
		LargestControlField:    19445,
		XMRecharged:            18308593,
		PortalsCaptured:        4273,
		UniquePortalsCaptured:  1246,

		ResonatorsDestroyed:         30607,
		PortalsNeutralized:          4444,
		EnemyLinksDestroyed:         5102,
		EnemyControlFieldsDestroyed: 2190,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 37458,
	},

	"haste_v0_nexus5-6.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    24219616,

		UniquePortalsVisited: 3453,
		PortalsDiscovered:    36,
		XMCollected:          101037503,

		DistanceWalked: 1921,

		ResonatorsDeployed:     50757,
		LinksCreated:           5156,
		ControlFieldsCreated:   2496,
		MindUnitsCaptured:      32712,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            25602971,
		PortalsCaptured:        5991,
		UniquePortalsCaptured:  1396,

		ResonatorsDestroyed:         43434,
		PortalsNeutralized:          6428,
		EnemyLinksDestroyed:         8030,
		EnemyControlFieldsDestroyed: 3708,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 49506,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus5-7.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    24323562,

		UniquePortalsVisited: 3482,
		PortalsDiscovered:    36,
		XMCollected:          101487477,

		DistanceWalked: 1934,

		ResonatorsDeployed:     50974,
		LinksCreated:           5171,
		ControlFieldsCreated:   2502,
		MindUnitsCaptured:      32748,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            25647448,
		PortalsCaptured:        6014,
		UniquePortalsCaptured:  1401,

		ResonatorsDestroyed:         43577,
		PortalsNeutralized:          6453,
		EnemyLinksDestroyed:         8078,
		EnemyControlFieldsDestroyed: 3734,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 49749,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus5-8.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    24861416,

		UniquePortalsVisited: 3491,
		PortalsDiscovered:    36,
		XMCollected:          103432896,

		DistanceWalked: 1978,

		ResonatorsDeployed:     52102,
		LinksCreated:           5214,
		ControlFieldsCreated:   2526,
		MindUnitsCaptured:      32776,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26309525,
		PortalsCaptured:        6115,
		UniquePortalsCaptured:  1406,

		ResonatorsDestroyed:         44216,
		PortalsNeutralized:          6543,
		EnemyLinksDestroyed:         8263,
		EnemyControlFieldsDestroyed: 3855,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 50719,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus5-9.png": profile.Profile{
		Nick:  "haste",
		Level: 13,
		AP:    16358734,

		UniquePortalsVisited: 3036,
		PortalsDiscovered:    24,
		XMCollected:          77118158,

		DistanceWalked: 1523,

		ResonatorsDeployed:     37644,
		LinksCreated:           3999,
		ControlFieldsCreated:   1885,
		MindUnitsCaptured:      31139,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            20335001,
		PortalsCaptured:        4559,
		UniquePortalsCaptured:  1284,

		ResonatorsDestroyed:         32998,
		PortalsNeutralized:          4797,
		EnemyLinksDestroyed:         5630,
		EnemyControlFieldsDestroyed: 2461,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 39823,
	},

	"haste_v0_nexus5-10.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    17008780,

		UniquePortalsVisited: 3048,
		PortalsDiscovered:    24,
		XMCollected:          78338002,

		DistanceWalked: 1552,

		ResonatorsDeployed:     38646,
		LinksCreated:           4300,
		ControlFieldsCreated:   2063,
		MindUnitsCaptured:      31462,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            20421902,
		PortalsCaptured:        4641,
		UniquePortalsCaptured:  1288,

		ResonatorsDestroyed:         33501,
		PortalsNeutralized:          4888,
		EnemyLinksDestroyed:         5764,
		EnemyControlFieldsDestroyed: 2544,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 40589,
	},

	"haste_v0_nexus5-11.png": profile.Profile{
		Nick:  "haste",
		Level: 13,
		AP:    15967035,

		UniquePortalsVisited: 3008,
		PortalsDiscovered:    24,
		XMCollected:          75156147,

		DistanceWalked: 1493,

		ResonatorsDeployed:     37229,
		LinksCreated:           3953,
		ControlFieldsCreated:   1845,
		MindUnitsCaptured:      31025,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            19799770,
		PortalsCaptured:        4478,
		UniquePortalsCaptured:  1255,

		ResonatorsDestroyed:         32236,
		PortalsNeutralized:          4705,
		EnemyLinksDestroyed:         5382,
		EnemyControlFieldsDestroyed: 2310,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 39236,
	},

	"haste_v0_nexus5-12.png": profile.Profile{
		Nick:  "haste",
		Level: 13,
		AP:    15211750,

		UniquePortalsVisited: 2915,
		PortalsDiscovered:    19,
		XMCollected:          72030337,

		DistanceWalked: 1439,

		ResonatorsDeployed:     35836,
		LinksCreated:           3656,
		ControlFieldsCreated:   1686,
		MindUnitsCaptured:      30634,
		LongestLinkEverCreated: 78,
		LargestControlField:    19445,
		XMRecharged:            18660322,
		PortalsCaptured:        4346,
		UniquePortalsCaptured:  1248,

		ResonatorsDestroyed:         31162,
		PortalsNeutralized:          4538,
		EnemyLinksDestroyed:         5196,
		EnemyControlFieldsDestroyed: 2232,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 38020,
	},

	"haste_v0_nexus5-13.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    24062506,

		UniquePortalsVisited: 3448,
		PortalsDiscovered:    36,
		XMCollected:          100613405,

		DistanceWalked: 1916,

		ResonatorsDeployed:     50251,
		LinksCreated:           5152,
		ControlFieldsCreated:   2494,
		MindUnitsCaptured:      32708,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            25588437,
		PortalsCaptured:        5935,
		UniquePortalsCaptured:  1391,

		ResonatorsDestroyed:         43060,
		PortalsNeutralized:          6383,
		EnemyLinksDestroyed:         7975,
		EnemyControlFieldsDestroyed: 3687,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 49294,
	},

	"haste_v0_nexus5-14.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    18781682,

		UniquePortalsVisited: 3089,
		PortalsDiscovered:    32,
		XMCollected:          86301856,

		DistanceWalked: 1715,

		ResonatorsDeployed:     42192,
		LinksCreated:           4657,
		ControlFieldsCreated:   2266,
		MindUnitsCaptured:      32236,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            22861940,
		PortalsCaptured:        5134,
		UniquePortalsCaptured:  1334,

		ResonatorsDestroyed:         36599,
		PortalsNeutralized:          5394,
		EnemyLinksDestroyed:         6386,
		EnemyControlFieldsDestroyed: 2838,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 44036,
	},

	"haste_v0_nexus5-15.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    25575067,

		UniquePortalsVisited: 3515,
		PortalsDiscovered:    36,
		XMCollected:          106888705,

		DistanceWalked: 2021,

		ResonatorsDeployed:     53825,
		LinksCreated:           5264,
		ControlFieldsCreated:   2546,
		MindUnitsCaptured:      32799,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26509239,
		PortalsCaptured:        6249,
		UniquePortalsCaptured:  1408,
		ModsDeployed:           403,

		ResonatorsDestroyed:         45095,
		PortalsNeutralized:          6674,
		EnemyLinksDestroyed:         8508,
		EnemyControlFieldsDestroyed: 4006,

		MaxTimePortalHeld:     114,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 52369,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus6-4.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    26707468,

		UniquePortalsVisited: 3572,
		PortalsDiscovered:    39,
		XMCollected:          111924182,

		DistanceWalked: 2098,

		ResonatorsDeployed:     56636,
		LinksCreated:           5374,
		ControlFieldsCreated:   2593,
		MindUnitsCaptured:      32903,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            27056201,
		PortalsCaptured:        6498,
		UniquePortalsCaptured:  1435,
		ModsDeployed:           819,

		ResonatorsDestroyed:         46737,
		PortalsNeutralized:          6921,
		EnemyLinksDestroyed:         9005,
		EnemyControlFieldsDestroyed: 4324,

		MaxTimePortalHeld:     137,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks:           54326,
		GlyphHackPoints: 73,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus6-5.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    28454788,

		UniquePortalsVisited: 3687,
		PortalsDiscovered:    40,
		XMCollected:          118912429,

		DistanceWalked: 2186,

		ResonatorsDeployed:     60710,
		LinksCreated:           5639,
		ControlFieldsCreated:   2736,
		MindUnitsCaptured:      33206,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            27608651,
		PortalsCaptured:        6827,
		UniquePortalsCaptured:  1470,
		ModsDeployed:           1249,

		ResonatorsDestroyed:         49323,
		PortalsNeutralized:          7277,
		EnemyLinksDestroyed:         9748,
		EnemyControlFieldsDestroyed: 4738,

		MaxTimePortalHeld:     154,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 10,

		Hacks:           56411,
		GlyphHackPoints: 4297,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus5-16.png": profile.Profile{
		Nick:  "haste",
		Level: 13,
		AP:    16156376,

		UniquePortalsVisited: 3032,
		PortalsDiscovered:    24,
		XMCollected:          76076011,

		DistanceWalked: 1506,

		ResonatorsDeployed:     37394,
		LinksCreated:           3960,
		ControlFieldsCreated:   1852,
		MindUnitsCaptured:      31038,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            19877113,
		PortalsCaptured:        4523,
		UniquePortalsCaptured:  1274,

		ResonatorsDestroyed:         32695,
		PortalsNeutralized:          4758,
		EnemyLinksDestroyed:         5527,
		EnemyControlFieldsDestroyed: 2401,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 39515,
	},

	"haste_v0_nexus5-17.png": profile.Profile{
		Nick:  "haste",
		Level: 13,
		AP:    16846515,

		UniquePortalsVisited: 3048,
		PortalsDiscovered:    24,
		XMCollected:          78111812,

		DistanceWalked: 1545,

		ResonatorsDeployed:     38335,
		LinksCreated:           4203,
		ControlFieldsCreated:   2019,
		MindUnitsCaptured:      31409,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            20403879,
		PortalsCaptured:        4632,
		UniquePortalsCaptured:  1288,

		ResonatorsDestroyed:         33418,
		PortalsNeutralized:          4876,
		EnemyLinksDestroyed:         5747,
		EnemyControlFieldsDestroyed: 2536,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 40419,
	},

	"haste_v0_nexus5-18.png": profile.Profile{
		Nick:  "haste",
		Level: 13,
		AP:    15273144,

		UniquePortalsVisited: 2980,
		PortalsDiscovered:    20,
		XMCollected:          72828849,

		DistanceWalked: 1451,

		ResonatorsDeployed:     36082,
		LinksCreated:           3661,
		ControlFieldsCreated:   1689,
		MindUnitsCaptured:      30644,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            19057556,
		PortalsCaptured:        4349,
		UniquePortalsCaptured:  1250,

		ResonatorsDestroyed:         31282,
		PortalsNeutralized:          4559,
		EnemyLinksDestroyed:         5213,
		EnemyControlFieldsDestroyed: 2238,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 38343,
	},

	"viddy_v0_unknown.png": profile.Profile{
		Nick:  "viddy",
		Level: 14,
		AP:    17115185,

		UniquePortalsVisited: 3009,
		PortalsDiscovered:    15,
		XMCollected:          97394185,

		DistanceWalked: 1353,

		ResonatorsDeployed:     38722,
		LinksCreated:           3470,
		ControlFieldsCreated:   1467,
		MindUnitsCaptured:      15179,
		LongestLinkEverCreated: 166,
		LargestControlField:    1250,
		XMRecharged:            30993663,
		PortalsCaptured:        5782,
		UniquePortalsCaptured:  1468,
		ModsDeployed:           1468,

		ResonatorsDestroyed:         37317,
		PortalsNeutralized:          5211,
		EnemyLinksDestroyed:         6135,
		EnemyControlFieldsDestroyed: 2934,

		MaxTimePortalHeld:     185,
		MaxTimeLinkMaintained: 67,
		MaxLinkLengthXDays:    98,
		MaxTimeFieldHeld:      64,
		LargestFieldMUsXDays:  5153,

		Hacks: 37552,
	},

	"viddy_v0_unknown-1.png": profile.Profile{
		Nick:  "viddy",
		Level: 13,
		AP:    15181213,

		UniquePortalsVisited: 2789,
		PortalsDiscovered:    15,
		XMCollected:          89935217,

		DistanceWalked: 1272,

		ResonatorsDeployed:     34982,
		LinksCreated:           3181,
		ControlFieldsCreated:   1286,
		MindUnitsCaptured:      14754,
		LongestLinkEverCreated: 166,
		LargestControlField:    1250,
		XMRecharged:            29128903,
		PortalsCaptured:        5089,
		UniquePortalsCaptured:  1375,

		ResonatorsDestroyed:         33275,
		PortalsNeutralized:          4565,
		EnemyLinksDestroyed:         5380,
		EnemyControlFieldsDestroyed: 2570,

		MaxTimePortalHeld:     155,
		MaxTimeLinkMaintained: 67,
		MaxLinkLengthXDays:    98,
		MaxTimeFieldHeld:      64,
		LargestFieldMUsXDays:  5153,

		Hacks: 34808,
	},

	"haste_v0_nexus5-19.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    17709965,

		UniquePortalsVisited: 3058,
		PortalsDiscovered:    29,
		XMCollected:          81346452,

		DistanceWalked: 1617,

		ResonatorsDeployed:     39980,
		LinksCreated:           4466,
		ControlFieldsCreated:   2167,
		MindUnitsCaptured:      31943,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            21398558,
		PortalsCaptured:        4797,
		UniquePortalsCaptured:  1298,

		ResonatorsDestroyed:         34694,
		PortalsNeutralized:          5072,
		EnemyLinksDestroyed:         6010,
		EnemyControlFieldsDestroyed: 2664,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 41971,
	},

	"haste_v0_nexus5-20.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    17709965,

		UniquePortalsVisited: 3058,
		PortalsDiscovered:    29,
		XMCollected:          81346452,

		DistanceWalked: 1617,

		ResonatorsDeployed:     39980,
		LinksCreated:           4466,
		ControlFieldsCreated:   2167,
		MindUnitsCaptured:      31943,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            21398558,
		PortalsCaptured:        4797,
		UniquePortalsCaptured:  1298,

		ResonatorsDestroyed:         34694,
		PortalsNeutralized:          5072,
		EnemyLinksDestroyed:         6010,
		EnemyControlFieldsDestroyed: 2664,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 41971,
	},

	"haste_v0_nexus5-21.png": profile.Profile{
		Nick:  "haste",
		Level: 13,
		AP:    16487110,

		UniquePortalsVisited: 3036,
		PortalsDiscovered:    24,
		XMCollected:          77359545,

		DistanceWalked: 1528,

		ResonatorsDeployed:     37839,
		LinksCreated:           4040,
		ControlFieldsCreated:   1910,
		MindUnitsCaptured:      31178,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            20335001,
		PortalsCaptured:        4601,
		UniquePortalsCaptured:  1284,

		ResonatorsDestroyed:         33212,
		PortalsNeutralized:          4849,
		EnemyLinksDestroyed:         5664,
		EnemyControlFieldsDestroyed: 2476,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 39964,
	},

	"haste_v0_nexus5-22.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    18571386,

		UniquePortalsVisited: 3081,
		PortalsDiscovered:    31,
		XMCollected:          85417585,

		DistanceWalked: 1692,

		ResonatorsDeployed:     41826,
		LinksCreated:           4593,
		ControlFieldsCreated:   2232,
		MindUnitsCaptured:      32146,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            22593900,
		PortalsCaptured:        5089,
		UniquePortalsCaptured:  1333,

		ResonatorsDestroyed:         36340,
		PortalsNeutralized:          5351,
		EnemyLinksDestroyed:         6319,
		EnemyControlFieldsDestroyed: 2799,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 43614,
	},

	"haste_v0_nexus5-23.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    17401066,

		UniquePortalsVisited: 3052,
		PortalsDiscovered:    28,
		XMCollected:          79744982,

		DistanceWalked: 1577,

		ResonatorsDeployed:     39388,
		LinksCreated:           4367,
		ControlFieldsCreated:   2098,
		MindUnitsCaptured:      31535,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            20786548,
		PortalsCaptured:        4745,
		UniquePortalsCaptured:  1293,

		ResonatorsDestroyed:         34334,
		PortalsNeutralized:          5017,
		EnemyLinksDestroyed:         5944,
		EnemyControlFieldsDestroyed: 2626,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 41119,
	},

	"haste_v0_nexus5-24.png": profile.Profile{
		Nick:  "haste",
		Level: 13,
		AP:    16754299,

		UniquePortalsVisited: 3048,
		PortalsDiscovered:    24,
		XMCollected:          77985729,

		DistanceWalked: 1539,

		ResonatorsDeployed:     38133,
		LinksCreated:           4146,
		ControlFieldsCreated:   1986,
		MindUnitsCaptured:      31357,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            20403879,
		PortalsCaptured:        4630,
		UniquePortalsCaptured:  1288,

		ResonatorsDestroyed:         33418,
		PortalsNeutralized:          4876,
		EnemyLinksDestroyed:         5747,
		EnemyControlFieldsDestroyed: 2536,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 40259,
	},

	"haste_v0_nexus5-25.png": profile.Profile{
		Nick:  "haste",
		Level: 14,
		AP:    17581951,

		UniquePortalsVisited: 3053,
		PortalsDiscovered:    28,
		XMCollected:          80482237,

		DistanceWalked: 1595,

		ResonatorsDeployed:     39683,
		LinksCreated:           4429,
		ControlFieldsCreated:   2137,
		MindUnitsCaptured:      31715,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            21007953,
		PortalsCaptured:        4767,
		UniquePortalsCaptured:  1293,

		ResonatorsDestroyed:         34519,
		PortalsNeutralized:          5037,
		EnemyLinksDestroyed:         6001,
		EnemyControlFieldsDestroyed: 2663,

		MaxTimePortalHeld:     109,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 41470,
	},

	"gnarfoz_v0_unknown.png": profile.Profile{
		Nick:  "Gnarfoz",
		Level: 10,
		AP:    5114971,

		UniquePortalsVisited: 1253,
		PortalsDiscovered:    38,
		XMCollected:          31181555,

		DistanceWalked: 345,

		ResonatorsDeployed:     11342,
		LinksCreated:           2803,
		ControlFieldsCreated:   725,
		MindUnitsCaptured:      440607,
		LongestLinkEverCreated: 94,
		LargestControlField:    152278,
		XMRecharged:            15887881,
		PortalsCaptured:        901,
		UniquePortalsCaptured:  402,

		ResonatorsDestroyed:         7918,
		PortalsNeutralized:          812,
		EnemyLinksDestroyed:         1765,
		EnemyControlFieldsDestroyed: 741,

		MaxTimePortalHeld:     107,
		MaxTimeLinkMaintained: 94,
		MaxLinkLengthXDays:    111,
		MaxTimeFieldHeld:      54,
		LargestFieldMUsXDays:  12718,

		Hacks: 15128,
	},

	// Monthly screenshot.
	"gnarfoz_v0_unknown-1.png": profile.Profile{
		Nick:  "Gnarfoz",
		Level: 10,
		AP:    6968731,

		UniquePortalsVisited: 773,
		PortalsDiscovered:    0,
		XMCollected:          4720898,

		DistanceWalked: 70,

		ResonatorsDeployed:     1498,
		LinksCreated:           316,
		ControlFieldsCreated:   153,
		MindUnitsCaptured:      1975,
		LongestLinkEverCreated: 3,
		LargestControlField:    327,
		XMRecharged:            1410524,
		PortalsCaptured:        169,
		UniquePortalsCaptured:  132,
		ModsDeployed:           445,

		ResonatorsDestroyed:         1767,
		PortalsNeutralized:          295,
		EnemyLinksDestroyed:         288,
		EnemyControlFieldsDestroyed: 162,

		MaxTimePortalHeld:     187,
		MaxTimeLinkMaintained: 24,
		MaxLinkLengthXDays:    20,
		MaxTimeFieldHeld:      24,
		LargestFieldMUsXDays:  5428,

		UniqueMissionsCompleted: 24,

		Hacks: 1894,

		InnovatorLevel: 9,
	},

	"arilds_v0_unknown.png": profile.Profile{
		Nick:  "ArildS",
		Level: 10,
		AP:    15400656,

		UniquePortalsVisited: 1590,
		PortalsDiscovered:    54,
		XMCollected:          83677640,

		DistanceWalked: 814,

		ResonatorsDeployed:     29195,
		LinksCreated:           3613,
		ControlFieldsCreated:   1608,
		MindUnitsCaptured:      254349,
		LongestLinkEverCreated: 207,
		LargestControlField:    114599,
		XMRecharged:            27971604,
		PortalsCaptured:        3182,
		UniquePortalsCaptured:  884,

		ResonatorsDestroyed:         28549,
		PortalsNeutralized:          3764,
		EnemyLinksDestroyed:         6398,
		EnemyControlFieldsDestroyed: 3351,

		MaxTimePortalHeld:     306,
		MaxTimeLinkMaintained: 89,
		MaxLinkLengthXDays:    628,
		MaxTimeFieldHeld:      89,
		LargestFieldMUsXDays:  116533,

		Hacks: 40960,
	},

	"forferdet_v0_unknown-1.png": profile.Profile{
		Nick:  "forferdet",
		Level: 11,
		AP:    12750621,

		UniquePortalsVisited: 1079,
		XMCollected:          53925206,

		DistanceWalked: 929,

		ResonatorsDeployed:     22753,
		LinksCreated:           5074,
		ControlFieldsCreated:   2637,
		MindUnitsCaptured:      18954,
		LongestLinkEverCreated: 45,
		LargestControlField:    692,
		XMRecharged:            24629947,
		PortalsCaptured:        2294,
		UniquePortalsCaptured:  706,

		ResonatorsDestroyed:         15851,
		PortalsNeutralized:          2303,
		EnemyLinksDestroyed:         3046,
		EnemyControlFieldsDestroyed: 1451,

		MaxTimePortalHeld:     37,
		MaxTimeLinkMaintained: 55,
		MaxLinkLengthXDays:    93,
		MaxTimeFieldHeld:      52,
		LargestFieldMUsXDays:  540,

		Hacks: 21386,
	},

	"steino_v0_unknown-2.png": profile.Profile{
		Nick:  "steino",
		Level: 8,
		AP:    1637518,

		UniquePortalsVisited: 777,
		PortalsDiscovered:    2,
		XMCollected:          4073906,

		DistanceWalked: 174,

		ResonatorsDeployed:     3453,
		LinksCreated:           696,
		ControlFieldsCreated:   363,
		MindUnitsCaptured:      10740,
		LongestLinkEverCreated: 3,
		LargestControlField:    378,
		XMRecharged:            485829,
		PortalsCaptured:        274,
		UniquePortalsCaptured:  99,
		ModsDeployed:           4,

		ResonatorsDestroyed:         1762,
		PortalsNeutralized:          277,
		EnemyLinksDestroyed:         283,
		EnemyControlFieldsDestroyed: 136,

		MaxTimePortalHeld:     16,
		MaxTimeLinkMaintained: 14,
		MaxLinkLengthXDays:    5,
		MaxTimeFieldHeld:      6,
		LargestFieldMUsXDays:  760,

		Hacks: 2849,

		InnovatorLevel: 3,
	},

	"bentis_v0_unknown.png": profile.Profile{
		Nick:  "Bentis",
		Level: 16,
		AP:    40795452,

		UniquePortalsVisited: 7953,
		PortalsDiscovered:    22,
		XMCollected:          113750958,

		DistanceWalked: 1965,

		ResonatorsDeployed:     58346,
		LinksCreated:           6374,
		ControlFieldsCreated:   3084,
		MindUnitsCaptured:      73932,
		LongestLinkEverCreated: 99,
		LargestControlField:    22944,
		XMRecharged:            42141305,
		PortalsCaptured:        14709,
		UniquePortalsCaptured:  4609,
		ModsDeployed:           547,

		ResonatorsDestroyed:         75702,
		PortalsNeutralized:          11971,
		EnemyLinksDestroyed:         14331,
		EnemyControlFieldsDestroyed: 6348,

		MaxTimePortalHeld:     178,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    1619,
		MaxTimeFieldHeld:      51,
		LargestFieldMUsXDays:  26317,

		UniqueMissionsCompleted: 7,

		Hacks: 39176,

		InnovatorLevel: 13,
	},

	"forferdet_v0_unknown-2.png": profile.Profile{
		Nick:  "forferdet",
		Level: 15,
		AP:    30157205,

		UniquePortalsVisited: 1885,
		PortalsDiscovered:    14,
		XMCollected:          121396161,

		DistanceWalked: 1601,

		ResonatorsDeployed:     49864,
		LinksCreated:           8215,
		ControlFieldsCreated:   3901,
		MindUnitsCaptured:      26106,
		LongestLinkEverCreated: 156,
		LargestControlField:    977,
		XMRecharged:            40035020,
		PortalsCaptured:        5393,
		UniquePortalsCaptured:  1141,
		ModsDeployed:           3030,

		ResonatorsDestroyed:         45035,
		PortalsNeutralized:          6219,
		EnemyLinksDestroyed:         9436,
		EnemyControlFieldsDestroyed: 4731,

		MaxTimePortalHeld:     68,
		MaxTimeLinkMaintained: 55,
		MaxLinkLengthXDays:    2543,
		MaxTimeFieldHeld:      52,
		LargestFieldMUsXDays:  2632,

		UniqueMissionsCompleted: 9,

		Hacks:           44838,
		GlyphHackPoints: 35239,
	},

	"forferdet_v0_unknown-3.png": profile.Profile{
		Nick:  "forferdet",
		Level: 15,
		AP:    29044336,

		UniquePortalsVisited: 1874,
		PortalsDiscovered:    13,
		XMCollected:          116276662,

		DistanceWalked: 1560,

		ResonatorsDeployed:     48179,
		LinksCreated:           8033,
		ControlFieldsCreated:   3843,
		MindUnitsCaptured:      25903,
		LongestLinkEverCreated: 156,
		LargestControlField:    977,
		XMRecharged:            38881810,
		PortalsCaptured:        5184,
		UniquePortalsCaptured:  1136,
		ModsDeployed:           2658,

		ResonatorsDestroyed:         42729,
		PortalsNeutralized:          5912,
		EnemyLinksDestroyed:         8910,
		EnemyControlFieldsDestroyed: 4465,

		MaxTimePortalHeld:     59,
		MaxTimeLinkMaintained: 55,
		MaxLinkLengthXDays:    2543,
		MaxTimeFieldHeld:      52,
		LargestFieldMUsXDays:  2632,

		UniqueMissionsCompleted: 7,

		Hacks:           43350,
		GlyphHackPoints: 27254,

		InnovatorLevel: 11,
	},

	"zyp_v0_unknown-2.png": profile.Profile{
		Nick:  "zyp",
		Level: 9,
		AP:    2400020,

		UniquePortalsVisited: 782,
		XMCollected:          6440635,

		DistanceWalked: 137,

		ResonatorsDeployed:     3560,
		LinksCreated:           1116,
		ControlFieldsCreated:   528,
		MindUnitsCaptured:      30116,
		LongestLinkEverCreated: 97,
		LargestControlField:    7221,
		XMRecharged:            3259370,
		PortalsCaptured:        471,
		UniquePortalsCaptured:  322,

		ResonatorsDestroyed:         3163,
		PortalsNeutralized:          355,
		EnemyLinksDestroyed:         866,
		EnemyControlFieldsDestroyed: 371,

		MaxTimePortalHeld:     32,
		MaxTimeLinkMaintained: 25,
		MaxLinkLengthXDays:    83,
		MaxTimeFieldHeld:      15,
		LargestFieldMUsXDays:  3117,

		Hacks: 3077,
	},

	"wexp_v0_unknown-2.png": profile.Profile{
		Nick:  "wexp",
		Level: 8,
		AP:    1788413,

		UniquePortalsVisited: 271,
		XMCollected:          9425170,

		DistanceWalked: 195,

		ResonatorsDeployed:     2685,
		LinksCreated:           796,
		ControlFieldsCreated:   424,
		MindUnitsCaptured:      31119,
		LongestLinkEverCreated: 31,
		LargestControlField:    9875,
		XMRecharged:            5257960,
		PortalsCaptured:        246,
		UniquePortalsCaptured:  125,

		ResonatorsDestroyed:         1475,
		PortalsNeutralized:          223,
		EnemyLinksDestroyed:         345,
		EnemyControlFieldsDestroyed: 175,

		MaxTimePortalHeld:     26,
		MaxTimeLinkMaintained: 19,
		MaxLinkLengthXDays:    114,
		MaxTimeFieldHeld:      16,
		LargestFieldMUsXDays:  9695,

		Hacks: 3299,
	},

	"greenduck_v0_unknown.png": profile.Profile{
		Nick:  "Greenduck",
		Level: 15,
		AP:    41899205,

		UniquePortalsVisited: 1439,
		PortalsDiscovered:    84,
		XMCollected:          128676839,

		DistanceWalked: 2422,

		ResonatorsDeployed:     65685,
		LinksCreated:           22831,
		ControlFieldsCreated:   9977,
		MindUnitsCaptured:      149265,
		LongestLinkEverCreated: 253,
		LargestControlField:    851,
		XMRecharged:            26909815,
		PortalsCaptured:        5334,
		UniquePortalsCaptured:  781,
		ModsDeployed:           311,

		ResonatorsDestroyed:         50650,
		PortalsNeutralized:          6044,
		EnemyLinksDestroyed:         9434,
		EnemyControlFieldsDestroyed: 4665,

		MaxTimePortalHeld:     84,
		MaxTimeLinkMaintained: 52,
		MaxLinkLengthXDays:    95,
		MaxTimeFieldHeld:      52,
		LargestFieldMUsXDays:  1873,

		UniqueMissionsCompleted: 1,

		Hacks: 66036,

		InnovatorLevel: 13,
	},

	"panoptos_v0_unknown.jpeg": profile.Profile{
		Nick:  "panoptos",
		Level: 11,
		AP:    11543207,

		UniquePortalsVisited: 2442,
		XMCollected:          32191703,

		DistanceWalked: 549,

		ResonatorsDeployed:     16865,
		LinksCreated:           2816,
		ControlFieldsCreated:   1139,
		MindUnitsCaptured:      18063,
		LongestLinkEverCreated: 17,
		LargestControlField:    987,
		XMRecharged:            9847371,
		PortalsCaptured:        2386,
		UniquePortalsCaptured:  1233,
		ModsDeployed:           608,

		ResonatorsDestroyed:         18522,
		PortalsNeutralized:          2524,
		EnemyLinksDestroyed:         4023,
		EnemyControlFieldsDestroyed: 1937,

		MaxTimePortalHeld:     51,
		MaxTimeLinkMaintained: 41,
		MaxLinkLengthXDays:    51,
		MaxTimeFieldHeld:      41,
		LargestFieldMUsXDays:  2839,

		Hacks: 11564,

		InnovatorLevel: 9,
	},

	"steino_v0_unknown-3.png": profile.Profile{
		Nick:  "steino",
		Level: 8,
		AP:    1667211,

		UniquePortalsVisited: 780,
		PortalsDiscovered:    4,
		XMCollected:          4524546,

		DistanceWalked: 184,

		ResonatorsDeployed:     3502,
		LinksCreated:           708,
		ControlFieldsCreated:   368,
		MindUnitsCaptured:      10745,
		LongestLinkEverCreated: 105,
		LargestControlField:    378,
		XMRecharged:            897264,
		PortalsCaptured:        280,
		UniquePortalsCaptured:  102,
		ModsDeployed:           8,

		ResonatorsDestroyed:         1786,
		PortalsNeutralized:          280,
		EnemyLinksDestroyed:         284,
		EnemyControlFieldsDestroyed: 136,

		MaxTimePortalHeld:     49,
		MaxTimeLinkMaintained: 33,
		MaxLinkLengthXDays:    10,
		MaxTimeFieldHeld:      16,
		LargestFieldMUsXDays:  760,

		Hacks: 2908,

		InnovatorLevel: 3,
	},

	"haste_v0_nexus6-6.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    31565998,

		UniquePortalsVisited: 3843,
		PortalsDiscovered:    40,
		XMCollected:          130451374,

		DistanceWalked: 2344,

		ResonatorsDeployed:     66444,
		LinksCreated:           5944,
		ControlFieldsCreated:   2900,
		MindUnitsCaptured:      33752,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            28103024,
		PortalsCaptured:        7505,
		UniquePortalsCaptured:  1537,
		ModsDeployed:           1950,

		ResonatorsDestroyed:         54700,
		PortalsNeutralized:          8035,
		EnemyLinksDestroyed:         11241,
		EnemyControlFieldsDestroyed: 5562,

		MaxTimePortalHeld:     184,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 24,

		Hacks:                  60282,
		GlyphHackPoints:        12116,
		ConsecutiveDaysHacking: 13,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus6-7.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    31728659,

		UniquePortalsVisited: 3845,
		PortalsDiscovered:    40,
		XMCollected:          130836794,

		DistanceWalked: 2354,

		ResonatorsDeployed:     66685,
		LinksCreated:           5947,
		ControlFieldsCreated:   2901,
		MindUnitsCaptured:      33775,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            28106324,
		PortalsCaptured:        7541,
		UniquePortalsCaptured:  1537,
		ModsDeployed:           1972,

		ResonatorsDestroyed:         55023,
		PortalsNeutralized:          8073,
		EnemyLinksDestroyed:         11343,
		EnemyControlFieldsDestroyed: 5619,

		MaxTimePortalHeld:     186,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 24,

		Hacks:                  60412,
		GlyphHackPoints:        12296,
		ConsecutiveDaysHacking: 14,
	},

	"viddy_v0_unknown-2.png": profile.Profile{
		Nick:  "viddy",
		Level: 13,
		AP:    14536389,

		UniquePortalsVisited: 2775,
		PortalsDiscovered:    15,
		XMCollected:          87329517,

		DistanceWalked: 1231,

		ResonatorsDeployed:     33868,
		LinksCreated:           3050,
		ControlFieldsCreated:   1179,
		MindUnitsCaptured:      14415,
		LongestLinkEverCreated: 166,
		LargestControlField:    1250,
		XMRecharged:            28380817,
		PortalsCaptured:        4915,
		UniquePortalsCaptured:  1355,

		ResonatorsDestroyed:         32313,
		PortalsNeutralized:          4420,
		EnemyLinksDestroyed:         5132,
		EnemyControlFieldsDestroyed: 2434,

		MaxTimePortalHeld:     151,
		MaxTimeLinkMaintained: 67,
		MaxLinkLengthXDays:    98,
		MaxTimeFieldHeld:      64,
		LargestFieldMUsXDays:  5153,

		Hacks: 33814,
	},

	"jonheier_v0_unknown.png": profile.Profile{
		Nick:  "jonheier",
		Level: 10,
		AP:    5909679,

		UniquePortalsVisited: 1446,
		PortalsDiscovered:    1,
		XMCollected:          25198081,

		DistanceWalked: 645,

		ResonatorsDeployed:     12275,
		LinksCreated:           1716,
		ControlFieldsCreated:   1061,
		MindUnitsCaptured:      49393,
		LongestLinkEverCreated: 9,
		LargestControlField:    19962,
		XMRecharged:            11079834,
		PortalsCaptured:        1493,
		UniquePortalsCaptured:  410,

		ResonatorsDestroyed:         10617,
		PortalsNeutralized:          1631,
		EnemyLinksDestroyed:         1484,
		EnemyControlFieldsDestroyed: 605,

		MaxTimePortalHeld:     120,
		MaxTimeLinkMaintained: 62,
		MaxLinkLengthXDays:    22,
		MaxTimeFieldHeld:      55,
		LargestFieldMUsXDays:  1323,

		Hacks: 9299,
	},

	"arneaase_v0_unknown.png": profile.Profile{
		Nick:  "arneaase",
		Level: 10,
		AP:    9476114,

		UniquePortalsVisited: 967,
		PortalsDiscovered:    17,
		XMCollected:          43025847,

		DistanceWalked: 774,

		ResonatorsDeployed:     21602,
		LinksCreated:           2211,
		ControlFieldsCreated:   790,
		MindUnitsCaptured:      19942,
		LongestLinkEverCreated: 625,
		LargestControlField:    10197,
		XMRecharged:            10836742,
		PortalsCaptured:        2074,
		UniquePortalsCaptured:  586,

		ResonatorsDestroyed:         15706,
		PortalsNeutralized:          2057,
		EnemyLinksDestroyed:         3243,
		EnemyControlFieldsDestroyed: 1724,

		MaxTimePortalHeld:     231,
		MaxTimeLinkMaintained: 16,
		MaxLinkLengthXDays:    4354,
		MaxTimeFieldHeld:      12,
		LargestFieldMUsXDays:  1117,

		Hacks: 23987,
	},

	"zyp_v0_unknown-3.png": profile.Profile{
		Nick:  "zyp",
		Level: 10,
		AP:    4131752,

		UniquePortalsVisited: 2035,
		XMCollected:          12116116,

		DistanceWalked: 253,

		ResonatorsDeployed:     5943,
		LinksCreated:           1383,
		ControlFieldsCreated:   656,
		MindUnitsCaptured:      80114,
		LongestLinkEverCreated: 172,
		LargestControlField:    24772,
		XMRecharged:            5548413,
		PortalsCaptured:        933,
		UniquePortalsCaptured:  680,
		ModsDeployed:           250,

		ResonatorsDestroyed:         5817,
		PortalsNeutralized:          755,
		EnemyLinksDestroyed:         1553,
		EnemyControlFieldsDestroyed: 722,

		MaxTimePortalHeld:     72,
		MaxTimeLinkMaintained: 25,
		MaxLinkLengthXDays:    180,
		MaxTimeFieldHeld:      15,
		LargestFieldMUsXDays:  3117,

		UniqueMissionsCompleted: 10,

		Hacks:                  4987,
		GlyphHackPoints:        52,
		ConsecutiveDaysHacking: 0,

		InnovatorLevel: 9,
	},

	"codemonkey_v0_unknown.png": profile.Profile{
		Nick:  "codemonkey",
		Level: 6,
		AP:    417574,

		UniquePortalsVisited: 438,
		XMCollected:          451823,

		DistanceWalked: 18,

		ResonatorsDeployed:     1240,
		LinksCreated:           256,
		ControlFieldsCreated:   30,
		MindUnitsCaptured:      13007,
		LongestLinkEverCreated: 6,
		LargestControlField:    8158,
		XMRecharged:            51075,
		PortalsCaptured:        102,
		UniquePortalsCaptured:  83,

		ResonatorsDestroyed:         373,
		PortalsNeutralized:          36,
		EnemyLinksDestroyed:         48,
		EnemyControlFieldsDestroyed: 18,

		MaxTimePortalHeld:     18,
		MaxTimeLinkMaintained: 3,
		MaxLinkLengthXDays:    0,
		MaxTimeFieldHeld:      0,
		LargestFieldMUsXDays:  12,

		Hacks: 1127,
	},

	"jinpark_v0_unknown.jpeg": profile.Profile{
		Nick:  "jinpark",
		Level: 2,
		AP:    5439,

		UniquePortalsVisited: 19,
		XMCollected:          11638,

		DistanceWalked: 1,

		ResonatorsDeployed:     18,
		LinksCreated:           3,
		LongestLinkEverCreated: 1,
		PortalsCaptured:        1,
		UniquePortalsCaptured:  1,

		MaxTimePortalHeld:     6,
		MaxTimeLinkMaintained: 0,
		MaxLinkLengthXDays:    0,

		Hacks: 25,
	},

	"erebwain_v0_unknown.png": profile.Profile{
		Nick:  "erebwain",
		Level: 14,
		AP:    23028553,

		UniquePortalsVisited: 2458,
		PortalsDiscovered:    80,
		XMCollected:          80835619,

		DistanceWalked: 1927,

		ResonatorsDeployed:     45545,
		LinksCreated:           5525,
		ControlFieldsCreated:   2869,
		MindUnitsCaptured:      43460,
		LongestLinkEverCreated: 361,
		LargestControlField:    1227,
		XMRecharged:            20071135,
		PortalsCaptured:        4786,
		UniquePortalsCaptured:  1199,
		ModsDeployed:           140,

		ResonatorsDestroyed:         34862,
		PortalsNeutralized:          4199,
		EnemyLinksDestroyed:         8410,
		EnemyControlFieldsDestroyed: 4305,

		MaxTimePortalHeld:     167,
		MaxTimeLinkMaintained: 28,
		MaxLinkLengthXDays:    442,
		MaxTimeFieldHeld:      27,
		LargestFieldMUsXDays:  1687,

		UniqueMissionsCompleted: 1,

		Hacks: 48764,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus6-8.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    27422228,

		UniquePortalsVisited: 3633,
		PortalsDiscovered:    40,
		XMCollected:          114908313,

		DistanceWalked: 2137,

		ResonatorsDeployed:     58270,
		LinksCreated:           5442,
		ControlFieldsCreated:   2633,
		MindUnitsCaptured:      32961,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            27555026,
		PortalsCaptured:        6671,
		UniquePortalsCaptured:  1441,
		ModsDeployed:           1004,

		ResonatorsDestroyed:         48047,
		PortalsNeutralized:          7117,
		EnemyLinksDestroyed:         9339,
		EnemyControlFieldsDestroyed: 4496,

		MaxTimePortalHeld:     145,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 7,

		Hacks:           55088,
		GlyphHackPoints: 1193,

		InnovatorLevel: 13,
	},

	"forferdet_v0_unknown-4.png": profile.Profile{
		Nick:  "forferdet",
		Level: 14,
		AP:    24469249,

		UniquePortalsVisited: 1677,
		PortalsDiscovered:    13,
		XMCollected:          94900504,

		DistanceWalked: 1350,

		ResonatorsDeployed:     40648,
		LinksCreated:           6950,
		ControlFieldsCreated:   3384,
		MindUnitsCaptured:      23695,
		LongestLinkEverCreated: 156,
		LargestControlField:    977,
		XMRecharged:            33450470,
		PortalsCaptured:        4297,
		UniquePortalsCaptured:  1027,
		ModsDeployed:           1101,

		ResonatorsDestroyed:         34381,
		PortalsNeutralized:          4761,
		EnemyLinksDestroyed:         7000,
		EnemyControlFieldsDestroyed: 3466,

		MaxTimePortalHeld:     53,
		MaxTimeLinkMaintained: 55,
		MaxLinkLengthXDays:    2543,
		MaxTimeFieldHeld:      52,
		LargestFieldMUsXDays:  2632,

		Hacks: 36483,

		InnovatorLevel: 11,
	},

	"haste_v0_nexus5-26.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    25243305,

		UniquePortalsVisited: 3505,
		PortalsDiscovered:    36,
		XMCollected:          104724752,

		DistanceWalked: 2003,

		ResonatorsDeployed:     52943,
		LinksCreated:           5230,
		ControlFieldsCreated:   2532,
		MindUnitsCaptured:      32783,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26381951,
		PortalsCaptured:        6199,
		UniquePortalsCaptured:  1407,
		ModsDeployed:           244,

		ResonatorsDestroyed:         44718,
		PortalsNeutralized:          6614,
		EnemyLinksDestroyed:         8419,
		EnemyControlFieldsDestroyed: 3960,

		MaxTimePortalHeld:     111,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 51480,

		InnovatorLevel: 13,
	},

	"yggdar_v0_unknown.png": profile.Profile{
		Nick:  "Yggdar",
		Level: 6,
		AP:    324558,

		UniquePortalsVisited: 174,
		XMCollected:          984574,

		DistanceWalked: 23,

		ResonatorsDeployed:     732,
		LinksCreated:           349,
		ControlFieldsCreated:   44,
		MindUnitsCaptured:      150,
		LongestLinkEverCreated: 1,
		LargestControlField:    76,
		XMRecharged:            400247,
		PortalsCaptured:        53,
		UniquePortalsCaptured:  37,
		ModsDeployed:           14,

		ResonatorsDestroyed:         224,
		PortalsNeutralized:          10,
		EnemyLinksDestroyed:         16,
		EnemyControlFieldsDestroyed: 4,

		MaxTimePortalHeld:     10,
		MaxTimeLinkMaintained: 10,
		MaxLinkLengthXDays:    0,
		MaxTimeFieldHeld:      10,
		LargestFieldMUsXDays:  12,

		Hacks: 1210,

		InnovatorLevel: 3,
	},

	"tufte_v0_iphone-18.jpeg": profile.Profile{
		Nick:  "Tufte",
		Level: 11,
		AP:    8590390,

		UniquePortalsVisited: 2031,
		XMCollected:          41159953,

		DistanceWalked: 691,

		ResonatorsDeployed:     13679,
		LinksCreated:           2053,
		ControlFieldsCreated:   1194,
		MindUnitsCaptured:      624168,
		LongestLinkEverCreated: 304,
		LargestControlField:    610877,
		XMRecharged:            19327328,
		PortalsCaptured:        1331,
		UniquePortalsCaptured:  720,
		ModsDeployed:           1326,

		ResonatorsDestroyed:         10393,
		PortalsNeutralized:          1287,
		EnemyLinksDestroyed:         2365,
		EnemyControlFieldsDestroyed: 1296,

		MaxTimePortalHeld:     44,
		MaxTimeLinkMaintained: 23,
		MaxLinkLengthXDays:    274,
		MaxTimeFieldHeld:      20,
		LargestFieldMUsXDays:  354224,

		UniqueMissionsCompleted: 13,

		Hacks:                  10383,
		GlyphHackPoints:        1484,
		ConsecutiveDaysHacking: 6,

		AgentsSuccessfullyRecruited: 1,

		InnovatorLevel: 3,
	},

	"leiasqz_v0_unknown.png": profile.Profile{
		Nick:  "leiasqz",
		Level: 8,
		AP:    2238929,

		UniquePortalsVisited: 464,
		XMCollected:          4779263,

		DistanceWalked: 93,

		ResonatorsDeployed:     5063,
		LinksCreated:           1917,
		ControlFieldsCreated:   262,
		MindUnitsCaptured:      15592,
		LongestLinkEverCreated: 19,
		LargestControlField:    3705,
		XMRecharged:            1739218,
		PortalsCaptured:        464,
		UniquePortalsCaptured:  195,
		ModsDeployed:           381,

		ResonatorsDestroyed:         4426,
		PortalsNeutralized:          245,
		EnemyLinksDestroyed:         1036,
		EnemyControlFieldsDestroyed: 311,

		MaxTimePortalHeld:     5,
		MaxTimeLinkMaintained: 7,
		MaxLinkLengthXDays:    1,
		MaxTimeFieldHeld:      7,
		LargestFieldMUsXDays:  77,

		UniqueMissionsCompleted: 2,

		Hacks:                  4908,
		GlyphHackPoints:        359,
		ConsecutiveDaysHacking: 13,
	},

	"haste_v0_nexus6-9.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    27422228,

		UniquePortalsVisited: 3633,
		PortalsDiscovered:    40,
		XMCollected:          114908313,

		DistanceWalked: 2137,

		ResonatorsDeployed:     58270,
		LinksCreated:           5442,
		ControlFieldsCreated:   2633,
		MindUnitsCaptured:      32961,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            27555026,
		PortalsCaptured:        6671,
		UniquePortalsCaptured:  1441,
		ModsDeployed:           1004,

		ResonatorsDestroyed:         48047,
		PortalsNeutralized:          7117,
		EnemyLinksDestroyed:         9339,
		EnemyControlFieldsDestroyed: 4496,

		MaxTimePortalHeld:     145,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 7,

		Hacks:           55088,
		GlyphHackPoints: 1193,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus6-10.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    27040256,

		UniquePortalsVisited: 3633,
		PortalsDiscovered:    39,
		XMCollected:          113605960,

		DistanceWalked: 2119,

		ResonatorsDeployed:     57407,
		LinksCreated:           5407,
		ControlFieldsCreated:   2613,
		MindUnitsCaptured:      32935,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            27486162,
		PortalsCaptured:        6569,
		UniquePortalsCaptured:  1441,
		ModsDeployed:           924,

		ResonatorsDestroyed:         47385,
		PortalsNeutralized:          7017,
		EnemyLinksDestroyed:         9169,
		EnemyControlFieldsDestroyed: 4403,

		MaxTimePortalHeld:     141,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 4,

		Hacks:           54682,
		GlyphHackPoints: 367,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus6-11.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    28397342,

		UniquePortalsVisited: 3685,
		PortalsDiscovered:    40,
		XMCollected:          118583310,

		DistanceWalked: 2182,

		ResonatorsDeployed:     60575,
		LinksCreated:           5636,
		ControlFieldsCreated:   2734,
		MindUnitsCaptured:      33204,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            27608651,
		PortalsCaptured:        6818,
		UniquePortalsCaptured:  1469,
		ModsDeployed:           1245,

		ResonatorsDestroyed:         49205,
		PortalsNeutralized:          7267,
		EnemyLinksDestroyed:         9712,
		EnemyControlFieldsDestroyed: 4717,

		MaxTimePortalHeld:     153,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		UniqueMissionsCompleted: 10,

		Hacks:           56372,
		GlyphHackPoints: 4097,

		InnovatorLevel: 13,
	},

	"haste_v0_nexus6-12.png": profile.Profile{
		Nick:  "haste",
		Level: 15,
		AP:    26230336,

		UniquePortalsVisited: 3554,
		PortalsDiscovered:    37,
		XMCollected:          109539190,

		DistanceWalked: 2062,

		ResonatorsDeployed:     55272,
		LinksCreated:           5345,
		ControlFieldsCreated:   2576,
		MindUnitsCaptured:      32866,
		LongestLinkEverCreated: 163,
		LargestControlField:    19445,
		XMRecharged:            26539748,
		PortalsCaptured:        6367,
		UniquePortalsCaptured:  1421,
		ModsDeployed:           653,

		ResonatorsDestroyed:         45849,
		PortalsNeutralized:          6782,
		EnemyLinksDestroyed:         8815,
		EnemyControlFieldsDestroyed: 4228,

		MaxTimePortalHeld:     129,
		MaxTimeLinkMaintained: 61,
		MaxLinkLengthXDays:    15,
		MaxTimeFieldHeld:      61,
		LargestFieldMUsXDays:  11600,

		Hacks: 53550,

		InnovatorLevel: 13,
	},

	"nittingduck_v0_unknown.png": profile.Profile{
		Nick:  "nittingduck",
		Level: 11,
		AP:    10670514,

		UniquePortalsVisited: 1336,
		XMCollected:          49052789,

		DistanceWalked: 986,

		ResonatorsDeployed:     15508,
		LinksCreated:           3352,
		ControlFieldsCreated:   1665,
		MindUnitsCaptured:      656684,
		LongestLinkEverCreated: 304,
		LargestControlField:    610289,
		XMRecharged:            18927968,
		PortalsCaptured:        1639,
		UniquePortalsCaptured:  620,
		ModsDeployed:           1874,

		ResonatorsDestroyed:         11331,
		PortalsNeutralized:          1487,
		EnemyLinksDestroyed:         2954,
		EnemyControlFieldsDestroyed: 1524,

		MaxTimePortalHeld:     51,
		MaxTimeLinkMaintained: 38,
		MaxLinkLengthXDays:    281,
		MaxTimeFieldHeld:      27,
		LargestFieldMUsXDays:  354179,

		UniqueMissionsCompleted: 11,

		Hacks:                  12688,
		GlyphHackPoints:        13381,
		ConsecutiveDaysHacking: 14,

		InnovatorLevel: 3,
	},

	"nittingduck_v0_unknown-1.png": profile.Profile{
		Nick:  "nittingduck",
		Level: 11,
		AP:    10368373,

		UniquePortalsVisited: 1332,
		XMCollected:          47138634,

		DistanceWalked: 971,

		ResonatorsDeployed:     15138,
		LinksCreated:           3263,
		ControlFieldsCreated:   1626,
		MindUnitsCaptured:      649448,
		LongestLinkEverCreated: 304,
		LargestControlField:    610289,
		XMRecharged:            18001208,
		PortalsCaptured:        1597,
		UniquePortalsCaptured:  615,
		ModsDeployed:           1817,

		ResonatorsDestroyed:         11004,
		PortalsNeutralized:          1448,
		EnemyLinksDestroyed:         2885,
		EnemyControlFieldsDestroyed: 1498,

		MaxTimePortalHeld:     43,
		MaxTimeLinkMaintained: 38,
		MaxLinkLengthXDays:    281,
		MaxTimeFieldHeld:      27,
		LargestFieldMUsXDays:  354179,

		UniqueMissionsCompleted: 11,

		Hacks:                  12381,
		GlyphHackPoints:        11833,
		ConsecutiveDaysHacking: 5,
	},

	"hbrx_v0_unknown.png": profile.Profile{
		Nick:  "hbrx",
		Level: 15,
		AP:    29608688,

		UniquePortalsVisited: 2162,
		PortalsDiscovered:    6,
		XMCollected:          100096614,

		DistanceWalked: 2181,

		ResonatorsDeployed:     55549,
		LinksCreated:           7627,
		ControlFieldsCreated:   2926,
		MindUnitsCaptured:      56930,
		LongestLinkEverCreated: 128,
		LargestControlField:    8756,
		XMRecharged:            25872226,
		PortalsCaptured:        8051,
		UniquePortalsCaptured:  1526,
		ModsDeployed:           1269,

		ResonatorsDestroyed:         62013,
		PortalsNeutralized:          7760,
		EnemyLinksDestroyed:         12573,
		EnemyControlFieldsDestroyed: 6024,

		MaxTimePortalHeld:     149,
		MaxTimeLinkMaintained: 62,
		MaxLinkLengthXDays:    478,
		MaxTimeFieldHeld:      58,
		LargestFieldMUsXDays:  13797,

		UniqueMissionsCompleted: 5,

		Hacks:                  51421,
		GlyphHackPoints:        16995,
		ConsecutiveDaysHacking: 34,

		InnovatorLevel: 13,
	},
}

var (
	vTop            bool
	vNick           bool
	vLevel          bool
	vAP             bool
	vBottom         bool
	expensive       bool
	overrideProfile int
)

func init() {
	conf.Load("../config.json")
	InitConfig()

	model.Init()

	flag.BoolVar(&vTop, "top", false, "Validate validate top part of OCR image.")
	flag.BoolVar(&vNick, "nick", false, "Validate validate nick part of OCR image.")
	flag.BoolVar(&vLevel, "level", false, "Validate validate level part of OCR image.")
	flag.BoolVar(&vAP, "ap", false, "Validate validate ap part of OCR image.")
	flag.BoolVar(&vBottom, "bottom", false, "Validate validate bottom part of OCR image.")
	flag.BoolVar(&expensive, "expensive", false, "Only run expensive tests.")
	flag.IntVar(&overrideProfile, "profile", 0, "Profile to use.")

	flag.Parse()
}

func validateNick(t *testing.T, file string, e profile.Profile, p profile.Profile) {
	if p.Nick != e.Nick {
		t.Errorf("%s: .Nick: Got %v Expected %v", file, p.Nick, e.Nick)
	}
}

func validateLevel(t *testing.T, file string, e profile.Profile, p profile.Profile) {
	if p.Level != e.Level {
		t.Errorf("%s: .Level: Got %v Expected %v", file, p.Level, e.Level)
	}
}

func validateAP(t *testing.T, file string, e profile.Profile, p profile.Profile) {
	if p.AP != e.AP {
		t.Errorf("%s: .AP: Got %v Expected %v", file, p.AP, e.AP)
	}
}

func validateTop(t *testing.T, file string, e profile.Profile, p profile.Profile) {
	validateNick(t, file, e, p)
	validateLevel(t, file, e, p)
	validateAP(t, file, e, p)
}

func validateBottom(t *testing.T, file string, e profile.Profile, p profile.Profile) {
	if p.UniquePortalsVisited != e.UniquePortalsVisited {
		t.Errorf("%s: .UniquePortalsVisited: Got %v Expected %v", file, p.UniquePortalsVisited, e.UniquePortalsVisited)
	}

	if p.PortalsDiscovered != e.PortalsDiscovered {
		t.Errorf("%s: .PortalsDiscovered: Got %v Expected %v", file, p.PortalsDiscovered, e.PortalsDiscovered)
	}

	if p.XMCollected != e.XMCollected {
		t.Errorf("%s: .XMCollected: Got %v Expected %v", file, p.XMCollected, e.XMCollected)
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

	if p.InnovatorLevel != e.InnovatorLevel {
		t.Errorf("%s: .InnovatorLevel: Got %v Expected %v", file, p.InnovatorLevel, e.InnovatorLevel)
	}

	if p.Hacks != e.Hacks {
		t.Errorf("%s: .Hacks: Got %v Expected %v", file, p.Hacks, e.Hacks)
	}

	if p.GlyphHackPoints != e.GlyphHackPoints {
		t.Errorf("%s: .GlyphHackPoints: Got %v Expected %v", file, p.GlyphHackPoints, e.GlyphHackPoints)
	}

	if p.UniqueMissionsCompleted != e.UniqueMissionsCompleted {
		t.Errorf("%s: .UniqueMissionsCompleted: Got %v Expected %v", file, p.UniqueMissionsCompleted, e.UniqueMissionsCompleted)
	}

	if p.AgentsSuccessfullyRecruited != e.AgentsSuccessfullyRecruited {
		t.Errorf("%s: .AgentsSuccessfullyRecruited: Got %v Expected %v", file, p.AgentsSuccessfullyRecruited, e.AgentsSuccessfullyRecruited)
	}
}

// Test wrapper
func w(t *testing.T, file string) {
	if !testing.Short() {
		t.Skip("skipping test in expensive mode.")
	}

	conf.Config.Cache = "testdata/"

	t.Parallel()

	p := New(file, overrideProfile)
	p.Split()

	e := testData[file]

	if vTop || vNick {
		p.ProcessNick()
		validateNick(t, file, e, p.Profile)
	}

	if vTop || vLevel {
		p.ProcessLevel()
		validateLevel(t, file, e, p.Profile)
	}

	if vTop || vAP {
		p.ProcessAP()
		validateAP(t, file, e, p.Profile)
	}

	if vBottom {
		p.ProcessBottom()
		p.ProcessInnovator()
		validateBottom(t, file, e, p.Profile)
	}

	p.CleanUp()
}

func rowToQueueProfile(rows *sql.Rows) userQueueEntry {
	var e userQueueEntry
	err := rows.Scan(
		&e.Id, &e.File, &e.OCRProfile, &e.Processed, &e.Processtime,

		&e.Profile.Id, &e.Profile.Timestamp, &e.Profile.Nick, &e.Profile.Level,
		&e.Profile.AP,

		&e.Profile.UniquePortalsVisited, &e.Profile.PortalsDiscovered,
		&e.Profile.XMCollected,

		&e.Profile.DistanceWalked,

		&e.Profile.ResonatorsDeployed, &e.Profile.LinksCreated,
		&e.Profile.ControlFieldsCreated, &e.Profile.MindUnitsCaptured,
		&e.Profile.LongestLinkEverCreated, &e.Profile.LargestControlField,
		&e.Profile.XMRecharged, &e.Profile.PortalsCaptured,
		&e.Profile.UniquePortalsCaptured, &e.Profile.ModsDeployed,

		&e.Profile.ResonatorsDestroyed, &e.Profile.PortalsNeutralized,
		&e.Profile.EnemyLinksDestroyed, &e.Profile.EnemyControlFieldsDestroyed,

		&e.Profile.MaxTimePortalHeld, &e.Profile.MaxTimeLinkMaintained,
		&e.Profile.MaxLinkLengthXDays, &e.Profile.MaxTimeFieldHeld,
		&e.Profile.LargestFieldMUsXDays,

		&e.Profile.UniqueMissionsCompleted,

		&e.Profile.Hacks, &e.Profile.GlyphHackPoints,
		&e.Profile.ConsecutiveDaysHacking,

		&e.Profile.AgentsSuccessfullyRecruited,

		&e.Profile.InnovatorLevel,
	)

	if err != nil {
		fmt.Println(err)
	}

	return e
}

func TestOCRAllExpensive(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	rows := model.GetAllUserQueuesWithProfiles()

	for rows.Next() {
		e := rowToQueueProfile(rows)

		fmt.Println("Testing ", e.File)

		p := New(e.File, e.OCRProfile)
		p.Split()
		p.Process()

		validateTop(t, e.File, e.Profile, p.Profile)
		validateBottom(t, e.File, e.Profile, p.Profile)

		p.CleanUp()
	}
}

func TestOCR_480x_wexp_v0_unknown(t *testing.T) {
	w(t, "wexp_v0_unknown.png")
}

func TestOCR_480x_wexp_v0_unknown_1(t *testing.T) {
	// Invalid Max Link Length x Days
	w(t, "wexp_v0_unknown-1.png")
}

func TestOCR_480x_wexp_v0_unknown_2(t *testing.T) {
	// Invalid ap
	w(t, "wexp_v0_unknown-2.png")
}

func TestOCR_640x_mvnch_v1640_iphone5(t *testing.T) {
	w(t, "mvnch_v1640_iphone5.jpeg")
}

func TestOCR_640x_mvnch_v1640_iphone5_1(t *testing.T) {
	w(t, "mvnch_v1640_iphone5-1.jpeg")
}

func TestOCR_640x_mvnch_vunknown_iphone5(t *testing.T) {
	// Invalid nick and AP.
	w(t, "mvnch_vunknown_iphone5.jpeg")
}

func TestOCR_640x_mvnch_vunknown_iphone5_1(t *testing.T) {
	// Invalid unique portals visited
	w(t, "mvnch_vunknown_iphone5-1.jpeg")
}

func TestOCR_640x_Scissorhill_v0_iphone(t *testing.T) {
	w(t, "scissorhill_v0_iphone.jpeg")
}

func TestOCR_640x_Tufte_v1660_iphone(t *testing.T) {
	w(t, "tufte_v1660_iphone.jpeg")
}

func TestOCR_640x_Tufte_v1670_iphone(t *testing.T) {
	w(t, "tufte_v1670_iphone.jpeg")
}

func TestOCR_640x_Tufte_v1700_iphone(t *testing.T) {
	w(t, "tufte_v1700_iphone.jpeg")
}

func TestOCR_640x_Tufte_v1700_iphone_1(t *testing.T) {
	w(t, "tufte_v1700_iphone-1.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone(t *testing.T) {
	w(t, "tufte_v0_iphone.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_1(t *testing.T) {
	// Invalid portals captured and neutralized
	w(t, "tufte_v0_iphone-1.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_2(t *testing.T) {
	// Invalid level
	w(t, "tufte_v0_iphone-2.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_3(t *testing.T) {
	// Invalid AP
	w(t, "tufte_v0_iphone-3.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_4(t *testing.T) {
	// Invalid level
	w(t, "tufte_v0_iphone-4.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_5(t *testing.T) {
	// Invalid level, ap, resonators deployed, unique portals captured
	w(t, "tufte_v0_iphone-5.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_6(t *testing.T) {
	// Invalid level, largest control field, max time link maintained
	w(t, "tufte_v0_iphone-6.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_7(t *testing.T) {
	// Invalid mind units captured, enemy links destroyed, max link length x days
	w(t, "tufte_v0_iphone-7.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_8(t *testing.T) {
	// Invalid hacks
	w(t, "tufte_v0_iphone-8.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_9(t *testing.T) {
	// Invalid unique portals visited, mind units captured, xm recharged, enemy
	// control fields destroyed.
	w(t, "tufte_v0_iphone-9.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_10(t *testing.T) {
	// Invalid mind units captured
	w(t, "tufte_v0_iphone-10.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_11(t *testing.T) {
	// Invalid distance walked and glyph hack points.
	w(t, "tufte_v0_iphone-11.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_12(t *testing.T) {
	// Invalid mods deployed
	w(t, "tufte_v0_iphone-12.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_13(t *testing.T) {
	// Invalid unique portal visited, distance walked, control fields created,
	// portals captured, mods deployed and hacks.
	w(t, "tufte_v0_iphone-13.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_14(t *testing.T) {
	// Invalid xm collected.
	w(t, "tufte_v0_iphone-14.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_15(t *testing.T) {
	// Invalid xm collected
	w(t, "tufte_v0_iphone-15.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_16(t *testing.T) {
	// Invalid portals neutralized.
	w(t, "tufte_v0_iphone-16.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_17(t *testing.T) {
	// Invalid max time portal held, max time link maintained, max link length
	// x days.
	w(t, "tufte_v0_iphone-17.jpeg")
}

func TestOCR_640x_mvnch_v0_iphone_1(t *testing.T) {
	// Invalid enemy control fields destroyed.
	w(t, "mvnch_v0_iphone-1.jpeg")
}

func TestOCR_640x_mvnch_v0_iphone_2(t *testing.T) {
	// Invalid level.
	w(t, "mvnch_v0_iphone-2.jpeg")
}

func TestOCR_640x_panoptos_v0_unknown(t *testing.T) {
	// Invalid level.
	w(t, "panoptos_v0_unknown.jpeg")
}

func TestOCR_640x_jinpark_v0_unknown(t *testing.T) {
	// Invalid max time portal held.
	w(t, "jinpark_v0_unknown.jpeg")
}

func TestOCR_640x_Tufte_v0_iphone_18(t *testing.T) {
	// Invalid glyph hack points.
	w(t, "tufte_v0_iphone-18.jpeg")
}

func TestOCR_720x_sockerdricka_v0_unknown(t *testing.T) {
	w(t, "sockerdricka_v0_unknown.png")
}

func TestOCR_720x_arilds_v0_unknown(t *testing.T) {
	w(t, "arilds_v0_unknown.png")
}

func TestOCR_768x_oteckeh_v1630_nexus4(t *testing.T) {
	w(t, "oteckeh_v1630_nexus4.png")
}

func TestOCR_768x_oteckeh_v1690_nexus4(t *testing.T) {
	w(t, "oteckeh_v1690_nexus4.png")
}

func TestOCR_768x_zyp_v1630_unknown(t *testing.T) {
	w(t, "zyp_v1630_unknown.png")
}

func TestOCR_768x_zyp_v1690_unknown(t *testing.T) {
	w(t, "zyp_v1690_unknown.png")
}

func TestOCR_768x_zyp_v1690_unknown_1(t *testing.T) {
	w(t, "zyp_v1690_unknown-1.png")
}

func TestOCR_768x_zyp_v0_unknown(t *testing.T) {
	w(t, "zyp_v0_unknown.png")
}

func TestOCR_768x_zyp_v1700_unknown(t *testing.T) {
	w(t, "zyp_v1700_unknown.png")
}

func TestOCR_768x_zyp_v1711_unknown(t *testing.T) {
	w(t, "zyp_v1711_unknown.png")
}

func TestOCR_768x_zyp_v0_unknown_1(t *testing.T) {
	// Invalid nick and level.
	w(t, "zyp_v0_unknown-1.png")
}

func TestOCR_768x_zyp_v0_unknown_2(t *testing.T) {
	// Invalid nick.
	w(t, "zyp_v0_unknown-2.png")
}

func TestOCR_768x_zyp_v0_unknown_3(t *testing.T) {
	// Invalid nick.
	w(t, "zyp_v0_unknown-3.png")
}

func TestOCR_768x_codemonkey_v0_unknown(t *testing.T) {
	// Invalid level.
	w(t, "codemonkey_v0_unknown.png")
}

func TestOCR_1080x_steino_v0_unknown(t *testing.T) {
	// Invalid nick.
	w(t, "steino_v0_unknown.png")
}

func TestOCR_1080x_steino_v0_unknown_1(t *testing.T) {
	// Invalid AP
	w(t, "steino_v0_unknown-1.png")
}

func TestOCR_1080x_Bentis_v0_unknown_1(t *testing.T) {
	// Invalid nick
	w(t, "bentis_v0_unknown.png")
}

func TestOCR_1080x_haste_v1620_nexus5(t *testing.T) {
	w(t, "haste_v1620_nexus5.png")
}

func TestOCR_1080x_forferdet_v0_unknown(t *testing.T) {
	w(t, "forferdet_v0_unknown.png")
}

func TestOCR_1080x_forferdet_v0_unknown_1(t *testing.T) {
	w(t, "forferdet_v0_unknown-1.png")
}

func TestOCR_1080x_forferdet_v0_unknown_2(t *testing.T) {
	w(t, "forferdet_v0_unknown-2.png")
}

func TestOCR_1080x_forferdet_v0_unknown_3(t *testing.T) {
	w(t, "forferdet_v0_unknown-3.png")
}

func TestOCR_1080x_erebwain_v1620_s4active(t *testing.T) {
	w(t, "erebwain_v1620_s4active.png")
}

func TestOCR_1080x_erebwain_v1630_unknown(t *testing.T) {
	w(t, "erebwain_v1630_unknown.png")
}

func TestOCR_1080x_haste_v1630_nexus5(t *testing.T) {
	w(t, "haste_v1630_nexus5.png")
}

func TestOCR_1080x_haste_v1630_nexus5_1(t *testing.T) {
	w(t, "haste_v1630_nexus5-1.png")
}

func TestOCR_1080x_haste_v1660_nexus5(t *testing.T) {
	w(t, "haste_v1660_nexus5.png")
}

func TestOCR_1080x_haste_v1660_nexus5_1(t *testing.T) {
	w(t, "haste_v1660_nexus5-1.png")
}

func TestOCR_1080x_haste_v1670_nexus5(t *testing.T) {
	w(t, "haste_v1670_nexus5.png")
}

func TestOCR_1080x_mrwolfe_v1640_unknown(t *testing.T) {
	w(t, "mrwolfe_v1640_unknown.png")
}

func TestOCR_1080x_leiasqz_v1711_unknown(t *testing.T) {
	w(t, "leiasqz_v1711_unknown.png")
}

func TestOCR_1080x_forferdet_v1680_oneplusone(t *testing.T) {
	w(t, "forferdet_v1680_oneplusone.png")
}

func TestOCR_1080x_forferdet_v1700_oneplusone(t *testing.T) {
	w(t, "forferdet_v1700_oneplusone.png")
}

func TestOCR_1080x_madder79_v1680_unknown(t *testing.T) {
	w(t, "madder79_v1680_unknown.png")
}

func TestOCR_1080x_haste_v0_nexus5(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5.png")
}

func TestOCR_1080x_haste_v0_nexus5_1(t *testing.T) {
	// Invalid level.
	w(t, "haste_v0_nexus5-1.png")
}

func TestOCR_1080x_haste_v0_nexus5_2(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-2.png")
}

func TestOCR_1080x_haste_v0_nexus5_3(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-3.png")
}

func TestOCR_1080x_haste_v0_nexus5_4(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-4.png")
}

func TestOCR_1080x_haste_v0_nexus5_5(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-5.png")
}

func TestOCR_1080x_haste_v0_nexus5_6(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-6.png")
}

func TestOCR_1080x_haste_v0_nexus5_7(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-7.png")
}

func TestOCR_1080x_haste_v0_nexus5_8(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-8.png")
}

func TestOCR_1080x_haste_v0_nexus5_9(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-9.png")
}

func TestOCR_1080x_haste_v0_nexus5_10(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-10.png")
}

func TestOCR_1080x_haste_v0_nexus5_11(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-11.png")
}

func TestOCR_1080x_haste_v0_nexus5_12(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-12.png")
}

func TestOCR_1080x_haste_v0_nexus5_13(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-13.png")
}

func TestOCR_1080x_haste_v0_nexus5_14(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus5-14.png")
}

func TestOCR_1080x_haste_v0_nexus5_15(t *testing.T) {
	// Invalid AP. links created, control fields created, xm recharged, portals
	// captured, portals neutralized, enemy control fields destroyed, max time
	// link maintained, max time field held, largest field mus x days, hacks
	w(t, "haste_v0_nexus5-15.png")
}

func TestOCR_1080x_haste_v0_nexus5_16(t *testing.T) {
	// Invalid resonators destroyed.
	w(t, "haste_v0_nexus5-16.png")
}

func TestOCR_1080x_haste_v0_nexus5_17(t *testing.T) {
	// Invalid enemy control fields destroyed.
	w(t, "haste_v0_nexus5-17.png")
}

func TestOCR_1080x_haste_v0_nexus5_18(t *testing.T) {
	// Invalid mind units captured and largest field mus x days.
	w(t, "haste_v0_nexus5-18.png")
}

func TestOCR_1080x_viddy_v0_unknown(t *testing.T) {
	// Invalid control fields created.
	w(t, "viddy_v0_unknown.png")
}

func TestOCR_1080x_viddy_v0_unknown_1(t *testing.T) {
	// Invalid control fields created, portals neutralized, max time link
	// maintained.
	w(t, "viddy_v0_unknown-1.png")
}

func TestOCR_1080x_haste_v0_nexus5_19(t *testing.T) {
	// Invalid resonators destroyed.
	w(t, "haste_v0_nexus5-19.png")
}

func TestOCR_1080x_haste_v0_nexus5_20(t *testing.T) {
	// Invalid level.
	w(t, "haste_v0_nexus5-20.png")
}

func TestOCR_1080x_haste_v0_nexus5_21(t *testing.T) {
	// Invalid ap.
	w(t, "haste_v0_nexus5-21.png")
}

func TestOCR_1080x_haste_v0_nexus5_22(t *testing.T) {
	// Invalid distance walked, resonators deployed, mind units captured,
	// longest link ever created, resonators destroyed, enemy links destroyed,
	// max time field held, largest field mus x days, hacks.
	w(t, "haste_v0_nexus5-22.png")
}

func TestOCR_1080x_haste_v0_nexus5_23(t *testing.T) {
	// Invalid hacks.
	w(t, "haste_v0_nexus5-23.png")
}

func TestOCR_1080x_haste_v0_nexus5_24(t *testing.T) {
	// Invalid links created, control fields created, portals captured, portals
	// neutralized, enemy control fields destroyed, max time portal held,
	// largest fieild mus x days.
	w(t, "haste_v0_nexus5-24.png")
}

func TestOCR_1080x_haste_v0_nexus5_25(t *testing.T) {
	// Invalid enemy links destroyed.
	w(t, "haste_v0_nexus5-25.png")
}

func TestOCR_1080x_gnarfoz_v0_unknown(t *testing.T) {
	// Invalid mind units captured.
	w(t, "gnarfoz_v0_unknown.png")
}

func TestOCR_1080x_gnarfoz_v0_unknown_1(t *testing.T) {
	// Invalid resonators destroyed.
	w(t, "gnarfoz_v0_unknown-1.png")
}

func TestOCR_1080x_steino_v0_unknown_2(t *testing.T) {
	// Invalid max time field held.
	w(t, "steino_v0_unknown-2.png")
}

func TestOCR_1080x_Geenduck_v0_unknown(t *testing.T) {
	// Invalid nick.
	w(t, "greenduck_v0_unknown.png")
}

func TestOCR_1080x_steino_v0_unknown_3(t *testing.T) {
	// Invalid ap.
	w(t, "steino_v0_unknown-3.png")
}

func TestOCR_1080x_viddy_v0_unknown_2(t *testing.T) {
	// Invalid ap, max time field held.
	w(t, "viddy_v0_unknown-2.png")
}

func TestOCR_1080x_jonheier_v0_unknown(t *testing.T) {
	// Invalid nick, xm collected, portals neutralized.
	w(t, "jonheier_v0_unknown.png")
}

func TestOCR_1080x_arneaase_v0_unknown(t *testing.T) {
	// Invalid nick.
	w(t, "arneaase_v0_unknown.png")
}

func TestOCR_1080x_erebwain_v0_unknown(t *testing.T) {
	// Invalid longest link ever created.
	w(t, "erebwain_v0_unknown.png")
}

func TestOCR_1080x_forferdet_v0_unknown_4(t *testing.T) {
	// Invalid ap.
	w(t, "forferdet_v0_unknown-4.png")
}

func TestOCR_1080x_haste_v0_nexus5_26(t *testing.T) {
	// Invalid max time portal held.
	w(t, "haste_v0_nexus5-26.png")
}

func TestOCR_1080x_yggdar_v0_unknown(t *testing.T) {
	// Invalid ap.
	w(t, "yggdar_v0_unknown.png")
}

func TestOCR_1080x_leiasqz_v0_unknown(t *testing.T) {
	// Invalid ap.
	w(t, "leiasqz_v0_unknown.png")
}

func TestOCR_1200x_hbrx_v0_unknown(t *testing.T) {
	w(t, "hbrx_v0_unknown.png")
}

func TestOCR_1200x_forferdet_v0_unknowntablet(t *testing.T) {
	w(t, "forferdet_v0_unknowntablet.png")
}

func TestOCR_1200x_viddy_v1640_nexus7_2013(t *testing.T) {
	w(t, "viddy_v1640_nexus7_2013.png")
}

func TestOCR_1200x_forferdet_v0_unknown_1(t *testing.T) {
	w(t, "forferdet_v0_unknown-1.png")
}

func TestOCR_1440x_haste_v1680_nexus6(t *testing.T) {
	w(t, "haste_v1680_nexus6.png")
}

func TestOCR_1440x_haste_v1680_nexus6_1(t *testing.T) {
	w(t, "haste_v1680_nexus6-1.png")
}

func TestOCR_1440x_haste_v1680_nexus6_2(t *testing.T) {
	w(t, "haste_v1680_nexus6-1.png")
}

func TestOCR_1440x_haste_v1690_nexus6(t *testing.T) {
	w(t, "haste_v1690_nexus6.png")
}

func TestOCR_1440x_haste_v1690_nexus6_1(t *testing.T) {
	w(t, "haste_v1690_nexus6-1.png")
}

func TestOCR_1440x_haste_v1711_nexus6_1(t *testing.T) {
	w(t, "haste_v1711_nexus6.png")
}

func TestOCR_1440x_haste_v0_nexus6(t *testing.T) {
	// Invalid AP
	w(t, "haste_v0_nexus6.png")
}

func TestOCR_1440x_haste_v0_nexus6_1(t *testing.T) {
	// Invalid resonators deployed, links created, longest link ever created,
	// xm recharged, resonators destroyed, portals neutralized, enemy links,
	// destroyed, max time link maintained, max time field held, largest field
	// x mu days, glyph hack points.
	w(t, "haste_v0_nexus6-1.png")
}

func TestOCR_1440x_haste_v0_nexus6_2(t *testing.T) {
	// Invalid resonators deployed
	w(t, "haste_v0_nexus6-2.png")
}

func TestOCR_1440x_haste_v0_nexus6_3(t *testing.T) {
	// Invalid glyph hack points.
	w(t, "haste_v0_nexus6-3.png")
}

func TestOCR_1440x_haste_v0_nexus6_4(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus6-4.png")
}

func TestOCR_1440x_haste_v0_nexus6_5(t *testing.T) {
	// Invalid AP.
	w(t, "haste_v0_nexus6-5.png")
}

func TestOCR_1440x_haste_v0_nexus6_6(t *testing.T) {
	// Invalid unique missions completed
	w(t, "haste_v0_nexus6-6.png")
}

func TestOCR_1440x_haste_v0_nexus6_7(t *testing.T) {
	// Invalid unique portals visited, longest link ever created. unique
	// missions completed.
	w(t, "haste_v0_nexus6-7.png")
}

func TestOCR_1440x_haste_v0_nexus6_8(t *testing.T) {
	// Invalid hacks.
	w(t, "haste_v0_nexus6-8.png")
}

func TestOCR_1440x_haste_v0_nexus6_9(t *testing.T) {
	// Invalid hacks.
	w(t, "haste_v0_nexus6-9.png")
}

func TestOCR_1440x_haste_v0_nexus6_10(t *testing.T) {
	// Invalid resonators destroyed.
	w(t, "haste_v0_nexus6-10.png")
}

func TestOCR_1440x_haste_v0_nexus6_11(t *testing.T) {
	// Invalid longest link created.
	w(t, "haste_v0_nexus6-11.png")
}

func TestOCR_1440x_haste_v0_nexus6_12(t *testing.T) {
	// Invalid mind units captured.
	w(t, "haste_v0_nexus6-12.png")
}

func TestOCR_1440x_nittingduck_v0_unknown(t *testing.T) {
	// Invalid unique portals visited, control fields created, mint units
	// captured, largest control field, xm recharged, portals captured,
	// unique portals captured, hacks, largest control field.
	w(t, "nittingduck_v0_unknown.png")
}

func TestOCR_1440x_nittingduck_v0_unknown_1(t *testing.T) {
	// Invalid links created
	w(t, "nittingduck_v0_unknown-1.png")
}

func TestOCR_1536x_Scissorhill_v0_ipad(t *testing.T) {
	w(t, "scissorhill_v0_ipad.jpeg")
}
