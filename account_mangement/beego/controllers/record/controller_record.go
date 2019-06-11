package record

import (
	"accout_management/models"
	"accout_management/models/models_record"
	"github.com/astaxie/beego"
)

type FundRecord struct{
	beego.Controller
}


func (this*FundRecord)GetWeekConsum(){
	phone_num,err:=this.GetInt64("phone_num")
	if err!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err.Error()),"",this.Ctx)
		return
	}
	date ,err2:=models_record.GetWeekConsum(phone_num)
	if err2!=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err2.Error()),"",this.Ctx)
		return
	}


	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),date,this.Ctx)
}


func(this*FundRecord)GetWeekEarn(){
	phone_num,err:=this.GetInt64("phone_num")
	if err!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err.Error()),"",this.Ctx)
		return
	}
	date2,err3:=models_record.GetWeekEarn(phone_num)
	if err3!=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err3.Error()),"",this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),date2,this.Ctx)
}


func(this*FundRecord)GetMonthConsum(){
	phone_num,err:=this.GetInt64("phone_num")
	if err!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err.Error()),"",this.Ctx)
		return
	}
	res,err2:=models_record.GetMonthConsum(phone_num)
	if err2!=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err2.Error()),"",this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),res,this.Ctx)
}

func (this*FundRecord)GetAnalysis(){
	phone_num,err:=this.GetInt64("phone_num")
	if err!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err.Error()),"",this.Ctx)
		return
	}
	err2 ,res:=models_record.GetAnalysis(phone_num)
	if err2!=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err2.Error()),"",this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),res,this.Ctx)
}