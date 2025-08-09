package repo

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// user repo --> ur
func (ur *UserRepository) GetUserInformation() string {
	return "Chi Bao Learning"
}
