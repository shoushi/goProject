package whisper

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Whisper(app fyne.App) {
	whisperWin := app.NewWindow("whisper")

	// todo 后面访问server获取好友列表
	data := []string{"shoushi", "caoyang", "little white", "ysy"}
	datalist := binding.NewStringList()
	userList := widget.NewListWithData(datalist, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(di binding.DataItem, co fyne.CanvasObject) {
		i := di.(binding.String)
		txt, _ := i.Get()
		label1 := co.(*widget.Label)
		label1.SetText(txt)
	})
	chaterWidget := container.NewGridWithRows(2, userList)
	for _, v := range data {
		datalist.Append(v)
	}
	userList.OnSelected = func(id widget.ListItemID) {
		log.Println("和", data[id], "聊天")
		go Chat(app, data[id])
	}
	whisperWin.Resize(fyne.NewSize(300, 1000))
	whisperWin.SetContent(chaterWidget)
	whisperWin.SetIcon(theme.HomeIcon())
	whisperWin.Show()
}
