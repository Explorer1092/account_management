package information

import (
	"accout_management/models"
	"accout_management/models/models_information"
	"github.com/astaxie/beego"
)

type PersonInfoController struct{
	beego.Controller
}



func (this* PersonInfoController)Register(){
	Phone_num,err:=this.GetInt64("username")
	Login_passwd := this.GetString("password")
	if err != nil {
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,"请正确输入"),nil,this.Ctx)
		return
	}

	status,err := models_information.Is_Register(Phone_num)
	if err != nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err.Error()),nil,this.Ctx)
		return
	}
	if status{
		models.HandleError(models.User_Exist_ERROR,models.GetErrMsg(models.User_Exist_ERROR,""),nil,this.Ctx)
		return
	}
	err =models_information.Register(Phone_num,Login_passwd)
	if err != nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err.Error()),nil,this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),Phone_num,this.Ctx)
}


func(this* PersonInfoController)Login() {
	Phone_num ,err:=this.GetInt64("phone_num")
	Login_passwd := this.GetString("login_passwd")

	if err != nil {
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,"请正确输入"),nil,this.Ctx)
		return
	}
	safe,err :=models_information.Check_Safe(Phone_num,Login_passwd)
	if err != nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err.Error()),nil,this.Ctx)
		return
	}
	if safe{

		models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),"true",this.Ctx)
		return
	}else{
		var status bool
		if status,err =models_information.Is_Register(Phone_num);err!=nil{
			models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err.Error()),nil,this.Ctx)
			return
		}
		if !status {
			models.HandleError(models.User_Not_Exist_ERROR, models.GetErrMsg(models.User_Not_Exist_ERROR, ""), nil, this.Ctx)
			return
		}else{
			models.HandleError(models.Account_Passwd_Error,models.GetErrMsg(models.Account_Passwd_Error,""),nil,this.Ctx)
			return
		}
	}
}

func (this*PersonInfoController)GetName(){
	Phone_num, err := this.GetInt64("phone_num")
	if err != nil{
		models.HandleError(models.ERR_ARG, models.GetErrMsg(models.ERR_ARG, err.Error()), nil, this.Ctx)
		return
	}
	name,err2:=models_information.GetName(Phone_num)
	if err2 !=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err.Error()),nil,this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),name,this.Ctx)
}


func (this*PersonInfoController)GetImg(){
	Phone_num, err := this.GetInt64("phone_num")
	if err != nil{
		models.HandleError(models.ERR_ARG, models.GetErrMsg(models.ERR_ARG, err.Error()), nil, this.Ctx)
		return
	}
	img_path,err2 :=models_information.GetImg(Phone_num)
	if err2 !=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err2.Error()),nil,this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),img_path,this.Ctx)
}

func (this*PersonInfoController)ChangePassword() {
	Phone_num, err := this.GetInt64("username")
	Old_passwd := this.GetString("oldpwd")
	New_passwd := this.GetString("newpwd")
	if err != nil {
		models.HandleError(models.ERR_ARG, models.GetErrMsg(models.ERR_ARG, err.Error()), nil, this.Ctx)
		return
	}
	safe, err := models_information.Check_Safe(Phone_num, Old_passwd)
	if err != nil {
		models.HandleError(models.SQL_ERROR, models.GetErrMsg(models.SQL_ERROR, err.Error()), nil, this.Ctx)
		return
	}
	if !safe {
		var status bool
		if status, err = models_information.Is_Register(Phone_num); err != nil {
			models.HandleError(models.SQL_ERROR, models.GetErrMsg(models.SQL_ERROR, err.Error()), nil, this.Ctx)
			return
		}
		if !status {
			models.HandleError(models.User_Not_Exist_ERROR, models.GetErrMsg(models.User_Not_Exist_ERROR, ""), nil, this.Ctx)
			return
		} else {
			models.HandleError(models.Account_Passwd_Error, models.GetErrMsg(models.Account_Passwd_Error, ""), nil, this.Ctx)
			return
		}
	}
	err3 := models_information.Change_Password(Phone_num, New_passwd)
	if err3 != nil {
		models.HandleError(models.SQL_ERROR, models.GetErrMsg(models.SQL_ERROR, err.Error()), nil, this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS, models.GetErrMsg(models.SUCCESS, ""), nil, this.Ctx)

}

func (this*PersonInfoController)SetName(){
	Phone_num,err:= this.GetInt64("phone")
	Name:=this.GetString("full_name")
	if err != nil{
		models.HandleError(models.ERR_ARG, models.GetErrMsg(models.ERR_ARG, err.Error()), nil, this.Ctx)
		return
	}
	err2:=models_information.Set_Name(Phone_num,Name)
	if err2 != nil{
		models.HandleError(models.SQL_ERROR, models.GetErrMsg(models.SQL_ERROR, err2.Error()), nil, this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),nil,this.Ctx)
}

func (this*PersonInfoController)CheckIn(){
	PhoneNum,err:=this.GetInt64("Phone_num")
	if err !=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,"1:"+err.Error()),nil,this.Ctx)
		return
	}
	count ,err2 :=models_information.Check_In(PhoneNum)
	if err2 !=nil{
		models.HandleError(models.ERR_SERVER,models.GetErrMsg(models.ERR_ARG,err2.Error()),nil,this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),count,this.Ctx)
	return
}



