package po

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string
	RoleNote string `gorm:"type:text"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
