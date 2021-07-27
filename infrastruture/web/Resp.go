package web

const (
	CodeInvalidParam = 10400
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
	No      int16       `json:"no"`      //页号
}

type PageParam struct {
	No   int16 `json:"no"`   // 页号
	Size int16 `json:"size"` // 分页大小
}

type ErrorResp struct {
	Code    int16                  `json:"code"`   //业务代码
	Message string                 `json:"msg"`    //错误消息
	Errors  map[string]interface{} `json:"errors"` //额外数据
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
