package vo

type UserRegistrationRequest struct {
	Email   string `json:"email" binding:"required"`
	Purpose string `json:"purpose" binding:"required"`
}
