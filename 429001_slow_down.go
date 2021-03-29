package erresponse

import (
	"github.com/kklab-com/gone-httpstatus"
	"github.com/kklab-com/goth-erresponse/constant"
	"github.com/kklab-com/goth-kkerror"
)

var SlowDownTooFast = Collection.Register(&DefaultErrorResponse{
	StatusCode:  httpstatus.TooManyRequests,
	Name:        constant.ErrorSlowDown,
	Description: "too fast, rate limit exceeded",
	DefaultKKError: kkerror.DefaultKKError{
		ErrorLevel:    kkerror.Normal,
		ErrorCategory: kkerror.Client,
		ErrorCode:     "429001",
	},
})
