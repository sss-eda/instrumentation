package domain

// SiteID TODO
type SiteID interface {
	Equals(SiteID) bool
	String() string
}

// Site - This is an entity
type site struct {
	id           SiteID
	name         string
	abbreviation string
}

// ID TODO
func (entity site) ID() SiteID {
	return entity.id
}
