package erresponse

import (
	"fmt"

	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var _InvalidRequestWithMessage = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.BadRequest,
	Name:        constant.ErrorInvalidRequest,
	Description: "",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "400102",
	},
})

func InvalidRequestWithMessage(format string, params ...interface{}) ErrorResponse {
	return &DefaultErrorResponse{
		StatusCode:  httpstatus.BadRequest,
		Name:        constant.ErrorInvalidRequest,
		Description: "",
		DefaultKKError: kkerror.DefaultKKError{
			ErrorLevel:    kkerror.Normal,
			ErrorCategory: kkerror.Client,
			ErrorCode:     "400102",
			ErrorMessage:  fmt.Sprintf(format, params...),
		},
	}
}
