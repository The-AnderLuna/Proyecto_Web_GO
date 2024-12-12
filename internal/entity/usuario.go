package entity

import (
	"time"
)

type Usuario struct {
	ID            uint      `json:"id"`
	Password      string    `json:"password"`
	Nombre        string    `json:"nombre"`
	Apellidos     string    `json:"apellidos"`
	Rol           string    `json:"rol"`
	Email         string    `json:"email"`
	Telefono      string    `json:"telefono"`
	Estado        string    `json:"estado"`
	FechaRegistro time.Time `json:"fecha_registro"`
}
