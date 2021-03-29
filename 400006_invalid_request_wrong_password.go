package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var InvalidRequestWrongPassword = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.BadRequest,
	Name:        constant.ErrorInvalidRequest,
	Description: "wrong password",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "400006",
	},
})
