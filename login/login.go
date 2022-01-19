package login

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Login(myWindow fyne.Window) {
	email := widget.NewEntry()
	password := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "email", Widget: email}, {Text: "password", Widget: password}},
		OnSubmit: func() {
			// 登录访问后台
			postServer(email.Text, password.Text)
			myWindow.Close()
		}, OnCancel: func() {
			myWindow.Close()
		},
	}
	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.SetContent(form)
}

func postServer(email string, password string) {
	// password进行md5加密
	data := []byte(password)
	pass := fmt.Sprintf("%x", md5.Sum(data))
	log.Println(pass)
	req := LoginEntity{
		Email:    email,
		Password: pass,
	}
	log.Println(req.Email, req.Password)
	// 生成json报文
	jsonStr, err := json.Marshal(req)
	if err != nil {
		log.Println("生成json错误")
	}
	log.Println(string(jsonStr))
	http.Post("http://localhost:8080/mock/status", "application/json", strings.NewReader(string(jsonStr)))
}
