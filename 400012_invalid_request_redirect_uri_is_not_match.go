package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var InvalidRequestRedirectUriIsNotMatch = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.BadRequest,
	Name:        constant.ErrorInvalidRequest,
	Description: "redirect_uri is not match",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "400012",
	},
})
