package models_operation

import (
	"accout_management/conf"
	_ "accout_management/conf"
	"accout_management/database"
	"bufio"
	_ "database/sql"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)


type Charset string
const(
	UTF8 =Charset("UTF-8")
	GB18030 =Charset("GB18030")
)



type Bill struct{
	Phone_num int64
	Id int
	Time string
	Date string
	Type string
	Otherside string
	Goodsname string
	Way string
	Price string
	Classification string
	Timestamp int64
}


func Analysis(goodsname string ,otherside string)(billType string ){
	info:=fmt.Sprintf("%s %s",goodsname,otherside)

	foodConsum:=[]string{"饭","餐","菜","肉","蛋","奶","汁","果","水","饮料","食品","肯德基","麦当劳","必胜客","蜜雪冰城","海底捞","烧烤","串串","火锅","炒","面","米","粉","饺子","馍","汤"}
	trafficConsum:=[]string{"地铁","公交","车","滴滴","12306","交通","航","机","高铁"}
	transfer:=[]string{"转账"}
	educationConsum:=[]string{"书","本","电脑","当当","笔","大学","中学","高中","幼儿园","研究","学费","教育","学习"}
	lifeConsum:=[]string{"淘宝","京东","拼多多","购物","超市","商城","充值","费用","娱乐","游戏","旅游","付款","商品","阿里","微信","支付宝"}

	for i:=0;i<len(foodConsum);i++{
		status:=strings.Index(info,foodConsum[i])
		if status !=-1{
			billType="饮食消费"
			return
		}
	}
	for i:=0;i<len(trafficConsum);i++{
		status:=strings.Index(info,trafficConsum[i])
		if status !=-1{
			billType="交通消费"
			return
		}
	}
	for i:=0;i<len(transfer);i++{
		status:=strings.Index(info,transfer[i])
		if status !=-1{
			billType="转账"
			return
		}
	}
	for i:=0;i<len(educationConsum);i++{
		status:=strings.Index(info,educationConsum[i])
		if status !=-1{
			billType="教育消费"
			return
		}
	}
	for i:=0;i<len(lifeConsum);i++{
		status:=strings.Index(info,lifeConsum[i])
		if status!=-1{
			billType="生活消费"
			return
		}
	}
	billType="其他"
	return
}


func Upload_img(phone_num int64,path string)(err error){
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d",phone_num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=?",phone_num).Scan(&id)
	if err != nil{
		return fmt.Errorf("1=> %s",err)
	}

	//update_sql:=fmt.Sprintf("update  person_info set img_src ='%s' where id =%d",path,id)
	_,err =database.GetDb().Exec("update  person_info set img_src =? where id =?",path,id)
	if err!=nil{
		return err
	}
	return nil
}





//
func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func GetData(phone_num int64,file_name string)(err error){
	var bill []Bill
	f,err:=os.Open(file_name)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	defer f.Close()
	rd:= bufio.NewReader(f)
	for {
		line,err :=rd.ReadString('\n')
		if err!=nil || io.EOF ==err{
			break
		}
		res :=strings.Split(line,",")
		if len(res)==17 {
			var res_bill Bill

			result:=strings.Split(res[2]," ")

			res_bill.Date=result[0]
			res_bill.Time = result[1]
			res_bill.Type = "支付宝"
			res_bill.Otherside = res[7]
			res_bill.Goodsname = res[8]
			res_bill.Price = res[9]
			res_bill.Way = res[10]
			res_bill.Classification=Analysis(res_bill.Goodsname,res_bill.Otherside)
			bill =append(bill,res_bill)
		}
	}
	if len(bill)<2{
		fmt.Println(len(bill))
		return nil
	}
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d",phone_num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=?",phone_num).Scan(&id)
	if err != nil{
		return fmt.Errorf("1=> %s",err)
	}

	for i:=1;i<len(bill);i++{
		var existcount int
		//select_sql:=fmt.Sprintf("select count(1) from bill_info where id=%d and date='%s' and time='%s' and type='支付宝' ",id,bill[i].Date,bill[i].Time)
		if err=database.GetDb().QueryRow("select count(1) from bill_info where id=? and date=? and time=? and type='支付宝' ",id,bill[i].Date,bill[i].Time).Scan(&existcount);err!=nil{
			return
		}
		if existcount !=0{
			continue
		}
		nowtimes,err:=time.Parse("2006-01-02 15:04:05",fmt.Sprintf("%s %s",bill[i].Date,bill[i].Time))
		if err!=nil{
			return fmt.Errorf("time parse failed!%s,%s",bill[i].Date,bill[i].Time)
		}
		bill[i].Timestamp=nowtimes.Unix()
		//sql:=fmt.Sprintf("insert into bill_info(id,time,date,type,otherside,goodsname,way,price,timestamp,classification)values(%d,'%s','%s','%s','%s','%s','%s','%s',%d,'%s')",id,bill[i].Time,bill[i].Date,bill[i].Type,bill[i].Otherside,bill[i].Goodsname,bill[i].Way, bill[i].Price,bill[i].Timestamp,bill[i].Classification)
		_,err=database.GetDb().Exec("insert into bill_info(id,time,date,type,otherside,goodsname,way,price,timestamp,classification)values(?,?,?,?,?,?,?,?,?,?)",id,bill[i].Time,bill[i].Date,bill[i].Type,bill[i].Otherside,bill[i].Goodsname,bill[i].Way, bill[i].Price,bill[i].Timestamp,bill[i].Classification)
		if err !=nil{
			continue
		}
	}
	return nil
}

