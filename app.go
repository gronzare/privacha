package main

import (
	"context"
	"crypto/sha256"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

type Message struct {
	Sender  string
	Time    string
	Content string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetMsgList() []Message {
	ReadQR()
	png := EncodeQR("大家好我是丁真999")
	PNGToClipboard(png)
	return []Message{
		{"aaa", "123", "hhh"},
		{"bbb", "124", "hhhhh你好"},
	}
}

// return JSONstring slice
func (a *App) ScanNewMsg() []string {
	var newMsg []string
	QREncodedStrings := ReadQR()
	for _, encryptedStr := range QREncodedStrings {
		digest := sha256.Sum256([]byte(encryptedStr))
		_, exist := MessageMap[digest]
		if !exist {
			jsonStr := string(TextCipher.Decrypt(encryptedStr))
			fmt.Println("Decrypt QR result:", jsonStr)
			newMsg = append(newMsg, jsonStr)
			MessageMap[digest] = jsonStr
		}
	}
	return newMsg
}

func (a *App) SetPassword(pwd string) {
	TextCipher = NewCipher(pwd)
}

func (a *App) SetNickName(nickName string) {
	NickName = nickName
}

// encrypt json text, and convert to PNG, write into clipboard
func (a *App) EncryptText(textJson string) {
	encryptResult := TextCipher.Encrypt([]byte(textJson))
	pngByte := EncodeQR(encryptResult)
	PNGToClipboard(pngByte)
}
