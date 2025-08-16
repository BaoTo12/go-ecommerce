package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/BaoTo12/go-ecommerce/global"
)

type IUserAuthenticationRepository interface {
	AddOTP(ctx context.Context, email string, otp int, expiration time.Duration) error
}
type UserAuthenticationRepository struct{}

func NewUSerAuthenticationRepository() IUserAuthenticationRepository {
	return &UserAuthenticationRepository{}
}

func (r *UserAuthenticationRepository) AddOTP(ctx context.Context, email string, otp int, expiration time.Duration) error {
	// Nếu muốn một timeout cục bộ nhỏ (ví dụ 500ms) để tránh block quá lâu:
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	key := fmt.Sprintf("usr:%s:otp", email)
	fmt.Println(key)
	return global.Rdb.Set(ctx, key, otp, expiration).Err()
}
