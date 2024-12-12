package curso

import (
	"Proyecto_Web_GO/internal/entity"
	"Proyecto_Web_GO/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateCursoHandler struct {
	cursoService *services.CursoService
}

func NewCreateCursoHandler(r *gin.Engine, cursoService *services.CursoService) {
	handler := &CreateCursoHandler{
		cursoService: cursoService,
	}

	r.GET("/curso/create", handler.ShowCreateCursoPage)
	r.POST("/curso/create", handler.CreateCurso)
}

func (h *CreateCursoHandler) ShowCreateCursoPage(c *gin.Context) {
	c.HTML(http.StatusOK, "curso/create.html", nil)
}

func (h *CreateCursoHandler) CreateCurso(c *gin.Context) {
	var cursoRequest struct {
		Nombre        string `form:"nombre" binding:"required"`
		Descripcion   string `form:"descripcion" binding:"required"`
		DuracionHoras int    `form:"duracion_horas" binding:"required"`
	}

	if err := c.ShouldBind(&cursoRequest); err != nil {
		c.HTML(http.StatusBadRequest, "curso/create.html", gin.H{
			"Error": "Datos del curso inválidos",
		})
		return
	}

	curso := &entity.Curso{
		Nombre:             cursoRequest.Nombre,
		Descripcion:        cursoRequest.Descripcion,
		DuracionHoras:      cursoRequest.DuracionHoras,
		FechaCreacion:      time.Now(),
		FechaActualizacion: time.Now(),
	}

	if err := h.cursoService.CreateCurso(curso); err != nil {
		c.HTML(http.StatusInternalServerError, "curso/create.html", gin.H{
			"Error": "No se pudo crear el curso",
		})
		return
	}

	c.HTML(http.StatusOK, "curso/create.html", gin.H{
		"Message": "Curso creado exitosamente",
	})
}
