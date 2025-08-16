package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/BaoTo12/go-ecommerce/internal/repo"
	"github.com/BaoTo12/go-ecommerce/pkg/response"
	"github.com/BaoTo12/go-ecommerce/pkg/utils/crypto"
	"github.com/BaoTo12/go-ecommerce/pkg/utils/random"
	sendmail "github.com/BaoTo12/go-ecommerce/pkg/utils/sendMail"
)

type IUserService interface {
	Register(ctx context.Context, email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthenticationRepository
}

func NewUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthenticationRepository) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(ctx context.Context, email string, purpose string) int {

	// 0. Hash emails
	hashedEmail := crypto.HashEmail(email)
	fmt.Printf("Hashed Email::%s \n", hashedEmail)
	// 1. check whether email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrUserAlreadyExisted
	}
	// 2. Send new OTP
	otp := random.SixDigitsOTPGenerator()
	if purpose == "TEST" {
		otp = 123456
	}
	fmt.Printf("OTP is ::: %d \n", otp)
	// 3. Save OTP in redis
	err := us.userAuthRepo.AddOTP(ctx, hashedEmail, otp, 10*time.Minute)
	if err != nil {
		return response.ErrInvalidOTP
	}
	// 4. Send Email OTP
	errEmail := sendmail.SendTextMailOTP([]string{email}, "annoysitck@gmail.com", strconv.Itoa(otp))
	if errEmail != nil {
		return response.ErrSendEmailOTP
	}
	// 5. Check if OTP is available

	// 6. Prevent user spam ...

	return response.ErrCodeSuccess
}
