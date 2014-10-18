package ocr

import (
	"io/ioutil"
	"os"
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
}

func init() {
	f, _ := ioutil.TempFile("", "testconfig")
	defer os.Remove(f.Name())
	ioutil.WriteFile(f.Name(), []byte(`{"cache": "testdata/"}`), 0644)

	conf.Load(f.Name())
}

func validateOCR(t *testing.T, file string, p, e profile.Profile) {
	if p.Nick != e.Nick {
		t.Errorf("%s: .Nick: %v != %v", file, p.Nick, e.Nick)
	}

	if p.Level != e.Level {
		t.Errorf("%s: .Level: %v != %v", file, p.Level, e.Level)
	}

	if p.AP != e.AP {
		t.Errorf("%s: .AP: %v != %v", file, p.AP, e.AP)
	}

	if p.UniquePortalsVisited != e.UniquePortalsVisited {
		t.Errorf("%s: .UniquePortalsVisited: %v != %v", file, p.UniquePortalsVisited, e.UniquePortalsVisited)
	}

	if p.PortalsDiscovered != e.PortalsDiscovered {
		t.Errorf("%s: .PortalsDiscovered: %v != %v", file, p.PortalsDiscovered, e.PortalsDiscovered)
	}

	if p.XMCollected != e.XMCollected {
		t.Errorf("%s: .XMCollected: %v != %v", file, p.XMCollected, e.XMCollected)
	}

	if p.Hacks != e.Hacks {
		t.Errorf("%s: .Hacks: %v != %v", file, p.Hacks, e.Hacks)
	}

	if p.ResonatorsDeployed != e.ResonatorsDeployed {
		t.Errorf("%s: .ResonatorsDeployed: %v != %v", file, p.ResonatorsDeployed, e.ResonatorsDeployed)
	}

	if p.LinksCreated != e.LinksCreated {
		t.Errorf("%s: .LinksCreated: %v != %v", file, p.LinksCreated, e.LinksCreated)
	}

	if p.ControlFieldsCreated != e.ControlFieldsCreated {
		t.Errorf("%s: .ControlFieldsCreated: %v != %v", file, p.ControlFieldsCreated, e.ControlFieldsCreated)
	}

	if p.MindUnitsCaptured != e.MindUnitsCaptured {
		t.Errorf("%s: .MindUnitsCaptured: %v != %v", file, p.MindUnitsCaptured, e.MindUnitsCaptured)
	}

	if p.LongestLinkEverCreated != e.LongestLinkEverCreated {
		t.Errorf("%s: .LongestLinkEverCreated: %v != %v", file, p.LongestLinkEverCreated, e.LongestLinkEverCreated)
	}

	if p.LargestControlField != e.LargestControlField {
		t.Errorf("%s: .LargestControlField: %v != %v", file, p.LargestControlField, e.LargestControlField)
	}

	if p.XMRecharged != e.XMRecharged {
		t.Errorf("%s: .XMRecharged: %v != %v", file, p.XMRecharged, e.XMRecharged)
	}

	if p.PortalsCaptured != e.PortalsCaptured {
		t.Errorf("%s: .PortalsCaptured: %v != %v", file, p.PortalsCaptured, e.PortalsCaptured)
	}

	if p.UniquePortalsCaptured != e.UniquePortalsCaptured {
		t.Errorf("%s: .UniquePortalsCaptured: %v != %v", file, p.UniquePortalsCaptured, e.UniquePortalsCaptured)
	}

	if p.ResonatorsDestroyed != e.ResonatorsDestroyed {
		t.Errorf("%s: .ResonatorsDestroyed: %v != %v", file, p.ResonatorsDestroyed, e.ResonatorsDestroyed)
	}

	if p.PortalsNeutralized != e.PortalsNeutralized {
		t.Errorf("%s: .PortalsNeutralized: %v != %v", file, p.PortalsNeutralized, e.PortalsNeutralized)
	}

	if p.EnemyLinksDestroyed != e.EnemyLinksDestroyed {
		t.Errorf("%s: .EnemyLinksDestroyed: %v != %v", file, p.EnemyLinksDestroyed, e.EnemyLinksDestroyed)
	}

	if p.EnemyControlFieldsDestroyed != e.EnemyControlFieldsDestroyed {
		t.Errorf("%s: .EnemyControlFieldsDestroyed: %v != %v", file, p.EnemyControlFieldsDestroyed, e.EnemyControlFieldsDestroyed)
	}

	if p.DistanceWalked != e.DistanceWalked {
		t.Errorf("%s: .DistanceWalked: %v != %v", file, p.DistanceWalked, e.DistanceWalked)
	}

	if p.MaxTimePortalHeld != e.MaxTimePortalHeld {
		t.Errorf("%s: .MaxTimePortalHeld: %v != %v", file, p.MaxTimePortalHeld, e.MaxTimePortalHeld)
	}

	if p.MaxTimeLinkMaintained != e.MaxTimeLinkMaintained {
		t.Errorf("%s: .MaxTimeLinkMaintained: %v != %v", file, p.MaxTimeLinkMaintained, e.MaxTimeLinkMaintained)
	}

	if p.MaxLinkLengthXDays != e.MaxLinkLengthXDays {
		t.Errorf("%s: .MaxLinkLengthXDays: %v != %v", file, p.MaxLinkLengthXDays, e.MaxLinkLengthXDays)
	}

	if p.MaxTimeFieldHeld != e.MaxTimeFieldHeld {
		t.Errorf("%s: .MaxTimeFieldHeld: %v != %v", file, p.MaxTimeFieldHeld, e.MaxTimeFieldHeld)
	}

	if p.LargestFieldMUsXDays != e.LargestFieldMUsXDays {
		t.Errorf("%s: .LargestFieldMUsXDays: %v != %v", file, p.LargestFieldMUsXDays, e.LargestFieldMUsXDays)
	}
}

func TestOCR(t *testing.T) {
	for file, expected := range testData {
		res := runOCR(file)
		p := buildProfile(res)

		validateOCR(t, file, p, expected)
	}
}
