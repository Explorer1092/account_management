package models

import(
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
)

const(
	SUCCESS       =200
	ERR_ARG       = 4000
	ERR_SERVER    = 5000
	Account_Num_Error   =4001
	Account_Passwd_Error = 4002
	Paying_Passwd_Error = 4003
	User_Not_Exist_ERROR = 4004
	User_Exist_ERROR =4005
	Make_New_Card_Error =4006
	Upload_Failed=4007
	SQL_ERROR = 5001
)

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type NewResp struct {
	Code      int         `json:"code"`
	Message   string      `json:"msg"`
	TotalPage int         `json:"totalPage"`
	Data      interface{} `json:"data"`
}

func GetErrMsg(errcode int, key string) string {
	switch errcode {
	case SUCCESS:
		return  fmt.Sprintf("ok!")
	case ERR_ARG:
		return fmt.Sprintf("Arg err: %s",key)
	case Account_Num_Error:
		return fmt.Sprintf("用户账号错误:%s",key)
	case Account_Passwd_Error:
		return fmt.Sprintf("用户密码错误:%s",key)
	case Paying_Passwd_Error:
		return fmt.Sprintf("用户支付密码错误:%s",key)
	case User_Not_Exist_ERROR:
		return fmt.Sprintf("用户不存在:%s",key)
	case User_Exist_ERROR:
		return fmt.Sprintf("用户已注册:%s",key)
	case Make_New_Card_Error:
		return fmt.Sprintf("办理新银行卡失败:%s",key)
	case Upload_Failed:
		return fmt.Sprintf("上传失败:%s",key)
	case ERR_SERVER:
		return fmt.Sprintf("server err:%s",key)
	case SQL_ERROR:
		return fmt.Sprintf("数据库错误:%s",key)
	}
	return ""
}


func HandleError(status int, message string, data interface{}, ctx *context.Context, totalPage ...int) string {
	var resp interface{}
	if len(totalPage) == 0 {
		resp = Resp{
			Code:    status,
			Message: message,
			Data:    data,
		}
	} else {
		resp = NewResp{
			Code:      status,
			Message:   message,
			TotalPage: totalPage[0],
			Data:      data,
		}
	}
	res, err := json.Marshal(resp)
	if err != nil {
		return ""
	}

	if ctx != nil {
		ctx.ResponseWriter.Header().Set("Content-Type", "application/json;charset=utf-8")
		ctx.WriteString(string(res))
	}

	return string(res)
}
