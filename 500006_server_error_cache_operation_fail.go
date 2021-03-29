package erresponse

import (
	"fmt"

	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var ServerErrorCacheOperationFail = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.InternalServerError,
	Name:        constant.ErrorServerError,
	Description: "cache operation fail",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Critical,
		ErrorCategory: kkerror.Cache,
		ErrorCode:     "500006",
	},
})

func ServerErrorCacheOperationWithMessage(format string, params ...interface{}) ErrorResponse {
	return &DefaultErrorResponse{
		StatusCode:  httpstatus.InternalServerError,
		Name:        constant.ErrorServerError,
		Description: "cache operation fail",
		DefaultKKError: kkerror.DefaultKKError{
			ErrorLevel:    kkerror.Critical,
			ErrorCategory: kkerror.Cache,
			ErrorCode:     "500006",
			ErrorMessage:  fmt.Sprintf(format, params...),
		},
	}
}
