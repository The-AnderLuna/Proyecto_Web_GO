package Login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct{}

func NewLogoutHandler(r *gin.Engine) {
	handler := &LogoutHandler{}

	r.GET("/logout", handler.Logout)
}

func (h *LogoutHandler) Logout(c *gin.Context) {
	// Lógica para cerrar sesión, por ejemplo, eliminar la sesión del usuario
	c.Set("user", nil) // Limpiar los datos del usuario en la sesión
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Message": "Has cerrado sesión exitosamente",
	})
}
