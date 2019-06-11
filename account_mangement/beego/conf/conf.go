package conf

import (
	"fmt"
	"github.com/astaxie/beego"
)


var(
//mysql
	MyDbUser string
	MyDbPwd  string
	MyDbHost string
	MyDbPort string
	MyDbName string
	TaskCron string
	Dev_Aliuploadpath string
	Dev_Wechatuploadpath string

)

func init(){
	if MyDbUser = beego.AppConfig.String("mysql::mydbuser"); MyDbUser == "" {
		fmt.Println("MyDbUser path is not set , default is ", MyDbUser)
		panic("MyDbUser is not set , default is null")
	}
	if MyDbPwd = beego.AppConfig.String("mysql::mydbpwd"); MyDbPwd == "" {
		fmt.Println("MyDbPwd path is not set , default is ", MyDbPwd)
		//panic("MyDbPwd is not set , default is null")
	}
	if MyDbHost = beego.AppConfig.String("mysql::mydbhost"); MyDbHost == "" {
		fmt.Println("MyDbHost path is not set , default is ", MyDbHost)
		panic("MyDbHost is not set , default is null")
	}
	if MyDbPort = beego.AppConfig.String("mysql::mydbport"); MyDbPort == "" {
		fmt.Println("MyDbPort path is not set , default is ", MyDbPort)
		panic("MyDbPort is not set , default is null")
	}
	if MyDbName = beego.AppConfig.String("mysql::mydbname"); MyDbName == "" {
		fmt.Println("MyDbName path is not set , default is ", MyDbName)
		panic("MyDbName is not set , default is null")
	}
	if Dev_Aliuploadpath =beego.AppConfig.String("path::dev_aliuploadpath");Dev_Aliuploadpath==""{
		fmt.Println("Dev_uploadpath path is not set ,default is",Dev_Aliuploadpath)
		panic("Dev_uploadpath path is not set,default is null")
	}
	if Dev_Wechatuploadpath=beego.AppConfig.String("path::dev_wechatuploadpath");Dev_Wechatuploadpath==""{
		fmt.Println("Dev_Wechatuploadpath is not set ,default is ",Dev_Wechatuploadpath)
		panic("Dev_Wechatuploadpath path is not set,default is null")
	}

}