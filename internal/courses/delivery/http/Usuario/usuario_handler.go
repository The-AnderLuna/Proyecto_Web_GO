package Usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUsuarioHandler(r *gin.Engine) {
	r.GET("/usuario", ShowUsuarioPage)
}

func ShowUsuarioPage(c *gin.Context) {
	c.HTML(http.StatusOK, "Usuario.html", nil)
}
