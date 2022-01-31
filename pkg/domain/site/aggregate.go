package site

// Aggregate TODO
type Aggregate interface {
	Add(Name, Abbreviation) error
	Rename(Name) error
	ChangeAbbreviation(Abbreviation) error
	Remove() error
}
