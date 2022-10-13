package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"

	"github.com/getlantern/systray"
	"github.com/nfnt/resize"
)

func main() {
	onExit := func() {

	}

	systray.Run(onReady, onExit)
}

func onReady() {
	go func() {
		for range time.Tick(time.Second) {
			img := image.NewRGBA(image.Rect(0, 0, 96, 14))
			addLabel(img, 0, 12, time.Now().Format("15 : 04 : 05"))
			img_scaled := resize.Resize(960, 140, img, resize.NearestNeighbor)
			// img_scaled := img

			icon := img2bytes(img_scaled)
			systray.SetIcon(icon)
		}
	}()

	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{0, 0, 0, 255}
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: inconsolata.Regular8x16,
		Dot:  point,
	}
	d.DrawString(label)
}

func img2bytes(img image.Image) []byte {

	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)

	if err != nil {
		panic(err.Error())
	}
	return buf.Bytes()
}
