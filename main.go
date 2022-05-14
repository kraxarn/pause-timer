package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"strings"
)

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}

	var names []string
	obj := conn.Object("org.freedesktop.DBus", "/")
	if err = obj.Call("ListNames", 0).Store(&names); err != nil {
		panic("empty d-conn session")
	}

	for _, name := range names {
		if !strings.HasPrefix(name, "org.mpris.MediaPlayer2.") {
			continue
		}

		if id, err := getIdentifier(conn, name); err == nil {
			fmt.Printf("%s (%s)\n", id, name[23:])
		} else {
			fmt.Printf("%s", name)
		}
	}
}

func getIdentifier(conn *dbus.Conn, dest string) (string, error) {
	obj := conn.Object(dest, "/org/mpris/MediaPlayer2")

	prop, err := obj.GetProperty("org.mpris.MediaPlayer2.Identity")
	if err != nil {
		return "", err
	}

	var identity string
	if err = prop.Store(&identity); err != nil {
		return "", err
	}

	return identity, nil
}
