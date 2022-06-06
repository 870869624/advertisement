package errors

type Error struct {
	Code        int
	Information string
}

func New(Information string) *Error {
	instance := &Error{
		Information: Information,
	}
	return instance
}

func (e Error) Error() string {
	return e.Information
}
