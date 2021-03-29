package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var NotFoundUserNotFound = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.NotFound,
	Name:        constant.ErrorNotFound,
	Description: "user not found",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "404003",
	},
})
