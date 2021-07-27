package web

import "github.com/kataras/iris/v12/mvc"



func WrapPage(records interface{},total int64,current int16) mvc.Result {
	page := PageResp {
		Records: records,
		Total: total,
		No: current,
	}

	return mvc.Response{
		ContentType: "application/json",
		Object: page,
	}
}