package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var NotFoundCredentialNotFound = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.NotFound,
	Name:        constant.ErrorNotFound,
	Description: "credential not found",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Urgent,
		ErrorCategory: kkerror.Server,
		ErrorCode:     "404004",
	},
})
