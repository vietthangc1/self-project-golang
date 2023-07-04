package passwordx

import "golang.org/x/crypto/bcrypt"

type Password struct {
	Password string `json:"password"`
}

func (p *Password) HasingPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (p *Password) CheckPassword(current string) bool {
	inputPassword := []byte(p.Password)
	currentPassword := []byte(current)
	err := bcrypt.CompareHashAndPassword(currentPassword, inputPassword)
	return err == nil
}
