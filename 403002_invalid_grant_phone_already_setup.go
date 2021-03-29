package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var InvalidGrantPhoneAlreadySetup = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.Forbidden,
	Name:        constant.ErrorInvalidGrant,
	Description: "phone already setup",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "403002",
	},
})
