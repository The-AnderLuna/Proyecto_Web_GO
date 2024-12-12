package interfaces

import "Proyecto_Web_GO/internal/entity"

type CursoRepository interface {
	CreateCurso(curso *entity.Curso) error
	GetCursoByID(id uint) (*entity.Curso, error)
	GetAllCursos() ([]*entity.Curso, error)
	UpdateCurso(curso *entity.Curso) error
	DeleteCurso(id uint) error
}
