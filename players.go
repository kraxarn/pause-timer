package main

import (
	"github.com/godbus/dbus/v5"
	"strings"
)

type MediaPlayer struct {
	name     string
	identity string
}

func newMediaPlayer(name string, conn *dbus.Conn) MediaPlayer {
	player := MediaPlayer{
		name: name,
	}

	obj := conn.Object(name, "/org/mpris/MediaPlayer2")
	prop, err := obj.GetProperty("org.mpris.MediaPlayer2.Identity")
	if err != nil {
		return player
	}

	var identity string
	if err = prop.Store(&identity); err != nil {
		return player
	}

	player.identity = identity
	return player
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
