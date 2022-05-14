package main

import "flag"

type Flags struct {
	player  string
	minutes int
	list    bool
}

func parseFlags() Flags {
	var flags Flags

	flag.StringVar(&flags.player, "player", "", "Media player to use")
	flag.IntVar(&flags.minutes, "minutes", 5, "Minutes to wait")
	flag.BoolVar(&flags.list, "list", false, "List all available media players")
	flag.Parse()

	return flags
}
