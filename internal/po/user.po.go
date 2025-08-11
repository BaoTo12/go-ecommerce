package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid; type:varchar(255); not null; index:idx_uuidl unique"`
	UserName string    `gorm:"type:varchar(255); unique"`
	IsActive bool      `gorm:"type:boolean"`
	Roles    []Role    `gorm:"many2many:go_user_roles"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
