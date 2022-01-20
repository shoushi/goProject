package login

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myQQ/whisper"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Login(app fyne.App) {

	myWindow := app.NewWindow("Login")
	email := widget.NewEntry()
	password := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "email", Widget: email}, {Text: "password", Widget: password}},
		OnSubmit: func() {
			// 登录访问后台
			res := postServer(email.Text, password.Text)
			log.Println("post 结果", res.Success)
			if res.Success {
				myWindow.Close()
				go whisper.Whisper(app)
			} else {
				myWindow.Close()
			}
		}, OnCancel: func() {
			myWindow.Close()
		},
	}
	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.SetContent(form)
	myWindow.SetIcon(theme.LoginIcon())
	myWindow.Show()
}

func postServer(email string, password string) LoginResult {
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
	resp, err := http.Post("http://localhost:8080/mock/status", "application/json", strings.NewReader(string(jsonStr)))
	if err != nil {
		log.Println("访问后台服务报错!")
	}
	// 解析response
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	resRes := LoginResult{}
	json.Unmarshal(body, &resRes)
	return resRes
}
