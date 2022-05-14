package main

import "flag"

type Flags struct {
	current string
	minutes int
	list    bool
}

func parseFlags() *Flags {
	flags := &Flags{}

	flag.StringVar(&flags.current, "player", "", "Media player to use")
	flag.IntVar(&flags.minutes, "minutes", 5, "Minutes to wait")
	flag.BoolVar(&flags.list, "list", false, "List all available media players")
	flag.Parse()

	return flags
}
