package gohotkey

type Handler interface {
	Do(k Hotkey)
}

// See https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msg
type tagMSG struct {
	HWND   uintptr
	UINT   uintptr
	WPARAM uint //→ https://stackoverflow.com/a/2515285
	LPARAM int64
	DWORD  int32
	POINT  struct{ X, Y int64 }
}
