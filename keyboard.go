package gohotkey

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

type Keyboard struct {
	currentId uint
	delay     time.Duration
	list      map[uint]Hotkey
	handler   Handler
}

func NewKeyboard() *Keyboard {
	return &Keyboard{
		currentId: 0,
		delay:     time.Millisecond * 100,
		list:      make(map[uint]Hotkey),
	}
}

func (kb *Keyboard) SetHandler(h Handler) {
	kb.handler = h
}

func (kb *Keyboard) SetDelay(t time.Duration) {
	kb.delay = t
}

// Adds a key combination.
//
// Argument "char" see https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
func (kb *Keyboard) Add(isAlt bool, isCtrl bool, isShift bool, isWin bool, char rune) uint {
	kb.currentId++
	id := kb.currentId
	kb.list[id] = Hotkey{id, isAlt, isCtrl, isShift, isWin, char}
	return id
}

func (kb *Keyboard) Listen() error {
	var err error

	// §1
	user32 := syscall.NewLazyDLL("user32")
	err = user32.Load()
	if err != nil {
		return fmt.Errorf(`[AHupvk] loading error "user32", %v`, err.Error())
	}

	// §2
	registerHotKey := user32.NewProc("RegisterHotKey")
	err = registerHotKey.Find()
	if err != nil {
		return fmt.Errorf(`[VbY8wv] procedure "RegisterHotKey" not found, %v`, err.Error())
	}

	// §3
	for _, hk := range kb.list {
		var modifiers uint = 0x0000

		if hk.IsAlt {
			modifiers += 0x0001
		}

		if hk.IsCtrl {
			modifiers += 0x0002
		}

		if hk.IsShift {
			modifiers += 0x0004
		}

		if hk.IsWin {
			modifiers += 0x0008
		}

		// ↓ See https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerhotkey
		r1, r2, lastError := registerHotKey.Call(0, uintptr(hk.Id), uintptr(modifiers), uintptr(hk.Rune))
		if r1 != 1 {
			return fmt.Errorf(`[YhGnp9] registration error %s (%v, %v, %v)`, hk.String(), r1, r2, lastError.Error())
		}
	}

	// §4
	getMessageW := user32.NewProc("GetMessageW")
	err = getMessageW.Find()
	if err != nil {
		return fmt.Errorf(`[E4c7e3] procedure "GetMessageW" not found, %v`, err.Error())
	}

	// §5
	for {
		var msg = &tagMSG{}
		// ↓ See https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagew
		getMessageW.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0)

		if id := msg.WPARAM; id != 0 {
			hk, ok := kb.list[id]
			if ok {
				kb.handler.Do(hk)
			}
		}

		if kb.delay > 0 {
			time.Sleep(kb.delay)
		}
	}
}
