package response

const (
	ErrCodeSuccess = 2001 // Success
	ErrCodeParamInvalid = 2003 // Invalid parameter
)

var msg = map[int]string{
	ErrCodeSuccess: "success",
	ErrCodeParamInvalid: "invalid parameter",
}