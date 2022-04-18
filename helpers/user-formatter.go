package helpers

import (
	"github.com/aliepba/go-crypto/app/models"
)

type UserFormatter struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user models.User, token string) UserFormatter {
	formatter := UserFormatter{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return formatter
}
