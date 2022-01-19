package main

import (
	"myQQ/login"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	myWindow := a.NewWindow("Login")
	login.Login(myWindow)
	myWindow.ShowAndRun()
}
