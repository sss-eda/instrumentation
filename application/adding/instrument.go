package adding

import (
	"fmt"
	"strconv"
	"strings"
)

// InstrumentID TODO
type InstrumentID interface {
	Equals(InstrumentID) bool
}

// Instrument TODO - Aggregate Root, Site and InstrumentType are entities.
type Instrument struct {
	Name      string
	Site      Site
	Type      InstrumentType
	Iteration uint8
}

// Abbreviation TODO
func (instrument *Instrument) Abbreviation() string {
	abbreviation := strings.Join([]string{
		instrument.Site.Abbreviation,
		instrument.Type.Abbreviation,
		strconv.Itoa(int(instrument.Iteration)),
	}, "")

	var a int64 = 10

	b := uint8(a)
	fmt.Println(b)

	return abbreviation
}
