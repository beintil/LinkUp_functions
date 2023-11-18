package LUError

type Error interface {
	New(code int, message string) Error
	Error() string
	LUErrorCode() int
	HTTP(code int)
	HTTPCode() int
}
