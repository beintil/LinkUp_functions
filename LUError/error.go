package LUError

type LUError struct {
	ErrorCode   int
	Message     string
	HttpCode    int
	Description string
}

func (l *LUError) New(code int, message string) *LUError {
	l.ErrorCode = code
	l.Message = message
	return l
}

func (l *LUError) HTTP(code int) {
	l.HttpCode = code
}

func (l *LUError) Error() string {
	return l.Message
}

func (l *LUError) HTTPCode() int {
	return l.HttpCode
}

func (l *LUError) LUErrorCode() int {
	return l.ErrorCode
}

func (l *LUError) SetDescription(description string) *LUError {
	l.Description = description
	return l
}

func (l *LUError) GetDescription() string {
	return l.Description
}
