package erresponse

import (
	"fmt"

	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var ServerErrorSMSOperationFail = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.InternalServerError,
	Name:        constant.ErrorServerError,
	Description: "sms operation fail",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Critical,
		ErrorCategory: kkerror.Internal,
		ErrorCode:     "500008",
	},
})

func ServerErrorSMSOperationWithMessage(format string, params ...interface{}) ErrorResponse {
	return &DefaultErrorResponse{
		StatusCode:  httpstatus.InternalServerError,
		Name:        constant.ErrorServerError,
		Description: "sms operation fail",
		DefaultKKError: kkerror.DefaultKKError{
			ErrorLevel:    kkerror.Critical,
			ErrorCategory: kkerror.Internal,
			ErrorCode:     "500008",
			ErrorMessage:  fmt.Sprintf(format, params...),
		},
	}
}
