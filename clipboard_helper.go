package main

import "golang.design/x/clipboard"

func ToClipboard() {
	clipboard.Write(clipboard.FmtText, []byte("text data"))
}
func PNGToClipboard(pngByte []byte) {
	clipboard.Write(clipboard.FmtImage, pngByte)
}
