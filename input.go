package main

import (
	"fmt"
	"log"

	"golang.design/x/hotkey"
)

func listenHotkey(key hotkey.Key, mods ...hotkey.Modifier) (err error) {
	ms := []hotkey.Modifier{}
	ms = append(ms, mods...)
	hk := hotkey.New(ms, key)

	err = hk.Register()
	if err != nil {
		return
	}

	choice := mapKeyCodeToInteger(hk.String())
	log.Default().Printf(hk.String())
	log.Default().Printf(fmt.Sprint(choice))

	// Blocks until the hokey is triggered.
	<-hk.Keydown()
	updateWindow(choice)
	<-hk.Keyup()
	log.Printf("hotkey: %v is up\n", hk)
	hk.Unregister()
	return
}
