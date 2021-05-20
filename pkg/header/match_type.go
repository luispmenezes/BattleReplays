package header

type MatchType struct {
	typeId uint8
}

func (m MatchType) AsString() string {
	if m.typeId > 15 {
		return "ERROR"
	}

	idStrings := []string{"UNKNOWN",
		"QUICK2V2",
		"DEBUG",
		"PRIVATE",
		"TUTORIAL",
		"TRAINING",
		"QUICK3V3",
		"VSAI",
		"BRAWL",
		"CAMPAIGN",
		"BATTLEGROUNDS",
		"ROYALRUMBLESOLO",
		"ROYALRUMBLEDUO",
		"ROYALETUTORIAL",
		"ROYALEVSAI",
		"ROYALEPRIVATE"}

	return idStrings[m.typeId]
}
