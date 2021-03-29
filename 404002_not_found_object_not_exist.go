package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var NotFoundObjectNotExist = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.NotFound,
	Name:        constant.ErrorNotFound,
	Description: "target object is not exist",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "404002",
	},
})
