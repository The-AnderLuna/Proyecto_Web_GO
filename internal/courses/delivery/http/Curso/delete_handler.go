package curso

import (
	"Proyecto_Web_GO/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCursoHandler struct {
	cursoService *services.CursoService
}

func NewDeleteCursoHandler(r *gin.Engine, cursoService *services.CursoService) {
	handler := &DeleteCursoHandler{
		cursoService: cursoService,
	}

	r.GET("/curso/delete", handler.ShowDeleteCursoPage)
	r.POST("/curso/delete", handler.DeleteCurso)
}

func (h *DeleteCursoHandler) ShowDeleteCursoPage(c *gin.Context) {
	c.HTML(http.StatusOK, "curso/delete.html", nil)
}

func (h *DeleteCursoHandler) DeleteCurso(c *gin.Context) {
	idStr := c.PostForm("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "curso/delete.html", gin.H{
			"Error": "ID inválido",
		})
		return
	}

	if err := h.cursoService.DeleteCurso(uint(id)); err != nil {
		c.HTML(http.StatusInternalServerError, "curso/delete.html", gin.H{
			"Error": "No se pudo eliminar el curso",
		})
		return
	}

	c.HTML(http.StatusOK, "curso/delete.html", gin.H{
		"Message": "Curso eliminado exitosamente",
	})
}
