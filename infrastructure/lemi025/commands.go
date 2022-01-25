package lemi025

type commandType int

const (
	// ReadConfig TODO
	ReadConfig commandType = iota
	// ReadTime TODO
	ReadTime
	// SetTime TODO
	SetTime
)

// String TODO
func (ct commandType) String() string {
	return ""
}
