package response

const (
	ErrCodeSuccess        = 20001
	ErrCodeParamInvalid   = 20003
	ErrInvalidToken       = 30001
	ErrUserAlreadyExisted = 50001
)

var msg = map[int]string{
	ErrCodeSuccess:        "Success",
	ErrCodeParamInvalid:   "Email is invalid",
	ErrInvalidToken:       "Token is invalid",
	ErrUserAlreadyExisted: "User is already existed",
}
