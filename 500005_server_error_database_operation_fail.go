package erresponse

import (
	"fmt"

	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var ServerErrorDatabaseOperationFail = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.InternalServerError,
	Name:        constant.ErrorServerError,
	Description: "database operation fail",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Critical,
		ErrorCategory: kkerror.Database,
		ErrorCode:     "500005",
	},
})

func ServerErrorDatabaseOperationWithMessage(format string, params ...interface{}) ErrorResponse {
	return &DefaultErrorResponse{
		StatusCode:  httpstatus.InternalServerError,
		Name:        constant.ErrorServerError,
		Description: "database operation fail",
		DefaultKKError: kkerror.DefaultKKError{
			ErrorLevel:    kkerror.Critical,
			ErrorCategory: kkerror.Database,
			ErrorCode:     "500005",
			ErrorMessage:  fmt.Sprintf(format, params...),
		},
	}
}
