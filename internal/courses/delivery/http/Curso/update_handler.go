package curso

import (
	"Proyecto_Web_GO/internal/entity"
	"Proyecto_Web_GO/internal/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateCursoHandler struct {
	cursoService *services.CursoService
}

func NewUpdateCursoHandler(r *gin.Engine, cursoService *services.CursoService) {
	handler := &UpdateCursoHandler{
		cursoService: cursoService,
	}

	r.GET("/curso/update", handler.ShowUpdateCursoPage)
	r.POST("/curso/update", handler.UpdateCurso)
}

func (h *UpdateCursoHandler) ShowUpdateCursoPage(c *gin.Context) {
	c.HTML(http.StatusOK, "curso/update.html", nil)
}

func (h *UpdateCursoHandler) UpdateCurso(c *gin.Context) {
	var cursoRequest struct {
		ID            uint   `form:"id" binding:"required"`
		Nombre        string `form:"nombre" binding:"required"`
		Descripcion   string `form:"descripcion" binding:"required"`
		DuracionHoras int    `form:"duracion_horas" binding:"required"`
	}

	if err := c.ShouldBind(&cursoRequest); err != nil {
		c.HTML(http.StatusBadRequest, "curso/update.html", gin.H{
			"Error": "Datos del curso inválidos",
		})
		return
	}

	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "curso/update.html", gin.H{
			"Error": "ID inválido",
		})
		return
	}

	curso := &entity.Curso{
		ID:                 uint(id),
		Nombre:             cursoRequest.Nombre,
		Descripcion:        cursoRequest.Descripcion,
		DuracionHoras:      cursoRequest.DuracionHoras,
		FechaActualizacion: time.Now(),
	}

	if err := h.cursoService.UpdateCurso(curso); err != nil {
		c.HTML(http.StatusInternalServerError, "curso/update.html", gin.H{
			"Error": "No se pudo actualizar el curso",
		})
		return
	}

	c.HTML(http.StatusOK, "curso/update.html", gin.H{
		"Message": "Curso actualizado exitosamente",
	})
}
