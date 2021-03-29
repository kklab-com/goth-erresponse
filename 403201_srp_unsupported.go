package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var SRPUnsupported = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.Forbidden,
	Name:        constant.ErrorSRPUnsupported,
	Description: "",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "403201",
	},
})
