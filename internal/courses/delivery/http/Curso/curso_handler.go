package curso

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define el handler para mostrar la página de gestión de cursos
func NewCursoHandler(r *gin.Engine) {
	r.GET("/curso", ShowCursoPage)
}

// Muestra la página de gestión de cursos
func ShowCursoPage(c *gin.Context) {
	c.HTML(http.StatusOK, "Curso.html", nil)
}
