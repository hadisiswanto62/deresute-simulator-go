package enum

// Attribute represents idol's (or song's) primary attribute
type Attribute string

// All attributes
const (
	Cute    Attribute = "Cute"
	Cool              = "Cool"
	Passion           = "Passion"
	All               = "All"
)

// ForIdol returns all valid attributes for idols
func ForIdol() [3]Attribute {
	return [3]Attribute{Cute, Cool, Passion}
}

// ForSong returns all valid attributes for songs
func ForSong() [4]Attribute {
	return [4]Attribute{Cute, Cool, Passion, All}
}
