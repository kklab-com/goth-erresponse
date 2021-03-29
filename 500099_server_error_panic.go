package erresponse

import (
	"fmt"

	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var ServerErrorPanic = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.InternalServerError,
	Name:        constant.ErrorServerError,
	Description: "panic",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Urgent,
		ErrorCategory: kkerror.Server,
		ErrorCode:     "500099",
	},
})

func ServerErrorPanicWithMessage(format string, params ...interface{}) ErrorResponse {
	return &DefaultErrorResponse{
		StatusCode:  httpstatus.InternalServerError,
		Name:        constant.ErrorServerError,
		Description: "panic",
		DefaultKKError: kkerror.DefaultKKError{
			ErrorLevel:    kkerror.Urgent,
			ErrorCategory: kkerror.Server,
			ErrorCode:     "500099",
			ErrorMessage:  fmt.Sprintf(format, params...),
		},
	}
}
