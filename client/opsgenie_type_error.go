package client

/*const (
	RequestNotProcessed = "RequestNotProcessed"
	Unauthorized        = "401"
	NotFound            = "404"
)

var errorMappings = map[string]func(format string, args ...interface{}) error{
	RequestNotProcessed: NewRequestNotProcessedErrorf,
	Unauthorized:        NewUnauthorizedErrorf,
}

type ogError interface {
	error
	HeaderCode() string
	StatusCode() string
}

type requestNotProcessedError struct {
	message    string
	headerCode string//todo error header
	statusCode string
}

func NewRequestNotProcessedErrorf(format string, args ...interface{}) error {
	return &requestNotProcessedError{
		message:    fmt.Sprintf(format, args...),
		headerCode: RequestNotProcessed,
		statusCode: NotFound,
	}
}

func (e *requestNotProcessedError) Error() string {
	return e.message
}

func (e *requestNotProcessedError) HeaderCode() string {
	return RequestNotProcessed
}

func (e *requestNotProcessedError) StatusCode() string {
	return NotFound
}

type unauthorizedError struct {
	message    string
	headerCode string//remove
	statusCode string
}

func NewUnauthorizedErrorf(format string, args ...interface{}) error {
	return &unauthorizedError{
		message:    fmt.Sprintf(format, args...),
		headerCode: "",
		statusCode: Unauthorized,
	}
}

func (e *unauthorizedError) Error() string {
	return e.message
}

func (e *unauthorizedError) HeaderCode() string {
	return ""
}

func (e *unauthorizedError) StatusCode() string {
	return Unauthorized
}*/