func Ali_Read_Function(phone_num int64,file_name string)(err error){

	//_file_name :="../../static/Ali_excel/"+file_name
	_file_name :=fmt.Sprintf("%s/%s",conf.Dev_Aliuploadpath,file_name)
	fmt.Println(_file_name)
	file,err:=os.Open(_file_name)
	if err!=nil{
		fmt.Println("file open failed!")
		return
	}
	input,_:=ioutil.ReadAll(file)
	data :=ConvertByte2String(input,GB18030)

	new_name:=fmt.Sprintf("2-%s",file_name)
	err =ioutil.WriteFile(new_name, []byte(data),0777)
	if err!=nil{
		return
	}
	err =GetData(phone_num,new_name)
	if err !=nil{
		return
	}
	return nil
}

func GetWechatData(phone_num int64,file_name string)(err error){
	var bill []Bill
	f,err:=os.Open(file_name)
	if err!=nil{
		fmt.Println(err.Error()+"文件打不开")
		return err
	}
	defer f.Close()
	rd:= bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		res := strings.Split(line, ",")
		fmt.Println(res)
		if (len(res) == 11) {
			var res_bill Bill
			res_bill.Date =res[0]
			res_bill.Type = "微信"
			res_bill.Otherside = res[2]
			res_bill.Goodsname = res[3]
			res_bill.Way = res[4]
			res_bill.Price = res[5]
			res_bill.Classification=Analysis(res_bill.Goodsname,res_bill.Otherside)
			bill = append(bill, res_bill)
		}
	}

		fmt.Println(len(bill))
		if len(bill)<2{
			fmt.Println(len(bill))
			return nil
		}
		var id int
		//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d",phone_num)
		err =database.GetDb().QueryRow("select id from verification_info where phone_num=?",phone_num).Scan(&id)
		if err != nil{
			return fmt.Errorf("1=> %s",err)
		}

		for i:=1;i<len(bill);i++{

			res :=strings.Split(bill[i].Date," ")
			nowtimes,err:=time.Parse("2006-01-02 15:04:05",fmt.Sprintf("%s %s",res[0],res[1]))
			if err!=nil{
				return fmt.Errorf("time parse failed!")
			}

			bill[i].Timestamp=nowtimes.Unix()
			var existcount int
			//select_sql:=fmt.Sprintf("select count(1) from bill_info where id=%d and date='%s' and time='%s' and type='微信'",id,res[0],res[1])
			if err=database.GetDb().QueryRow("select count(1) from bill_info where id=? and date=? and time=? and type='微信'",id,res[0],res[1]).Scan(&existcount);err!=nil{
				return fmt.Errorf("scan err")
			}
			if existcount !=0{
				continue
			}

			res_price :=strings.Split(bill[i].Price,"¥")
			if (len(res_price) !=2){
				return fmt.Errorf("strings.Split failed!")

			}

			//sql := fmt.Sprintf("insert into bill_info(id,time,date,type,otherside,goodsname,way,price,timestamp,classification)values(%d,'%s','%s','%s','%s','%s','%s','%s',%d,'%s')", id, res[1], res[0], bill[i].Type, bill[i].Otherside, bill[i].Goodsname, bill[i].Way, res_price[1],bill[i].Timestamp,bill[i].Classification)
			_, err = database.GetDb().Exec("insert into bill_info(id,time,date,type,otherside,goodsname,way,price,timestamp,classification)values(?,?,?,?,?,?,?,?,?,?)", id, res[1], res[0], bill[i].Type, bill[i].Otherside, bill[i].Goodsname, bill[i].Way, res_price[1],bill[i].Timestamp,bill[i].Classification)
			if err != nil {
				continue
			}
		}
		return nil
}

func Wechat_Upload(phone_num int64,file_name string)(err error){
	_file_name:=fmt.Sprintf("%s/%s",conf.Dev_Wechatuploadpath,file_name)
	fmt.Println(_file_name)

	if err!=nil{
		fmt.Println("file open failed!")
		return
	}
	err =GetWechatData(phone_num,_file_name)
	if err!=nil{
		return err
	}
	return nil
}


