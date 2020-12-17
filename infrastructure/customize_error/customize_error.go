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
	stackTrace []string // err stack_trace_info
}

func New(err error, msg string) error {
	return &userDefinedError{err: err, msg: msg, stackTrace: formatErrorStack(err)}
}

func (u *userDefinedError) Error() string {
	s := []string{u.msg}
	if u.err != nil {
		s = append(s, u.err.Error())
	}
	if debug {
		s = append(s, u.stackTrace...)
	}
	return strings.Join(s, "\n")
}

func formatErrorStack(err error) []string {
	if err == nil {
		return nil
	}
	str := fmt.Sprintf("%+v", errors.WithStack(err))
	stack := strings.Split(str, "\n")
	return stack[5:]
}

// SetDebug enable/disable debug.
func SetDebug(enable bool) {
	debug = enable
}
