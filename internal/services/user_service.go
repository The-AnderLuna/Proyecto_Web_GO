package services

import (
	"Proyecto_Web_GO/internal/entity"
	"Proyecto_Web_GO/internal/interfaces"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(usuario *entity.Usuario) error {
	return s.repo.Save(usuario)
}

func (s *UserService) GetByEmail(email string) (*entity.Usuario, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) GetUserByID(id uint) (*entity.Usuario, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) UpdateUser(usuario *entity.Usuario) error {
	return s.repo.Update(usuario)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *UserService) GetAllUsers() ([]*entity.Usuario, error) {
	return s.repo.GetAll()
}

func (s *UserService) UpdatePassword(id uint, newPassword string) error {
	return s.repo.UpdatePassword(id, newPassword)
}
