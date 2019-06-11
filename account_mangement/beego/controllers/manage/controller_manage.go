package manage

import (
	"accout_management/models"
	"accout_management/models/models_operation"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"

	//"io"
	//"io/ioutil"
)

type OperateFund struct{
	beego.Controller
}

type Bill struct{
	Phone_num string `json:"phone_num"`
	Time string	`json:"time"`
	Date string	`json:"date"`
	Type string	`json:"type"`
	Otherside string `json:"otherside"`
	Goodsname string `json:"goodsname"`
	Way string	`json:"way"`
	Price string `json:"price"`
	Classification string `json:"classification"`
}

type NewBill struct{
	Phone_num string `json:"phone_num"`
	Time string	`json:"time"`
	Date string	`json:"date"`
	Type string	`json:"type"`
	Otherside string `json:"otherside"`
	Goodsname string `json:"goodsname"`
	Way string	`json:"way"`
	Price string `json:"price"`
	Classification string `json:"classification"`
	OldTime string 	`json:"oldtime"`
}

func(this *OperateFund)InputImg(){
	phone_num,err:=this.GetInt64("phone_num")
	f,h,err6 :=this.GetFile("file")
	if err !=nil {
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err.Error()),nil,this.Ctx)
		return
	}
	if err6!=nil {
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err6.Error()),nil,this.Ctx)
		return
	}
	defer f.Close()
	err2:= this.SaveToFile("file","static/img/"+h.Filename)
	if err2 !=nil{
		models.HandleError(models.Upload_Failed,models.GetErrMsg(models.Upload_Failed,err2.Error()),nil,this.Ctx)
		return
	}

    path :=fmt.Sprint("http://39.108.156.137:8000/static/img/",h.Filename)
	err4:=models_operation.Upload_img(phone_num,path)
	if err4 !=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err4.Error()),nil,this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),nil,this.Ctx)
}





func(this*OperateFund)UploadAliExcel(){
	Phone_num,err :=this.GetInt64("phone_num")
	f,h,err2 :=this.GetFile("file")
	if err!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err.Error()),nil,this.Ctx)
		return
	}
	if err2!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err2.Error()),nil,this.Ctx)
		return
	}

	index:=strings.Index(h.Filename,"alipay")
	fmt.Println(index)
	if index ==-1{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,"请上传正确的支付宝账单"),"",nil)
		return
	}


	err3:= this.SaveToFile("file","static/Ali_excel/"+h.Filename)
	fmt.Println(h.Header)
	if err3 !=nil{
		models.HandleError(models.Upload_Failed,models.GetErrMsg(models.Upload_Failed,err3.Error()),nil,this.Ctx)
		return
	}
	defer f.Close()
	err4 :=models_operation.Ali_Read_Function(Phone_num,h.Filename)
	if err4 !=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err4.Error()),nil,this.Ctx)
		fmt.Println(err4)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),nil,this.Ctx)
}


func(this*OperateFund)UploadWechatExcel(){
	Phone_num,err :=this.GetInt64("phone_num")
	f,h,err2 :=this.GetFile("file")
	if err!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err.Error()),nil,this.Ctx)
		return
	}
	if err2!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err2.Error()),nil,this.Ctx)
		return
	}
	index:=strings.Index(h.Filename,"微信")
	if index ==-1{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,"请上传正确的微信账单"),nil,this.Ctx)
		return
	}
	err3:= this.SaveToFile("file","static/Wechat_excel/"+h.Filename)
	fmt.Println(h.Header)
	if err3 !=nil{
		models.HandleError(models.Upload_Failed,models.GetErrMsg(models.Upload_Failed,err3.Error()),nil,this.Ctx)
		return
	}
	defer f.Close()

	err4:=models_operation.Wechat_Upload(Phone_num,h.Filename)
	if err4!=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err4.Error()),nil,this.Ctx)
		fmt.Println(err4)
		return
	}

	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),nil,this.Ctx)
}



func(this*OperateFund)GetBills(){
	Phone_num,err:=this.GetInt64("phone_num")
	Date:=this.GetString("date")

	if err!=nil {
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,""),"",this.Ctx)
		return
	}

	count,bills,err2:=models_operation.GetBills(Phone_num,Date)
	if err2!=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err2.Error()),bills,this.Ctx)
		return
	}

	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),bills,this.Ctx,count)

}


func(this*OperateFund)ModifyBill(){
	newBill :=NewBill{}
	if err:=json.Unmarshal(this.Ctx.Input.RequestBody,&newBill); err!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err.Error()),"",this.Ctx)
		return
	}

	phoneNum2,err3 :=strconv.ParseInt(newBill.Phone_num,10,64)
	if err3!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err3.Error()),"",this.Ctx)
		return
	}
	bill:=models_operation.Bill{
		Phone_num: phoneNum2,
		Time:      newBill.Time,
		Date:      newBill.Date,
		Type:      newBill.Type,
		Price:     newBill.Price,
		Otherside: newBill.Otherside,
		Way:       newBill.Way,
		Goodsname: newBill.Goodsname,
		Classification: newBill.Classification,
	}
	err4:=models_operation.ModifyBill(&bill,newBill.OldTime)
	if err4!=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err4.Error()),nil,this.Ctx)
		fmt.Println(err4.Error())
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),"",this.Ctx)
}

func(this*OperateFund)AddBill(){
	newBill :=Bill{}
	if err:=json.Unmarshal(this.Ctx.Input.RequestBody,&newBill); err!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err.Error()),"",this.Ctx)
		return
	}
	phoneNum2,err3 :=strconv.ParseInt(newBill.Phone_num,10,64)
	if err3!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err3.Error()),"",this.Ctx)
		return
	}
	bill:=models_operation.Bill{
		Phone_num: phoneNum2,
		Time:      newBill.Time,
		Date:      newBill.Date,
		Type:      newBill.Type,
		Price:     newBill.Price,
		Otherside: newBill.Otherside,
		Way:       newBill.Way,
		Goodsname: newBill.Goodsname,
		Classification: newBill.Classification,
	}
	err2 :=models_operation.AddBills(&bill)
	if err2!=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err2.Error()),nil,this.Ctx)
		fmt.Println(err2.Error())
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),"",this.Ctx)
}


func (this*OperateFund)RemoveBill(){
	Phone_num,err:=this.GetInt64("phone_num")
	Time:=this.GetString("time")
	Otherside:=this.GetString("otherside")
	Type:=this.GetString("type")
	GoodsName:=this.GetString("goodsname")
	if err!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err.Error()),"",this.Ctx)
		return
	}
	err2:=models_operation.Remove(Phone_num,Time,Otherside,Type,GoodsName)
	if err2 !=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err2.Error()),nil,this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),"",this.Ctx)
}


func (this*OperateFund)BatchRemoveBill(){
	//var times []string
	phone_num,err:=this.GetInt64("phone_num")
	if err!=nil{
		models.HandleError(models.ERR_ARG,models.GetErrMsg(models.ERR_ARG,err.Error()),"",this.Ctx)
		return
	}
	time:=this.GetString("time")
	err2:=models_operation.BatchRemoveBill(phone_num,time)
	if err2!=nil{
		models.HandleError(models.SQL_ERROR,models.GetErrMsg(models.SQL_ERROR,err2.Error()),nil,this.Ctx)
		return
	}
	models.HandleError(models.SUCCESS,models.GetErrMsg(models.SUCCESS,""),"",this.Ctx)
}