func GetBills(phone_num int64, date string)(count int,bills []Bill,err error){
	bills =make([]Bill,0)
	//limitsql := " limit " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d ",phone_num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=? ",phone_num).Scan(&id)
	if err != nil{
		return 0,bills,err
	}

	if date ==""{
		//sql:=fmt.Sprintf("select *from bill_info where id=%d order by date desc, time desc",id)
		rows,err:=database.GetDb().Query("select *from bill_info where id=? order by timestamp desc",id)
		if err !=nil{
			return 0,bills,err
		}

		defer rows.Close()

		for rows.Next(){
			var bill Bill
			if err=rows.Scan(&bill.Id,&bill.Time,&bill.Type,&bill.Otherside,&bill.Goodsname,&bill.Way,&bill.Price,&bill.Date,&bill.Timestamp,&bill.Classification);err!=nil{
				fmt.Println(err)
				continue
			}
			bill.Time=fmt.Sprintf("%s /%s",bill.Date,bill.Time)
			bills=append(bills,bill)
		}
		//countsql:=fmt.Sprintf("select count(1) from bill_info where id=%d ",id)
		err =database.GetDb().QueryRow("select count(1) from bill_info where id=? ",id).Scan(&count)
		if err!=nil{
			return 0,bills,err
		}
		return count,bills,nil
	}
		//sql2:=fmt.Sprintf("select *from bill_info where id=%d and date='%s' order by timestamp desc ",id,date)
		rows,err:=database.GetDb().Query("select *from bill_info where id=? and date=? order by time desc ",id,date)
		if err!=nil{
			return
		}

		defer rows.Close()
		for rows.Next(){
			var bill Bill
			if err=rows.Scan(&bill.Id,&bill.Time,&bill.Type,&bill.Otherside,&bill.Goodsname,&bill.Way,&bill.Price,&bill.Date,&bill.Timestamp,&bill.Classification);err!=nil{
				fmt.Println(err)
				continue
			}
			bill.Time=fmt.Sprintf("%s /%s",bill.Date,bill.Time)
			bills=append(bills,bill)
		}
		//countsql:=fmt.Sprintf("select count(1) from bill_info where id=%d and date= '%s' ",id,date)
		err =database.GetDb().QueryRow("select count(1) from bill_info where id=? and date= ? ",id,date).Scan(&count)
		if err!=nil{
			return 0,bills,err
		}
	return
}
func ModifyBill(newBill*Bill,oldtime string )(err error) {
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d ",newBill.Phone_num)
	err = database.GetDb().QueryRow("select id from verification_info where phone_num=? ", newBill.Phone_num).Scan(&id)
	if err != nil {
		return
	}
	olddata:=strings.Split(oldtime," /")
	if len(olddata)==2 {
		old_date:=olddata[0]
		old_time:=olddata[1]

		res_date := strings.Split(newBill.Date, "/")
		if len(res_date) != 3 {
			fmt.Println("split failed!")
			return fmt.Errorf("split failed!")
		}
		tmp, err := strconv.Atoi(res_date[1])
		if err != nil {
			return err
		}
		if tmp < 10 {
			res_date[1] = fmt.Sprintf("0%s", res_date[1])
		}
		tmp, err = strconv.Atoi(res_date[2])
		if err != nil {
			return err
		}
		if tmp < 10 {
			res_date[2] = fmt.Sprintf("0%s", res_date[2])
		}
		newBill.Date = fmt.Sprintf("%s-%s-%s", res_date[0], res_date[1], res_date[2])

		nowtimes, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", newBill.Date, newBill.Time))
		if err != nil {
			return err
		}
		newBill.Timestamp = nowtimes.Unix()

		_, err = database.GetDb().Exec("update bill_info set time=?,date=?,type=?,otherside=?,goodsname=?,way=?,price=?,classification=?,timestamp=? where id=? and  date=? and time=? ", newBill.Time, newBill.Date, newBill.Type, newBill.Otherside, newBill.Goodsname, newBill.Way, newBill.Price, newBill.Classification, newBill.Timestamp, id, old_date, old_time)
		if err != nil {
			return err
		}

	}else{
		fmt.Println(oldtime)
		fmt.Println(len(oldtime))
		return fmt.Errorf("len(oldtime)--->%d",len(oldtime))
	}
	return
}

