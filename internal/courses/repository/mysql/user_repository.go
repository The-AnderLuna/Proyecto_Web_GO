package mysql

import (
	"Proyecto_Web_GO/internal/entity"
	"Proyecto_Web_GO/internal/interfaces"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Save(usuario *entity.Usuario) error {
	// Hashear la contraseña antes de guardar
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usuario.Password = string(hashedPassword)

	_, err = r.DB.Exec("INSERT INTO Usuarios (password, nombre, apellidos, rol, email, telefono, estado, fecha_registro) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		usuario.Password, usuario.Nombre, usuario.Apellidos, usuario.Rol, usuario.Email, usuario.Telefono, usuario.Estado, usuario.FechaRegistro)
	return err
}

func (r *UserRepository) GetByEmail(email string) (*entity.Usuario, error) {
	var usuario entity.Usuario
	err := r.DB.QueryRow("SELECT id, password, nombre, apellidos, rol, email, telefono, estado, fecha_registro FROM Usuarios WHERE email = ?", email).
		Scan(&usuario.ID, &usuario.Password, &usuario.Nombre, &usuario.Apellidos, &usuario.Rol, &usuario.Email, &usuario.Telefono, &usuario.Estado, &usuario.FechaRegistro)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UserRepository) GetByID(id uint) (*entity.Usuario, error) {
	var usuario entity.Usuario
	err := r.DB.QueryRow("SELECT id, password, nombre, apellidos, rol, email, telefono, estado, fecha_registro FROM Usuarios WHERE id = ?", id).
		Scan(&usuario.ID, &usuario.Password, &usuario.Nombre, &usuario.Apellidos, &usuario.Rol, &usuario.Email, &usuario.Telefono, &usuario.Estado, &usuario.FechaRegistro)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UserRepository) Update(usuario *entity.Usuario) error {

	if usuario.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		usuario.Password = string(hashedPassword)
	}

	_, err := r.DB.Exec("UPDATE Usuarios SET password = ?, nombre = ?, apellidos = ?, rol = ?, email = ?, telefono = ?, estado = ? WHERE id = ?",
		usuario.Password, usuario.Nombre, usuario.Apellidos, usuario.Rol, usuario.Email, usuario.Telefono, usuario.Estado, usuario.ID)
	return err
}

func (r *UserRepository) Delete(id uint) error {
	_, err := r.DB.Exec("DELETE FROM Usuarios WHERE id = ?", id)
	return err
}

func (r *UserRepository) UpdatePassword(id uint, newPassword string) error {
	// Hashear la nueva contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec("UPDATE Usuarios SET password = ? WHERE id = ?", hashedPassword, id)
	return err
}

func (r *UserRepository) GetAll() ([]*entity.Usuario, error) {
	rows, err := r.DB.Query("SELECT id, nombre, apellidos, rol, email, telefono, estado, fecha_registro FROM Usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []*entity.Usuario
	for rows.Next() {
		var usuario entity.Usuario
		if err := rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Apellidos, &usuario.Rol, &usuario.Email, &usuario.Telefono, &usuario.Estado, &usuario.FechaRegistro); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, &usuario)
	}

	return usuarios, nil
}
