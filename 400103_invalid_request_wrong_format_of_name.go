package erresponse

import (
	"fmt"

	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var _InvalidRequestWrongFormatOfName = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.BadRequest,
	Name:        constant.ErrorInvalidRequest,
	Description: "parameter wrong format",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "400103",
		ErrorMessage:  fmt.Sprintf("%s has wrong format", "[name]"),
	},
})

func InvalidRequestWrongFormatOfName(name string) ErrorResponse {
	return &DefaultErrorResponse{
		StatusCode:  httpstatus.BadRequest,
		Name:        constant.ErrorInvalidRequest,
		Description: "parameter wrong format",
		DefaultKKError: kkerror.DefaultKKError{
			ErrorLevel:    kkerror.Normal,
			ErrorCategory: kkerror.Client,
			ErrorCode:     "400103",
			ErrorMessage:  fmt.Sprintf("%s has wrong format", name),
		},
	}
}
