package models

type User struct {
	Model
	Email             string `json:"email"`
	EncryptedPassword string `json:"-"`
}
