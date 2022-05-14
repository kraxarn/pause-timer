package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"os"
)

func main() {
	flags := parseFlags()
	if flags == nil {
		os.Exit(1)
		return
	}

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}

	var players []MediaPlayer
	players, err = getAllMediaPlayers(conn)
	if err != nil {
		panic(err)
	}

	if flags.list {
		for _, player := range players {
			fmt.Println(player.identity)
		}
		os.Exit(0)
	}

	for _, player := range players {
		fmt.Println(player.name, player.identity)
	}
}
