package health

import (
	"accout_management/models"
	"github.com/astaxie/beego"
)

type Health_Controller struct {
	beego.Controller
}

func(this* Health_Controller)GetHealthInfo(){
	beego.Info(this.Ctx.Input.Header("Content-Type"))
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),nil,this.Ctx)
}


