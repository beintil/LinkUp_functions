package LUError

type LError struct {
	ErrorCode   int
	Message     string
	HttpCode    int
	Description string
}

func (l *LError) New(code int, message string) Error {
	l.ErrorCode = code
	l.Message = message
	return l
}

func (l *LError) HTTP(code int) {
	l.HttpCode = code
}

func (l *LError) Error() string {
	return l.Message
}

func (l *LError) HTTPCode() int {
	return l.HttpCode
}

func (l *LError) LUErrorCode() int {
	return l.ErrorCode
}

func (l *LError) SetDescription(description string) Error {
	l.Description = description
	return l
}

func (l *LError) GetDescription() string {
	return l.Description
}
