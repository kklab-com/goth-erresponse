package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-kkerror"
)

var MethodNotAllowed = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.MethodNotAllowed,
	Name:        "method_not_allowed",
	Description: "",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "405001",
	},
})
