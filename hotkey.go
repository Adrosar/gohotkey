package gohotkey

import "fmt"

type Hotkey struct {
	Id      uint
	IsAlt   bool
	IsCtrl  bool
	IsShift bool
	IsWin   bool
	Rune    rune
}

func (k *Hotkey) String() string {
	return fmt.Sprintf(`<ID:%d, Char:%s, Code:%d, Alt:%t, Ctrl:%t, Shift:%t, Win:%t>`, k.Id, string(k.Rune), k.Rune, k.IsAlt, k.IsCtrl, k.IsShift, k.IsWin)
}

func (k *Hotkey) Code() int32 {
	return int32(k.Rune)
}

func (k *Hotkey) Char() string {
	return string(k.Rune)
}
