package customize_error

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

var debug bool

type userDefinedError struct {
	err        error
	msg        string
	stackTrace *[]string // err stack_trace_info
}

func New(err error, msg string) error {
	return &userDefinedError{err: err, msg: msg, stackTrace: formatErrorStack(err)}
}

func (u *userDefinedError) Error() string {
	msg := u.msg + "\n" + u.err.Error()
	if debug {
		if u.stackTrace != nil {
			msg += "\n" + strings.Join(*u.stackTrace, "\n")
		}
	}
	return msg
}

func formatErrorStack(err error) *[]string {
	str := fmt.Sprintf("%+v", errors.WithStack(err))
	str = strings.ReplaceAll(str, "\n\t", "\n")
	stack := strings.Split(str, "\n")
	stack = stack[5:]
	return &stack
}

// SetDebug enable/disable debug.
func SetDebug(enable bool) {
	debug = enable
}
