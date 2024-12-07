package consts

const (
	Success = 0
	ErrnoUnknownError = 1
	
	// param error 1xxx
	ErrnoBindRequestError = 1000
	ErrnoRequestValidateError = 1001

	// mysql error 2xxx

)


var ErrMsg = map[int]string{
	Success: "success",
	ErrnoUnknownError: "unknown error",
	ErrnoBindRequestError: "bind request error",
	ErrnoRequestValidateError: "request validate error",
}