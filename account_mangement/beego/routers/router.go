package routers

import (
	"accout_management/controllers/health"
	"accout_management/controllers/information"
	"accout_management/controllers/manage"
	"accout_management/controllers/record"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "TRPM-TOKEN", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))


//information
    beego.Router("/api/health", &health.Health_Controller{},"get:GetHealthInfo")

	beego.Router("/api/Register",&information.PersonInfoController{},"post:Register")

	beego.Router("/api/Login",&information.PersonInfoController{},"post:Login")

	beego.Router("/api/ChangePassword",&information.PersonInfoController{},"post:ChangePassword")

	beego.Router("/api/SetName",&information.PersonInfoController{},"post:SetName")

	beego.Router("/api/GetImg",&information.PersonInfoController{},"get:GetImg")

	beego.Router("/api/GetName",&information.PersonInfoController{},"get:GetName")

	beego.Router("/api/Check_in",&information.PersonInfoController{},"get:CheckIn")

	beego.Router("/api/InputImg",&manage.OperateFund{},"post:InputImg")

	beego.Router("/api/UploadAliExcel",&manage.OperateFund{},"post:UploadAliExcel")

	beego.Router("/api/UploadWechatExcel",&manage.OperateFund{},"post:UploadWechatExcel")

	beego.Router("/api/GetBills",&manage.OperateFund{},"get:GetBills")

	beego.Router("/api/AddBill",&manage.OperateFund{},"post:AddBill")

	beego.Router("/api/RemoveBill",&manage.OperateFund{},"post:RemoveBill")

	beego.Router("/api/BatchRemoveBill",&manage.OperateFund{},"post:BatchRemoveBill")

	beego.Router("/api/GetWeekConsum",&record.FundRecord{},"get:GetWeekConsum")

	beego.Router("/api/GetWeekEarn",&record.FundRecord{},"get:GetWeekEarn")

	beego.Router("/api/GetMonthConsum",&record.FundRecord{},"get:GetMonthConsum")

	beego.Router("/api/GetAnalysis",&record.FundRecord{},"get:GetAnalysis")

	beego.Router("/api/ModifyBill",&manage.OperateFund{},"post:ModifyBill")
}
