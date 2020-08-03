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

// AllIdolAttributes is all valid attributes for idols
var AllIdolAttributes = [3]Attribute{AttrCute, AttrCool, AttrPassion}

// AllSongAttributes is all valid attributes for songs
var AllSongAttributes = [4]Attribute{AttrCute, AttrCool, AttrPassion, AttrAll}
