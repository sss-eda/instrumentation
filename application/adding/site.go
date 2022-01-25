package adding

// SiteID TODO
type SiteID interface {
	Equals(SiteID) bool
}

// Site TODO
type Site struct {
	Name         string
	Abbreviation string
}
