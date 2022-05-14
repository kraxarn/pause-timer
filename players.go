package main

import (
	"github.com/godbus/dbus/v5"
	"strings"
)

type MediaPlayer struct {
	obj  dbus.BusObject
	name string
}

func newMediaPlayer(name string, conn *dbus.Conn) MediaPlayer {
	return MediaPlayer{
		name: name,
		obj:  conn.Object(name, "/org/mpris/MediaPlayer2"),
	}
}

func getAllMediaPlayers(conn *dbus.Conn) ([]MediaPlayer, error) {
	var names []string
	obj := conn.Object("org.freedesktop.DBus", "/")
	if err := obj.Call("ListNames", 0).Store(&names); err != nil {
		return nil, err
	}

	var players []MediaPlayer
	for _, name := range names {
		if strings.HasPrefix(name, "org.mpris.MediaPlayer2.") {
			players = append(players, newMediaPlayer(name, conn))
		}
	}

	return players, nil
}

func getProperty[V string](player *MediaPlayer, name string, value *V) error {
	prop, err := player.obj.GetProperty(name)
	if err != nil {
		return err
	}
	return prop.Store(value)
}

func (m *MediaPlayer) identity() string {
	var identity string
	_ = getProperty(m, "org.mpris.MediaPlayer2.Identity", &identity)
	return identity
}
