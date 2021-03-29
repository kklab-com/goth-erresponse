package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var ServerErrorOperationCantBeDone = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.InternalServerError,
	Name:        constant.ErrorServerError,
	Description: "operation can't be done",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Critical,
		ErrorCategory: kkerror.Process,
		ErrorCode:     "500003",
	},
})
