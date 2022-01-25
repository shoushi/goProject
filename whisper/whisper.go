package whisper

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var CurrentName string

func Whisper(app fyne.App) {
	whisperWin := app.NewWindow("whisper")

	// todo 后面访问server获取好友列表
	friendData := []string{"shoushi", "caoyang", "little white", "ysy"}
	friendDataList := binding.NewStringList()
	friendList := widget.NewListWithData(friendDataList, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(di binding.DataItem, co fyne.CanvasObject) {
		i := di.(binding.String)
		txt, _ := i.Get()
		label1 := co.(*widget.Label)
		label1.SetText(txt)
	})
	for _, v := range friendData {
		friendDataList.Append(v)
	}

	groupData := []string{"Just do it", "xxx", "xxx", "xxx"}
	groupDataList := binding.NewStringList()
	groupList := widget.NewListWithData(groupDataList, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(di binding.DataItem, co fyne.CanvasObject) {
		i := di.(binding.String)
		txt, _ := i.Get()
		label1 := co.(*widget.Label)
		label1.SetText(txt)
	})
	for _, v := range groupData {
		groupDataList.Append(v)
	}
	friendList.OnSelected = func(id widget.ListItemID) {
		log.Println("和", friendData[id], "聊天")
		go Chat(app, friendData[id])
	}
	friendTab := container.NewGridWithRows(2, friendList)
	groupTab := container.NewGridWithRows(2, groupList)
	tabs := container.NewAppTabs(
		container.NewTabItem("好友", friendTab),
		container.NewTabItem("群组", groupTab),
	)

	whisperWin.Resize(fyne.NewSize(300, 800))
	whisperWin.SetContent(tabs)
	whisperWin.SetIcon(theme.HomeIcon())
	whisperWin.Show()
}
