package interfaces

import "Proyecto_Web_GO/internal/entity"

type LoginRepository interface {
	Authenticate(email string) (*entity.Usuario, error)
}
