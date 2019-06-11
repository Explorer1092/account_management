package models_record

import (
	"accout_management/database"
	"fmt"
	"time"
)

type weekConsum struct{
	One float32	`json:"one"`
	Two float32  `json:"two"`
	Three float32 `json:"three"`
	Four float32 `json:"four"`
	Five float32 `json:"five"`
	Six float32 `json:"six"`
	Seven float32 `json:"seven"`
}


type weekEarn struct{
	One float32	`json:"one"`
	Two float32  `json:"two"`
	Three float32 `json:"three"`
	Four float32 `json:"four"`
	Five float32 `json:"five"`
	Six float32 `json:"six"`
	Seven float32 `json:"seven"`
}

func GetWeekConsum(Phone_num int64)(dates []weekConsum,err error){
	dates = make([]weekConsum,0)
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d ",Phone_num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=? ",Phone_num).Scan(&id)
	if err != nil{
		return
	}

	var arr [7]float32
	for i:=-6;i<=0;i++ {
		var count1 float32
		timenow := time.Now()
		yes_time := timenow.AddDate(0, 0, i)
		date1 := yes_time.Format("2006-01-02")

		//sql1 := fmt.Sprintf("select sum(price)  from bill_info where date ='%s'  and id=%d	and way ='支出' ", date1,id)

		if err = database.GetDb().QueryRow("select sum(price)  from bill_info where date =?  and id=?	and way ='支出' ", date1,id).Scan(&count1); err != nil {
			err=nil
		}
		arr[i+6]= count1
	}
		fmt.Println("ok")
		var date weekConsum
		date.Seven= arr[0]
		date.Six=arr[1]
		date.Five= arr[2]
		date.Four= arr[3]
		date.Three=arr[4]
		date.Two=arr[5]
		date.One=arr[6]

	dates= append(dates, date)//dates是代表支出

	return
}


func GetWeekEarn(Phone_num int64)(dates []weekEarn,err error){
	dates=make([]weekEarn,0)
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d ",Phone_num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=? ",Phone_num).Scan(&id)
	if err != nil{
		return
	}

	var arr [7]float32
	for i:=-6;i<=0;i++ {
		var count1 float32
		timenow := time.Now()
		yes_time := timenow.AddDate(0, 0, i)
		date1 := yes_time.Format("2006-01-02")

		//sql1 := fmt.Sprintf("select sum(price)  from bill_info where date ='%s'  and id=%d and way ='收入'", date1,id)

		if err = database.GetDb().QueryRow("select sum(price)  from bill_info where date =?  and id=? and way ='收入'", date1,id).Scan(&count1); err != nil {
			err=nil
		}
		arr[i+6]= count1
	}
	fmt.Println("ok")
	var date weekEarn
	date.Seven= arr[0]
	date.Six=arr[1]
	date.Five= arr[2]
	date.Four= arr[3]
	date.Three=arr[4]
	date.Two=arr[5]
	date.One=arr[6]
	dates= append(dates, date)//dates是代表支出
	return
}

func GetMonthConsum(phone_num int64)(res []string,err error){
	res =make([]string,2)
	res[0]="";res[1]=""//res[0]代表收入,res[1]代表支出
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d ",phone_num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=? ",phone_num).Scan(&id)
	if err != nil{
		return
	}

	nowtimes:=time.Now()
	lastmonth:=nowtimes.AddDate(0,-1,0)
	//sql3:=fmt.Sprintf("select sum(price) from bill_info where id =%d and way='收入' and timestamp between %d and %d",id,lastmonth.Unix(),nowtimes.Unix())
	err =database.GetDb().QueryRow("select sum(price) from bill_info where id =? and way='收入' and timestamp between ? and ?",id,lastmonth.Unix(),nowtimes.Unix()).Scan(&res[0])
	if err!=nil{
		return
	}
	//sql2:=fmt.Sprintf("select sum(price) from bill_info where id =%d and way='支出' and timestamp between %d and %d",id,lastmonth.Unix(),nowtimes.Unix())
	err =database.GetDb().QueryRow("select sum(price) from bill_info where id =? and way='支出' and timestamp between ? and ?",id,lastmonth.Unix(),nowtimes.Unix()).Scan(&res[1])
	if err!=nil{
		return
	}
	return res ,nil
}

func GetAnalysis(phone_num int64)(err error,res []string){

	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d ",phone_num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=? ",phone_num).Scan(&id)
	if err != nil{
		return
	}

	res =make([]string,6)
	nowtimes:=time.Now()
	lastmonth:=nowtimes.AddDate(0,-1,0)
	//sql2:=fmt.Sprintf("select sum(price) from bill_info where id=%d and classification='饮食消费' and timestamp between %d and %d",id,lastmonth.Unix(),nowtimes.Unix())
	if err =database.GetDb().QueryRow("select sum(price) from bill_info where id=? and classification='饮食消费' and timestamp between ? and ?",id,lastmonth.Unix(),nowtimes.Unix()).Scan(&res[0]);err!=nil{
		res[0]="0"
	}
	//sql3:=fmt.Sprintf("select sum(price) from bill_info where id=%d and classification='交通消费' and timestamp between %d and %d",id,lastmonth.Unix(),nowtimes.Unix())
	if err =database.GetDb().QueryRow("select sum(price) from bill_info where id=? and classification='交通消费' and timestamp between ? and ?",id,lastmonth.Unix(),nowtimes.Unix()).Scan(&res[1]);err!=nil{
		res[1]="0"
	}
	//sql4:=fmt.Sprintf("select sum(price) from bill_info where id=%d and classification='生活消费' and timestamp between %d and %d",id,lastmonth.Unix(),nowtimes.Unix())
	if err =database.GetDb().QueryRow("select sum(price) from bill_info where id=? and classification='生活消费' and timestamp between ? and ?",id,lastmonth.Unix(),nowtimes.Unix()).Scan(&res[2]);err!=nil{
		res[2]="0"
	}
	//sql5:=fmt.Sprintf("select sum(price) from bill_info where id=%d and classification='教育消费' and timestamp between %d and %d",id,lastmonth.Unix(),nowtimes.Unix())
	if err =database.GetDb().QueryRow("select sum(price) from bill_info where id=? and classification='教育消费' and timestamp between ? and ?",id,lastmonth.Unix(),nowtimes.Unix()).Scan(&res[3]);err!=nil{
		res[3]="0"
	}
	//sql6:=fmt.Sprintf("select sum(price) from bill_info where id=%d and classification='转账' and timestamp between %d and %d",id,lastmonth.Unix(),nowtimes.Unix())
	if err =database.GetDb().QueryRow("select sum(price) from bill_info where id=? and classification='转账' and timestamp between ? and ?",id,lastmonth.Unix(),nowtimes.Unix()).Scan(&res[4]);err!=nil{
		res[4]="0"
	}
	//sql7:=fmt.Sprintf("select sum(price) from bill_info where id=%d and classification='其他' and timestamp between %d and %d",id,lastmonth.Unix(),nowtimes.Unix())
	if err =database.GetDb().QueryRow("select sum(price) from bill_info where id=? and classification='其他' and timestamp between ? and ?",id,lastmonth.Unix(),nowtimes.Unix()).Scan(&res[5]);err!=nil{
		res[5]="0"
	}
	return nil,res
}