package player

type PlaybackStatus uint8

const (
	Playing PlaybackStatus = iota
	Paused  PlaybackStatus = iota
	Stopped PlaybackStatus = iota
)
