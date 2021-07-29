package web

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

var _validator =  validator.New()	

func ValidateRequest(req interface{}) mvc.Result {
	err:= _validator.Struct(req)

	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors := wrapValidationErrors(errs)
			return	mvc.Response {
				Code: iris.StatusUnprocessableEntity,
				Object: NewErrorResp().
					WithMessage("参数错误").
					WithCode(CodeInvalidParam).
					AppendError("detail",validationErrors),
			}
		}
	}
	return nil
}


func wrapValidationErrors(errs validator.ValidationErrors) []ValidationError {
	validationErrors := make([]ValidationError, 0, len(errs))
	for _, validationErr := range errs {
		validationErrors = append(validationErrors, ValidationError{
			ActualTag: validationErr.ActualTag(),
			Namespace: validationErr.Namespace(),
			Type:      validationErr.Type().String(),
			Value:     fmt.Sprintf("%v", validationErr.Value()),
			Param:     validationErr.Param(),
		})
	}
	return validationErrors
}