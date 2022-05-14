package player

import (
	"fmt"
	"github.com/godbus/dbus/v5"
)

type MediaPlayer struct {
	obj  dbus.BusObject
	name string
}

func NewMediaPlayer(name string, conn *dbus.Conn) MediaPlayer {
	return MediaPlayer{
		name: name,
		obj:  conn.Object(name, "/org/mpris/MediaPlayer2"),
	}
}

func getProperty[V string](player *MediaPlayer, name string, value *V) error {
	prop, err := player.obj.GetProperty(name)
	if err != nil {
		return err
	}
	return prop.Store(value)
}

func (m *MediaPlayer) Identity() string {
	var identity string
	_ = getProperty(m, "org.mpris.MediaPlayer2.Identity", &identity)
	return identity
}

func (m *MediaPlayer) PlaybackStatus() PlaybackStatus {
	var playbackStatus string
	err := getProperty(m, "org.mpris.MediaPlayer2.Player.PlaybackStatus", &playbackStatus)
	if err != nil {
		fmt.Println(err)
	}

	switch playbackStatus {
	case "Playing":
		return Playing

	case "Paused":
		return Paused

	default:
		return Stopped
	}
}

func (m *MediaPlayer) Pause() {
	call := m.obj.Call("org.mpris.MediaPlayer2.Player.Pause", 0)
	if call.Err != nil {
		println(call.Err)
	}
}
