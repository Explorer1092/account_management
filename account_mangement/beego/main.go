package main

import (
	_ "accout_management/routers"
	"github.com/astaxie/beego"
)

func main() {
	//cron.Init()
	beego.Run()
}

