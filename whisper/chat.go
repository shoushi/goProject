package whisper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Chat(app fyne.App, name string) {
	chatWin := app.NewWindow(name)
	record := widget.NewRichText()
	record.Resize(fyne.NewSize(500, 300))

	inputArea := widget.NewMultiLineEntry()
	// 输入框提交
	submitBut := widget.NewButton("submit", func() {
		sentence := inputArea.Text
		if sentence == "" {
			return
		}
		inputArea.Text = ""
		inputArea.Refresh()
		// 修改富文本内容
		record.Segments = append(record.Segments, &widget.TextSegment{Text: CurrentName + " " + time.Now().Format("2006-01-02 15:04:05")},
			&widget.TextSegment{Text: sentence})
		record.Refresh()
		submitServer(123, sentence)
	})
	submitBut.Resize(fyne.NewSize(200, 30))
	// 清空输入框
	clearBut := widget.NewButton("clear", func() {
		inputArea.Text = ""
		inputArea.Refresh()
	})
	clearBut.Resize(fyne.NewSize(200, 30))
	inputContent := container.NewVBox(inputArea, container.NewHBox(layout.NewSpacer(), clearBut, submitBut))
	chaterWidget := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), record, inputContent)
	chatWin.SetContent(chaterWidget)
	chatWin.Resize(fyne.NewSize(500, 400))
	chatWin.CenterOnScreen()
	chatWin.SetIcon(theme.MailSendIcon())
	chatWin.Show()
}

// submit message
func submitServer(id float64, message string) {
	msg := MsgEntity{id, message}
	jsonStr, err := json.Marshal(msg)
	if err != nil {
		resp, err := http.Post("http://localhost:8080/mock/status", "application/json", strings.NewReader(string(jsonStr)))
		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println(body)
			defer resp.Body.Close()
		}
	}

}
