package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/kraxarn/pause-timer/player"
	"os"
	"time"
)

func main() {
	flags := parseFlags()

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}

	var mediaPlayers []player.MediaPlayer
	mediaPlayers, err = player.GetAll(conn)
	if err != nil {
		panic(err)
	}

	if flags.list {
		for _, mediaPlayer := range mediaPlayers {
			fmt.Println(mediaPlayer.Identity())
		}
		os.Exit(0)
	}

	if len(mediaPlayers) > 1 && len(flags.player) == 0 {
		fmt.Printf("Multiple players found, specify one of:")
		player.PrintAll(mediaPlayers)
		fmt.Println()
		os.Exit(1)
	}

	var current *player.MediaPlayer
	for _, mediaPlayer := range mediaPlayers {
		if mediaPlayer.Identity() == flags.player {
			current = &mediaPlayer
			break
		}
	}

	if current == nil {
		fmt.Printf("Player \"%s\" not found, specify one of:", flags.player)
		player.PrintAll(mediaPlayers)
		fmt.Println()
		os.Exit(1)
	}

	for {
		// Wait for playback to start
		fmt.Println("Waiting for playback to start...")
		for {
			status := current.PlaybackStatus()
			if status == player.Playing {
				break
			}
			time.Sleep(time.Second)
		}

		// Started, print message first
		fmt.Printf("Playback started, waiting for ")
		if flags.minutes == 1 {
			fmt.Printf("%d minute", flags.minutes)
		} else {
			fmt.Printf("%d minutes", flags.minutes)
		}

		// Wait for flags.minutes minutes
		i := 0
		for {
			time.Sleep(time.Minute)
			fmt.Print(".")

			i++
			if i >= flags.minutes {
				break
			}
		}

		// Time's up, pause
		current.Pause()
	}
}
