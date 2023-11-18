package LUError

type Error struct {
	ErrorCode   int
	Message     string
	HttpCode    int
	Description string
}

func (l *Error) New(code int, message string) *Error {
	l.ErrorCode = code
	l.Message = message
	return l
}

func (l *Error) HTTP(code int) {
	l.HttpCode = code
}

func (l *Error) Error() string {
	return l.Message
}

func (l *Error) HTTPCode() int {
	return l.HttpCode
}

func (l *Error) LUErrorCode() int {
	return l.ErrorCode
}

func (l *Error) SetDescription(description string) *Error {
	l.Description = description
	return l
}

func (l *Error) GetDescription() string {
	return l.Description
}
