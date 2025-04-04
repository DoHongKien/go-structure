package response

const (
	ErrCodeSuccess = 2001 // Success
	ErrCodeFailed  = 2003 // Failed
)

var msg = map[int]string{
	ErrCodeSuccess: "success",
	ErrCodeFailed:  "failed",
}
