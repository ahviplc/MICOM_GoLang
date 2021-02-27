package main

import (
	"encoding/gob"
	"fmt"
	"github.com/ahviplc/GoJustToolc/UConsole"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	"mi.com/models"
	_ "mi.com/routers"
)

func init() {
	gob.Register(models.Manager{})
}

func main() {
	// 注册模板函数
	beego.AddFuncMap("unixToDate", models.UnixToDate)
	beego.AddFuncMap("formatImg", models.FormatImg)
	beego.AddFuncMap("formatAttr", models.FormatAttr)
	beego.AddFuncMap("cutStr", models.CutStr)
	beego.AddFuncMap("mul", models.Mul)
	beego.AddFuncMap("judge", models.Judge)

	// 配置session保存在redis里面
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "192.168.0.10:6379"
	// 打印一条直线并换行
	UConsole.PrintAStraightLine()
	// 打印 运行的前台网址信息 ip+端口+###
	fmt.Printf("...前台...http server Running on http://%s:%s", beego.AppConfig.String("httpip"), beego.AppConfig.String("httpport"))
	// 换行用
	fmt.Println()
	// 打印 运行的后台网址信息 ip+端口+###
	fmt.Printf("...后台...http server Running on http://%s:%s/%s", beego.AppConfig.String("httpip"), beego.AppConfig.String("httpport"), "admin")
	// 换行用
	fmt.Println()
	// 打印一条直线并换行
	UConsole.PrintAStraightLine()
	beego.Run()
	defer models.DB.Close()
}
