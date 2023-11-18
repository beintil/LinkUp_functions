package LUError

type LError struct {
	errorCode   int
	message     string
	httpCode    int
	description string
}

func New(code int, message string) Error {
	e := &LError{errorCode: code, message: message}
	return e
}

func (l *LError) HTTP(code int) {
	l.httpCode = code
}

func (l *LError) Error() string {
	return l.message
}

func (l *LError) HTTPCode() int {
	return l.httpCode
}

func (l *LError) LUErrorCode() int {
	return l.errorCode
}

func (l *LError) SetDescription(description string) Error {
	l.description = description
	return l
}

func (l *LError) GetDescription() string {
	return l.description
}
