package lemi025

import "github.com/sss-eda/instrumentation/domain/lemi025"

// Service TODO
type Service struct {
	Mutations Mutations
}

type Mutations struct {
	ReadConfig lemi025.ReadConfig
	ReadTime   lemi025.ReadTime
}
