package methods

import (
	"github.com/aliepba/go-crypto/app/models"
	"gorm.io/gorm"
)

type MethodUser interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindById(ID int) (models.User, error)
	Update(user models.User) (models.User, error)
}

type method struct {
	db *gorm.DB
}

func NewMethodUser(db *gorm.DB) *method {
	return &method{db}
}

func (m *method) Save(user models.User) (models.User, error) {
	err := m.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *method) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := m.db.Where("email= ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m *method) FindById(ID int) (models.User, error) {
	var user models.User
	err := m.db.Where("id= ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *method) Update(user models.User) (models.User, error) {
	err := m.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
