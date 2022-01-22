package whisper

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Chat(app fyne.App, name string) {
	chatWin := app.NewWindow(name)
	record := widget.NewRichText()

	inputArea := widget.NewMultiLineEntry()
	inputAreaBut := widget.NewButton("submit", func() {
		sentence := inputArea.Text
		log.Println(sentence)
		inputArea.Text = ""
		inputArea.Refresh()
		record.ParseMarkdown(sentence)
	})
	// tabs := container.NewAppTabs(container.NewTabItem("Tab 1", widget.NewLabel("Hello!")), container.NewTabItem("Tab 1", widget.NewLabel("World!")))
	chaterWidget := container.NewGridWithRows(3, record, inputArea, inputAreaBut)
	chatWin.SetContent(chaterWidget)
	chatWin.Resize(fyne.NewSize(300, 200))
	chatWin.Show()
}
