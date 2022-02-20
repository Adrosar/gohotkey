package main

import (
	"fmt"
	"log"

	"github.com/Adrosar/gohotkey"
)

type handler struct{}

func (h *handler) Do(k gohotkey.Hotkey) {
	fmt.Println(k.String())
}

func main() {
	k := gohotkey.NewKeyboard()
	k.SetHandler(&handler{})

	k.Add(true, true, false, false, 'K')    // ALT + CTRL + K
	k.Add(false, false, false, true, 'N')   // Win + N
	k.Add(false, false, false, false, 0xB3) // Additional "play/pause" button

	err := k.Listen()
	if err != nil {
		log.Fatalln(err)
	}
}