func AddBills( newBill *Bill)(err error){
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d ",newBill.Phone_num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=? ",newBill.Phone_num).Scan(&id)
	if err != nil{
		return
	}

	res_date:=strings.Split(newBill.Date,"/")
	if len(res_date) !=3{
		fmt.Println("split failed!")
		return fmt.Errorf("split failed!")
	}
	tmp,err:=strconv.Atoi(res_date[1])
	if err!=nil{
		return
	}
	if tmp<10{
		res_date[1]=fmt.Sprintf("0%s",res_date[1])
	}
	tmp,err =strconv.Atoi(res_date[2])
	if err!=nil{
		return
	}
	if tmp<10{
		res_date[2]=fmt.Sprintf("0%s",res_date[2])
	}
	newBill.Date=fmt.Sprintf("%s-%s-%s",res_date[0],res_date[1],res_date[2])

	nowtimes,err :=time.Parse("2006-01-02 15:04:05",fmt.Sprintf("%s %s",newBill.Date,newBill.Time))
	if err!=nil{
		return
	}
	newBill.Timestamp=nowtimes.Unix()
	var existCount int
	//select_sql:=fmt.Sprintf("select count(1) from bill_info where date='%s' and time='%s' and type='%s' and price='%s' and goodsname ='%s'",newBill.Date,newBill.Time,newBill.Type,newBill.Price,newBill.Goodsname)
	err =database.GetDb().QueryRow("select count(1) from bill_info where id=? and  date=? and time=? and type=? and price=? and goodsname =?",id,newBill.Date,newBill.Time,newBill.Type,newBill.Price,newBill.Goodsname).Scan(&existCount)
	if err!=nil{
		return
	}
	if existCount !=0 {
		_,err =database.GetDb().Exec("update bill_info set time=?,date=?,type=?,otherside=?,goodsname=?,way=?,price=?,classification=?,timestamp=? where id=? and  date=? and time=? and type=? and price=? and goodsname =?",newBill.Time,newBill.Date,newBill.Type,newBill.Otherside,newBill.Goodsname,newBill.Way, newBill.Price,newBill.Classification,newBill.Timestamp,id,newBill.Date,newBill.Time,newBill.Type,newBill.Price,newBill.Goodsname)
		if err!=nil{
			return
		}
	}

	//sql2:=fmt.Sprintf("insert into bill_info(id,time,date,type,otherside,goodsname,way,price,classification,timestamp) values(%d,'%s','%s','%s','%s','%s','%s','%s','%s',%d)",id,newBill.Time,newBill.Date,newBill.Type,newBill.Otherside,newBill.Goodsname,newBill.Way, newBill.Price,newBill.Classification,newBill.Timestamp)
	_,err =database.GetDb().Exec("insert into bill_info(id,time,date,type,otherside,goodsname,way,price,classification,timestamp) values(?,?,?,?,?,?,?,?,?,?)",id,newBill.Time,newBill.Date,newBill.Type,newBill.Otherside,newBill.Goodsname,newBill.Way, newBill.Price,newBill.Classification,newBill.Timestamp)
	if err!=nil{
		return err
	}
	return nil
}


func Remove(Phone_num int64,Time string ,Otherside string,Type string,Goodsname string)(err error){
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d ",Phone_num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=? ",Phone_num).Scan(&id)
	if err != nil{
		return
	}
	fmt.Println(Time,Otherside,Type,Goodsname)
	var date,time string
	res:=strings.Split(Time," /")
	if len(res) ==2{
		date =res[0]
		time =res[1]
	}
	//sql2:=fmt.Sprintf("delete from bill_info where id =%d AND time ='%s' AND otherside ='%s' AND type='%s' AND goodsname='%s' And date='%s'",id,time,Otherside,Type,Goodsname,date)
	_,err =database.GetDb().Exec("delete from bill_info where id =? AND time =? AND otherside =? AND type=? AND goodsname=? And date=?",id,time,Otherside,Type,Goodsname,date)
	if err!=nil{
		return err
	}
	return nil
}


func BatchRemoveBill(phone_num int64,time string)(err error){
	var id int
	//sql:=fmt.Sprintf("select id from verification_info where phone_num=%d ",phone_num)
	err =database.GetDb().QueryRow("select id from verification_info where phone_num=? ",phone_num).Scan(&id)
	if err != nil{
		return
	}
	var date []string
	var sql_time []string
	res_time:=strings.Split(time,",")


	for i:=0;i<len(res_time);i++{
		res:=strings.Split(res_time[i]," /")
		if len(res)==2{
			date=append(date,res[0])
			sql_time=append(sql_time,res[1])
		}
	}

	for i:=0;i<len(date)&&i<len(sql_time);i++{
		//insert_sql:=fmt.Sprintf("delete from bill_info where id=%d AND date='%s' AND time='%s'",id,date[i],sql_time[i])
		_,err =database.GetDb().Exec("delete from bill_info where id=? AND date=? AND time=?",id,date[i],sql_time[i])
		if err!=nil{
			fmt.Println(err)
		}
	}
	if err!=nil{
		return
	}
	return nil
}
