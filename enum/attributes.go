package enum

// Attribute represents idol's (or song's) primary attribute
type Attribute string

// All attributes
const (
	AttrCute    Attribute = "cute"
	AttrCool    Attribute = "cool"
	AttrPassion Attribute = "passion"
	AttrAll     Attribute = "all"
)

// AttrForIdol is all valid attributes for idols
var AttrForIdol = [3]Attribute{AttrCute, AttrCool, AttrPassion}

// AttrForSong is all valid attributes for songs
var AttrForSong = [4]Attribute{AttrCute, AttrCool, AttrPassion, AttrAll}

func GetAttribute(i int) Attribute {
	switch i {
	case 1:
		return AttrCute
	case 2:
		return AttrCool
	case 3:
		return AttrPassion
	case 4:
		return AttrAll
	}
	return AttrAll
}
