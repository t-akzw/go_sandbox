package domains

import (
	"golang.org/x/crypto/bcrypt"
	"local.packages/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserDomain struct {
	models.User
	Password string `json:"-"`
}

func CreateUser(email string) int {
	data := &models.User{Email: "tahoge@gmail.com"}
	txSave(data)

	return 1
}

func (user *UserDomain) UpdatePassword(password string) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.EncryptedPassword = string(hashed)
}
