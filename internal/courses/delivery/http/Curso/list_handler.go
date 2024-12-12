package curso

import (
	"Proyecto_Web_GO/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListCursoHandler struct {
	cursoService *services.CursoService
}

func NewListCursoHandler(r *gin.Engine, cursoService *services.CursoService) {
	handler := &ListCursoHandler{
		cursoService: cursoService,
	}

	r.GET("/curso/list", handler.ListCursos)
}

func (h *ListCursoHandler) ListCursos(c *gin.Context) {
	cursos, err := h.cursoService.GetAllCursos()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "curso/list.html", gin.H{
			"Error": "No se pudieron obtener los cursos",
		})
		return
	}

	c.HTML(http.StatusOK, "curso/list.html", gin.H{
		"Cursos": cursos,
	})
}
