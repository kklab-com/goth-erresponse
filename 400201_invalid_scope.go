package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var InvalidScope = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.BadRequest,
	Name:        constant.ErrorInvalidScope,
	Description: "",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "400201",
	},
})
