package main

import (
	"myQQ/login"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	login.Login(a)
	a.Run()
}
