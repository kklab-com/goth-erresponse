package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var InvalidClientExpiresAtIsTooLong = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.Unauthorized,
	Name:        constant.ErrorInvalidClient,
	Description: "client secret exp is too long",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "401002",
	},
})
