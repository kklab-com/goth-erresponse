package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var InvalidToken = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.Unauthorized,
	Name:        constant.ErrorInvalidToken,
	Description: "",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "401101",
	},
})
