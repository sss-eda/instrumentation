package lemi025

// Connection TODO
type Connection interface {
	ReadConfig() error
	ReadTime() error
	SetTime() error
}
