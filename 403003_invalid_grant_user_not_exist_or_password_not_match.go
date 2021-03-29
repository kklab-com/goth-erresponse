package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var InvalidGrantUserNotExistOrPasswordNotMatch = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.Forbidden,
	Name:        constant.ErrorInvalidGrant,
	Description: "user is not exist or password is not correct",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "403003",
	},
})
