package main

import (
	"fmt"
	"myQQ/login"
	"os"
	"strings"

	"fyne.io/fyne/v2/app"
	"github.com/flopp/go-findfont"
)

func main() {
	a := app.New()
	login.Login(a)
	a.Run()
}

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simhei.ttf") {
			fmt.Println(path)
			os.Setenv("FYNE_FONT", path) // 设置环境变量  // 取消环境变量 os.Unsetenv("FYNE_FONT")
			break
		}
	}
}
