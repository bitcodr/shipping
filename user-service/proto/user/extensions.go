package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/kevinburke/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("id", uuid.String())
}
