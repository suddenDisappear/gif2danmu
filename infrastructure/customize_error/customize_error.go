package customize_error

var debug bool

type userDefinedError struct {
	err error
	msg string
}

func New(err error, msg string) error {
	return &userDefinedError{err: err, msg: msg}
}

func (u *userDefinedError) Error() string {
	if debug {
		return u.msg + "\n" + u.err.Error()
	}
	return u.msg
}

func init() {
	debug = true
}
