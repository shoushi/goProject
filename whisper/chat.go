package whisper

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Chat(app fyne.App, name string) {
	chatWin := app.NewWindow(name)
	lab := "lao ke with name " + name
	chatlab := widget.NewLabel(lab)
	chatWin.SetContent(chatlab)
	chatWin.Resize(fyne.NewSize(300, 200))
	chatWin.Show()
}
