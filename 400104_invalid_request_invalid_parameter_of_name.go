package erresponse

import (
	"fmt"

	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var _InvalidRequestInvalidParameterOfName = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.BadRequest,
	Name:        constant.ErrorInvalidRequest,
	Description: "invalid parameter",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "400104",
		ErrorMessage:  fmt.Sprintf("%s has invalid data", "[name]"),
	},
})

func InvalidRequestInvalidDataOfName(name string) ErrorResponse {
	return &DefaultErrorResponse{
		StatusCode:  httpstatus.BadRequest,
		Name:        constant.ErrorInvalidRequest,
		Description: "invalid parameter",
		DefaultKKError: kkerror.DefaultKKError{
			ErrorLevel:    kkerror.Normal,
			ErrorCategory: kkerror.Client,
			ErrorCode:     "400104",
			ErrorMessage:  fmt.Sprintf("%s has invalid data", name),
		},
	}
}
