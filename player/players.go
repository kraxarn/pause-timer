package player

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"strings"
)

func GetAll(conn *dbus.Conn) ([]MediaPlayer, error) {
	var names []string
	obj := conn.Object("org.freedesktop.DBus", "/")
	if err := obj.Call("ListNames", 0).Store(&names); err != nil {
		return nil, err
	}

	var players []MediaPlayer
	for _, name := range names {
		if strings.HasPrefix(name, "org.mpris.MediaPlayer2.") {
			players = append(players, NewMediaPlayer(name, conn))
		}
	}

	return players, nil
}

func PrintAll(players []MediaPlayer) {
	for i, player := range players {
		if i == 0 {
			fmt.Printf(" \"%s\"", player.Identity())
		} else {
			fmt.Printf(", \"%s\"", player.Identity())
		}
	}
}
