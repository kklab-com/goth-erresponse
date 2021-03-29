package erresponse

import (
	"fmt"

	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var ServerErrorCrossServiceOperationFail = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.InternalServerError,
	Name:        constant.ErrorServerError,
	Description: "cross service operation fail",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Critical,
		ErrorCategory: kkerror.Internal,
		ErrorCode:     "500009",
	},
})

func ServerErrorCrossServiceOperationWithMessage(format string, params ...interface{}) ErrorResponse {
	return &DefaultErrorResponse{
		StatusCode:  httpstatus.InternalServerError,
		Name:        constant.ErrorServerError,
		Description: "cross service operation fail",
		DefaultKKError: kkerror.DefaultKKError{
			ErrorLevel:    kkerror.Critical,
			ErrorCategory: kkerror.Internal,
			ErrorCode:     "500000",
			ErrorMessage:  fmt.Sprintf(format, params...),
		},
	}
}
