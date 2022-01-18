package login

import (
	"crypto/md5"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Login(myApp fyne.App) {
	myWindow := myApp.NewWindow("Form Widget")

	username := widget.NewEntry()
	password := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "username", Widget: username}, {Text: "password", Widget: password}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", username.Text)
			log.Println("password:", password.Text)
			postServer(username.Text, password.Text)
			myWindow.Close()
		}, OnCancel: func() {
			myWindow.Close()
		},
	}

	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.SetContent(form)
	myWindow.Show()
}

func postServer(username string, password string) {
	data := []byte(password)
	pass := fmt.Sprintf("%x", md5.Sum(data))
	log.Println(pass)
}
