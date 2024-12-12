package mysql

import (
	"Proyecto_Web_GO/internal/entity"
	"database/sql"
)

type CursoRepository struct {
	DB *sql.DB
}

func NewCursoRepository(db *sql.DB) *CursoRepository {
	return &CursoRepository{DB: db}
}

func (r *CursoRepository) CreateCurso(curso *entity.Curso) error {
	query := "INSERT INTO Cursos (nombre, descripcion, duracion_horas, fecha_creacion, fecha_actualizacion) VALUES (?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, curso.Nombre, curso.Descripcion, curso.DuracionHoras, curso.FechaCreacion, curso.FechaActualizacion)
	return err
}

func (r *CursoRepository) GetCursoByID(id uint) (*entity.Curso, error) {
	var curso entity.Curso
	query := "SELECT id, nombre, descripcion, duracion_horas, fecha_creacion, fecha_actualizacion FROM Cursos WHERE id = ?"
	err := r.DB.QueryRow(query, id).Scan(&curso.ID, &curso.Nombre, &curso.Descripcion, &curso.DuracionHoras, &curso.FechaCreacion, &curso.FechaActualizacion)
	if err != nil {
		return nil, err
	}
	return &curso, nil
}

func (r *CursoRepository) GetAllCursos() ([]*entity.Curso, error) {
	var cursos []*entity.Curso
	query := "SELECT id, nombre, descripcion, duracion_horas, fecha_creacion, fecha_actualizacion FROM Cursos"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var curso entity.Curso
		if err := rows.Scan(&curso.ID, &curso.Nombre, &curso.Descripcion, &curso.DuracionHoras, &curso.FechaCreacion, &curso.FechaActualizacion); err != nil {
			return nil, err
		}
		cursos = append(cursos, &curso)
	}
	return cursos, nil
}

func (r *CursoRepository) UpdateCurso(curso *entity.Curso) error {
	query := "UPDATE Cursos SET nombre = ?, descripcion = ?, duracion_horas = ?, fecha_actualizacion = ? WHERE id = ?"
	_, err := r.DB.Exec(query, curso.Nombre, curso.Descripcion, curso.DuracionHoras, curso.FechaActualizacion, curso.ID)
	return err
}

func (r *CursoRepository) DeleteCurso(id uint) error {
	query := "DELETE FROM Cursos WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}
