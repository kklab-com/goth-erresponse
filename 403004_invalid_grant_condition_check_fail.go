package erresponse

import (
	httpstatus "github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	kkerror "github.com/kklab-com/goth-kkerror"
)

var InvalidGrantConditionCheckFail = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.Forbidden,
	Name:        constant.ErrorInvalidGrant,
	Description: "condition check fail",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "403004",
	},
})
