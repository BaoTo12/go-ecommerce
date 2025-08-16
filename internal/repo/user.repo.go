package repo

import (
	"github.com/BaoTo12/go-ecommerce/global"
	"github.com/BaoTo12/go-ecommerce/internal/model"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

func (*userRepository) GetUserByEmail(email string) bool {
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	return row != NumberNull
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
