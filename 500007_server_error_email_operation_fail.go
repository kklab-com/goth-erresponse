package erresponse

import (
	"fmt"

	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var ServerErrorEmailOperationFail = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.InternalServerError,
	Name:        constant.ErrorServerError,
	Description: "email operation fail",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Urgent,
		ErrorCategory: kkerror.Internal,
		ErrorCode:     "500007",
	},
})

func ServerErrorEmailOperationWithMessage(format string, params ...interface{}) ErrorResponse {
	return &DefaultErrorResponse{
		StatusCode:  httpstatus.InternalServerError,
		Name:        constant.ErrorServerError,
		Description: "email operation fail",
		DefaultKKError: kkerror.DefaultKKError{
			ErrorLevel:    kkerror.Urgent,
			ErrorCategory: kkerror.Internal,
			ErrorCode:     "500007",
			ErrorMessage:  fmt.Sprintf(format, params...),
		},
	}
}
