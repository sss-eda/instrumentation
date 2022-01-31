package instrumentType

// Aggregate TODO
type Aggregate interface {
	Rename(Name) error
	ChangeAbbreviation(Abbreviation) error
	Add(Name, Abbreviation) error
}
