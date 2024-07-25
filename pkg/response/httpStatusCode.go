package response

const (
	ErrCodeSuccess = 2001
	ErrCodeParamInvalid = 2003
	ErrInvalidToken = 2004
)

var msg = map[int]string{
	ErrCodeSuccess: "success",
	ErrCodeParamInvalid: "email is invalid",
	ErrInvalidToken: "invalid token",
}