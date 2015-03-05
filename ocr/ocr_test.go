package ocr

import (
	"flag"
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
}

var (
	onlyTop         bool
	onlyBottom      bool
	overrideProfile int
)

func init() {
	conf.Load("../config.json")
	conf.Config.Cache = "testdata/"
	InitConfig()

	flag.BoolVar(&onlyTop, "top", false, "Only validate top part of OCR image.")
	flag.BoolVar(&onlyBottom, "bottom", false, "Only validate bottom part of OCR image.")
	flag.IntVar(&overrideProfile, "profile", 0, "Profile to use.")

	flag.Parse()
}

func validateTop(t *testing.T, file string, p profile.Profile) {
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
}

func validateBottom(t *testing.T, file string, p profile.Profile) {
	e := testData[file]

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
	t.Parallel()

	p := New(file, overrideProfile)
	p.Split()

	if !onlyBottom {
		p.ProcessTop()
		validateTop(t, file, p.Profile)
	}

	if !onlyTop {
		p.ProcessBottom()
		p.ProcessInnovator()
		validateBottom(t, file, p.Profile)
	}

	p.CleanUp()
}

func TestOCR_480x_wexp_v0_unknown(t *testing.T) {
	w(t, "wexp_v0_unknown.png")
}

func TestOCR_480x_wexp_v0_unknown_1(t *testing.T) {
	// Invalid Max Link Length x Days
	w(t, "wexp_v0_unknown-1.png")
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

func TestOCR_720x_sockerdricka_v0_unknown(t *testing.T) {
	w(t, "sockerdricka_v0_unknown.png")
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

func TestOCR_1080x_steino_v0_unknown(t *testing.T) {
	// Invalid nick.
	w(t, "steino_v0_unknown.png")
}

func TestOCR_1080x_steino_v0_unknown_1(t *testing.T) {
	// Invalid AP
	w(t, "steino_v0_unknown-1.png")
}

func TestOCR_1080x_haste_v1620_nexus5(t *testing.T) {
	w(t, "haste_v1620_nexus5.png")
}

func TestOCR_1080x_forferdet_v0_unknown(t *testing.T) {
	w(t, "forferdet_v0_unknown.png")
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

func TestOCR_1200x_forferdet_v0_unknowntablet(t *testing.T) {
	w(t, "forferdet_v0_unknowntablet.png")
}

func TestOCR_1200x_viddy_v1640_nexus7_2013(t *testing.T) {
	w(t, "viddy_v1640_nexus7_2013.png")
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

func TestOCR_1536x_Scissorhill_v0_ipad(t *testing.T) {
	w(t, "scissorhill_v0_ipad.jpeg")
}
