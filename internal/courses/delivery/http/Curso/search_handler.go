package curso

import (
	"Proyecto_Web_GO/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchCursoHandler struct {
	cursoService *services.CursoService
}

func NewSearchCursoHandler(r *gin.Engine, cursoService *services.CursoService) {
	handler := &SearchCursoHandler{
		cursoService: cursoService,
	}

	r.GET("/curso/search", handler.ShowSearchCursoPage)
	r.POST("/curso/search", handler.SearchCurso)
}

func (h *SearchCursoHandler) ShowSearchCursoPage(c *gin.Context) {
	c.HTML(http.StatusOK, "curso/search.html", nil)
}

func (h *SearchCursoHandler) SearchCurso(c *gin.Context) {
	var searchRequest struct {
		ID uint `form:"id"`
	}

	if err := c.ShouldBind(&searchRequest); err != nil {
		c.HTML(http.StatusBadRequest, "curso/search.html", gin.H{
			"Error": "Datos de búsqueda inválidos",
		})
		return
	}

	curso, err := h.cursoService.GetCursoByID(searchRequest.ID)
	if err != nil {
		c.HTML(http.StatusNotFound, "curso/search.html", gin.H{
			"Error": "Curso no encontrado",
		})
		return
	}

	c.HTML(http.StatusOK, "curso/search.html", gin.H{
		"Curso": curso,
	})
}
