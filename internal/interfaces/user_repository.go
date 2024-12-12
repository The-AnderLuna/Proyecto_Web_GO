package interfaces

import (
	"Proyecto_Web_GO/internal/entity"
)

type UserRepository interface {
	Save(usuario *entity.Usuario) error
	GetByEmail(email string) (*entity.Usuario, error)
	GetByID(id uint) (*entity.Usuario, error)
	Update(usuario *entity.Usuario) error
	Delete(id uint) error
	GetAll() ([]*entity.Usuario, error)
	UpdatePassword(id uint, newPassword string) error
}
