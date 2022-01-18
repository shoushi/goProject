package main

import (
	"image/color"
	"myQQ/login"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func updateClock(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
func showAnother(a fyne.App) {
	win := a.NewWindow("Layout")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Here", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(400, 400))
	content := container.New(layout.NewGridLayout(2), text1, text2)
	win.SetContent(content)
	win.Show()
}
func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	clock := widget.NewLabel("")
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateClock(clock)
		}
	}()
	// go showAnother(a)
	go login.Login(a)
	w.ShowAndRun()
}
