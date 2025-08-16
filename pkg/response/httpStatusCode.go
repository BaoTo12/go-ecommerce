package response

const (
	ErrCodeSuccess        = 20001
	ErrCodeParamInvalid   = 20003
	ErrInvalidToken       = 30001
	ErrInvalidOTP         = 30002
	ErrSendEmailOTP       = 30003
	ErrUserAlreadyExisted = 50001
)

var msg = map[int]string{
	ErrCodeSuccess:        "Success",
	ErrCodeParamInvalid:   "Email is invalid",
	ErrInvalidToken:       "Token is invalid",
	ErrInvalidOTP:         "OTP is invalid",
	ErrSendEmailOTP:       "Send Email OTP failed",
	ErrUserAlreadyExisted: "User is already existed",
}
