# GoHotKey

A simple package to capture the keyboard in **Windows**.
_(tested on **Windows 10**, 64-bit)_



## Download

```
go get github.com/Adrosar/gohotkey
```



## How to use

```go
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
```
↑ [basic](example/basic/main.go) example



## How to build

Go to the directory where the `main.go` file is located and execute the command:

```
go build -o app.exe main.go
```



## Early stage
The software was created for personal use and is in the early stages of development.



## License

I put the software temporarily under the Go-compatible **BSD** license. If this prevents someone from using the software, do let me know and I'll consider changing it.



## Author

Adrian Gargula | [github.com/Adrosar](https://github.com/Adrosar) | [bitbucket.org/Adrosar](https://bitbucket.org/Adrosar)