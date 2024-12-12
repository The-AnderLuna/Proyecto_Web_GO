package services

import (
	"Proyecto_Web_GO/internal/entity"
	"Proyecto_Web_GO/internal/interfaces"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	repo interfaces.LoginRepository
}

func NewLoginService(repo interfaces.LoginRepository) *LoginService {
	return &LoginService{repo}
}

func (s *LoginService) Authenticate(email, password string) (*entity.Usuario, error) {
	usuario, err := s.repo.Authenticate(email)
	if err != nil {
		return nil, err
	}

	// Verificar la contraseña hasheada
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password))
	if err != nil {
		return nil, errors.New("usuario o contraseña incorrectos")
	}

	return usuario, nil
}
