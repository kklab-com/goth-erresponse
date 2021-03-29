package erresponse

import (
	"fmt"

	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var _NotFoundWithMessage = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.BadRequest,
	Name:        constant.ErrorNotFound,
	Description: "",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "404101",
	},
})

func NotFoundWithMessage(format string, params ...interface{}) ErrorResponse {
	return &DefaultErrorResponse{
		StatusCode:  httpstatus.BadRequest,
		Name:        constant.ErrorNotFound,
		Description: "",
		DefaultKKError: kkerror.DefaultKKError{
			ErrorLevel:    kkerror.Normal,
			ErrorCategory: kkerror.Client,
			ErrorCode:     "404101",
			ErrorMessage:  fmt.Sprintf(format, params...),
		},
	}
}
