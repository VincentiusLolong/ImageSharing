package security

import (
	"mestorage/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(data *models.Account, password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	data.Password = string(bytes)
}

func CheckPassword(data *models.Account, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
