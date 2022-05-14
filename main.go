package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"os"
)

func main() {
	flags := parseFlags()

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
			fmt.Println(player.identity())
		}
		os.Exit(0)
	}

	if len(players) > 1 && len(flags.player) == 0 {
		fmt.Printf("Multiple players found, specify one of:")
		printPlayerIdentities(players)
		fmt.Println()
		os.Exit(1)
	}

	var current *MediaPlayer
	for _, player := range players {
		if player.identity() == flags.player {
			current = &player
			break
		}
	}

	if current == nil {
		fmt.Printf("Player \"%s\" not found, specify one of:", flags.player)
		printPlayerIdentities(players)
		fmt.Println()
		os.Exit(1)
	}

	fmt.Println(":)")
}

func printPlayerIdentities(players []MediaPlayer) {
	for i, player := range players {
		if i == 0 {
			fmt.Printf(" \"%s\"", player.identity())
		} else {
			fmt.Printf(", \"%s\"", player.identity())
		}
	}
}
