package entity

import "time"

type Curso struct {
	ID                 uint      `json:"id"`
	Nombre             string    `json:"nombre"`
	Descripcion        string    `json:"descripcion"`
	DuracionHoras      int       `json:"duracion_horas"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}
