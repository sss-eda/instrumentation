package domain

// Site TODO
type Site interface {
	Rename(SiteName) error
	ChangeAbbreviation(SiteAbbreviation) error
	Remove() error
}

// SiteID TODO
type SiteID interface {
	Equals(SiteID) bool
	String() string
}

// SiteAbbreviation TODO
type SiteAbbreviation string

// SiteName TODO
type SiteName string
