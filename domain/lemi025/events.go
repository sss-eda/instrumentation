package lemi025

type EventType int

const (
	ConfigRead EventType = iota
	TimeRead
	TimeSet
)
