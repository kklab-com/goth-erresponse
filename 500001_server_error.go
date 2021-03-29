package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var ServerError = Collection.Register(&DefaultErrorResponse{
	StatusCode: httpstatus.InternalServerError,
	Name:       constant.ErrorServerError,
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Urgent,
		ErrorCategory: kkerror.Server,
		ErrorCode:     "500001",
	},
})
