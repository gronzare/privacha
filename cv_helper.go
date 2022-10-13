package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"

	"github.com/kbinani/screenshot"
	"github.com/makiuchi-d/gozxing"
	multiQR "github.com/makiuchi-d/gozxing/multi/qrcode"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func GetScreenImage() ([]*image.RGBA, error) {
	n := screenshot.NumActiveDisplays()
	images := make([]*image.RGBA, 0)
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}
	return images, nil
}

func ReadQR() []string {
	images, err := GetScreenImage()
	fmt.Println("Reading QR")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	QRTexts := make([]string, 0)
	for _, img := range images {
		// prepare BinaryBitmap
		bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

		// decode image
		qrReader := multiQR.NewQRCodeMultiReader()
		results, _ := qrReader.DecodeMultipleWithoutHint(bmp)
		//fmt.Println(result)
		for _, result := range results {
			QRTexts = append(QRTexts, result.GetText())
		}
	}
	fmt.Println(QRTexts)
	return QRTexts
}

func EncodeQR(text string) []byte {
	enc := qrcode.NewQRCodeWriter()
	img, _ := enc.Encode(text, gozxing.BarcodeFormat_QR_CODE, 600, 600, nil)

	buffer := bytes.NewBuffer([]byte{})

	// *BitMatrix implements the image.Image interface,
	// so it is able to be passed to png.Encode directly.
	_ = png.Encode(buffer, img)

	return buffer.Bytes()
}
