package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var InvalidGrantScopeNotGranted = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.Forbidden,
	Name:        constant.ErrorInsufficientScope,
	Description: "scope not granted",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "403101",
	},
})
