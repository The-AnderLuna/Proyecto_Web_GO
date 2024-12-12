package services

import (
	"Proyecto_Web_GO/internal/entity"
	"Proyecto_Web_GO/internal/interfaces"
)

type CursoService struct {
	repo interfaces.CursoRepository
}

func NewCursoService(repo interfaces.CursoRepository) *CursoService {
	return &CursoService{repo: repo}
}

func (s *CursoService) CreateCurso(curso *entity.Curso) error {
	return s.repo.CreateCurso(curso)
}

func (s *CursoService) GetCursoByID(id uint) (*entity.Curso, error) {
	return s.repo.GetCursoByID(id)
}

func (s *CursoService) GetAllCursos() ([]*entity.Curso, error) {
	return s.repo.GetAllCursos()
}

func (s *CursoService) UpdateCurso(curso *entity.Curso) error {
	return s.repo.UpdateCurso(curso)
}

func (s *CursoService) DeleteCurso(id uint) error {
	return s.repo.DeleteCurso(id)
}
