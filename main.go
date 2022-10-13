package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

var NickName string
var TextCipher *Cipher
var MessageMap map[[32]byte]string // key: QRCode digest ; value: decrypted message json

func main() {
	TextCipher = NewCipher("defaultpasswordpleasechange")
	MessageMap = make(map[[32]byte]string, 0)

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "SafeChat",
		Width:     400,
		Height:    800,
		Assets:    assets,
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
