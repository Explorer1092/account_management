package models_information

import (
	"accout_management/database"
	"fmt"

	"strings"
	"time"
)

func GetName(num int64)(str string,err error){
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d",num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=?",num).Scan(&id)
	if err !=nil{
		return "",nil
	}
	//select_sql := fmt.Sprintf("select name from person_info where id =%d",id)
	err =database.GetDb().QueryRow("select name from person_info where id =?",id).Scan(&str)
	if err !=nil{
		return "",nil
	}
	return str,nil
}

func GetImg(num int64)(str string ,err error){
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d",num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=?",num).Scan(&id)
	if err !=nil{
		return "",nil
	}
	//select_sql := fmt.Sprintf("select img_src from person_info where id =%d",id)
	err =database.GetDb().QueryRow("select img_src from person_info where id =?",id).Scan(&str)
	if err !=nil{
		return "",nil
	}
	return str,nil
}

func Is_Register(num int64)(status bool,err error) {
	if num < 10000000000 || num > 99999999999 { //判断是否为11位电话号码
		return false,fmt.Errorf("非法电话号码")
	}
	//countsql := fmt.Sprintf("select count(*) from verification_info where phone_num =%d",num)
	var count int
	err = database.GetDb().QueryRow("select count(*) from verification_info where phone_num =?",num).Scan(&count)
	if err!= nil{
		return
	}
	if count !=0 {
		return true,nil
	} else{
		return false ,nil
	}
}

//func Make_Token()(str string,err error){
//	mySigningKey := []byte("hzwy23")
//	// Create the Claims
//	claims := &jwt.StandardClaims{
//		NotBefore: int64(time.Now().Unix() - 1000),
//		ExpiresAt: int64(time.Now().Unix() + 1000),
//		Issuer:    "test",
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	ss, err := token.SignedString(mySigningKey)
//
//	t,err := jwt.Parse(ss, func(*jwt.Token) (interface{}, error) {
//		return mySigningKey, nil
//	})
//
//	if err != nil {
//		fmt.Errorf("parse with claims failed.", err)
//		return "",err
//	}
//
//}

func Register(account int64, passwd string )(err error){
	//sql:=fmt.Sprintf("insert into verification_info (phone_num,password)values(%d,'%s')",account,passwd)
	_,err =database.GetDb().Exec("insert into verification_info (phone_num,password)values(?,?)",account,passwd)
	if err != nil{
		return
	}
	var id int
	//sql2 :=fmt.Sprintf("select id from verification_info where phone_num =%d",account)
	if err = database.GetDb().QueryRow("select id from verification_info where phone_num =?",account).Scan(&id);err!=nil{
		return
	}
	//sql3:=fmt.Sprintf("insert into person_info (id)values(%d)",id)
	if _,err =database.GetDb().Exec("insert into person_info (id)values(?)",id);err!=nil{
		return
	}
	//sql4 :=fmt.Sprintf("insert into record_info (id)values(%d)",id)
	if _,err =database.GetDb().Exec("insert into bill_info (id)values(?)",id);err!=nil{
		return
	}
	return nil
}

func Check_Safe(account int64, passwd string)(status bool,err error){
	//sql := fmt.Sprintf("select count(1) from verification_info where phone_num=%d and password=%s",account,passwd )
	var count int
	if err := database.GetDb().QueryRow("select count(1) from verification_info where phone_num=? and password=?",account,passwd).Scan(&count);err != nil{
		return false,err
	}
	if count ==1{
		return true,nil
	} else{
		return false,nil
	}
}


func Change_Password(account int64,new_passwd string)(err error){
	//sql:=fmt.Sprintf("update verification_info set password = %s where phone_num =%d",new_passwd,account)
	_,err = database.GetDb().Exec("update verification_info set password =? where phone_num =?",new_passwd,account)
	return
}


func Set_Name(account int64,name string)(err error){
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d",account)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=?",account).Scan(&id)
	if err != nil{
		return
	}
	//sql2 := fmt.Sprintf("update person_info set name= '%s' where id= %d",name,id)
	_,err = database.GetDb().Exec("update person_info set name= ? where id= ?",name,id)
	return
}



func Check_In(Phone_num int64)(count int ,err error){
	//sql:=fmt.Sprintf("select id from verification_info where phone_num =%d",Phone_num)
	var id int
	err =database.GetDb().QueryRow("select id from verification_info where phone_num =?",Phone_num).Scan(&id)
	if err != nil{
		return -1,err
	}
	//sql2 := fmt.Sprintf("select check_time from person_info where id =%d",id)
	var sql_time string //example:2006-01-02 (string)
	err = database.GetDb().QueryRow("select check_time from person_info where id =?",id).Scan(&sql_time)
	if err != nil {
		return -1,err
	}
	nowdays:=time.Now().Format("2006-01-02 15:04:05")
	var start_time string
	if times:=strings.Split(nowdays," ");len(times)==2{
		start_time =times[0]+" 00:00:00"
	}else{
		return -1,fmt.Errorf("%s","split error!")
	}
	Format_start_time,err := time.Parse("2006-01-02 15:04:05",start_time)
	if err !=nil{
		return -1,err
	}
	Format_end_time:=Format_start_time.AddDate(0,0,1)

	Format_sql_time,err:=time.Parse("2006-01-02 15:04:05",sql_time)
	if err !=nil{
		return -1,err
	}
	if Format_sql_time.After(Format_start_time) && Format_sql_time.Before(Format_end_time){
		return -1,fmt.Errorf("%s","今天已经签到过了!")
	}

	new_sql_time:=time.Now().Format("2006-01-02 15:04:05")
	//order_sql:=fmt.Sprintf("update person_info set check_time='%s' , check_in=check_in+1 where id =%d",new_sql_time,id)
	if _,err = database.GetDb().Exec("update person_info set check_time=? , check_in=check_in+1 where id =?",new_sql_time,id);err!=nil{
		return -1,err
	}
	//select_sql := fmt.Sprintf("select check_in from person_info where id =%d",id)
	err =database.GetDb().QueryRow("select check_in from person_info where id =?",id).Scan(&count)
	if err !=nil{
		return -1,err
	}
	return count,nil
}






