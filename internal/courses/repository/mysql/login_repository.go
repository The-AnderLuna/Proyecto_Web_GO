package mysql

import (
	"Proyecto_Web_GO/internal/entity"
	"Proyecto_Web_GO/internal/interfaces"
	"database/sql"
	"errors"
)

type LoginRepository struct {
	DB *sql.DB
}

func NewLoginRepository(db *sql.DB) interfaces.LoginRepository {
	return &LoginRepository{
		DB: db,
	}
}

func (r *LoginRepository) Authenticate(email string) (*entity.Usuario, error) {
	var usuario entity.Usuario
	err := r.DB.QueryRow("SELECT id, password, nombre, apellidos, rol, email, telefono, estado, fecha_registro FROM Usuarios WHERE email = ?", email).
		Scan(&usuario.ID, &usuario.Password, &usuario.Nombre, &usuario.Apellidos, &usuario.Rol, &usuario.Email, &usuario.Telefono, &usuario.Estado, &usuario.FechaRegistro)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("usuario o contraseña incorrectos")
		}
		return nil, err
	}

	return &usuario, nil
}
