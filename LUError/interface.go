package LUError

type Error interface {
	Error() string
	LUErrorCode() int
	HTTP(code int) Error
	HTTPCode() int
	SetDescription(description string) Error
	GetDescription() string
}
