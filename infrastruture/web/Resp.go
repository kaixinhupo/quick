package web

import (
	e "github.com/kaixinhupo/quick/infrastruture/error"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)


const (
	CodeInvalidParam = 10400
	CodeUnknowError = 10500
)

type ValidationError struct {
	ActualTag string `json:"tag"`       // 规则
	Namespace string `json:"namespace"` // 属性
	Type      string `json:"type"`      // 类型
	Value     string `json:"value"`     // 值
	Param     string `json:"param"`     // 参数
}

type PageResp struct {
	Records interface{} `json:"records"` //记录
	Total   int64       `json:"total"`   //总记录数
	No      int       	`json:"no"`      //页号
}

type PageParam struct {
	No   int `json:"no"`   // 页号
	Size int `json:"size"` // 分页大小
}

type ErrorResp struct {
	Code    int16                  `json:"code"`   //业务代码
	Message string                 `json:"msg"`    //消息
	Errors  map[string]interface{} `json:"errors"` //额外数据
}

type SuccessResp struct {
	Code    int16  `json:"code"` //业务代码
	Message string `json:"msg"`  //消息
}

func NewErrorResp() ErrorResp {
	return ErrorResp{
		Errors: make(map[string]interface{}),
	}
}

func (e ErrorResp) WithCode(code int16) ErrorResp {
	e.Code = code
	return e
}

func (e ErrorResp) WithMessage(msg string) ErrorResp {
	e.Message = msg
	return e
}

func (e ErrorResp) WithErrors(errors map[string]interface{}) ErrorResp {
	for k, v := range errors {
		e.Errors[k] = v
	}
	return e
}

func (e ErrorResp) AppendError(key string, value interface{}) ErrorResp {
	e.Errors[key] = value
	return e
}

func WrapPage(records interface{}, total int64, current int) mvc.Result {
	page := PageResp{
		Records: records,
		Total:   total,
		No:      current,
	}

	return mvc.Response {
		ContentType: "application/json",
		Object:      page,
	}
}

func WrapSuccess() mvc.Result {
	resp := SuccessResp{
		Code:    0,
		Message: "success",
	}

	return mvc.Response{
		ContentType: "application/json",
		Object:      resp,
	}
}

func WrapError(err error) mvc.Result {
	resp := NewErrorResp()

	if err, ok := err.(*e.BizError); ok {
		resp.Code = int16(err.Code)
		resp.Message = err.Message
	} else {
		resp.Code = CodeUnknowError
		resp.Message ="未知错误"
		resp.AppendError("err",err)
	}
	return mvc.Response{
		ContentType: "application/json",
		Code: iris.StatusBadRequest,
		Object: resp,
	}
}

func WrapResp(resp interface{}) mvc.Result {
	return mvc.Response{
		ContentType: "application/json",
		Object:      resp,
	}
}