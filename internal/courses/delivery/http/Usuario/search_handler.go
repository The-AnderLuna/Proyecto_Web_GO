package Usuario

import (
	"Proyecto_Web_GO/internal/services"
	"log"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
)

type SearchUserHandler struct {
	userService *services.UserService
}

func NewSearchUserHandler(r *gin.Engine, userService *services.UserService) {
	handler := &SearchUserHandler{
		userService: userService,
	}

	r.GET("/usuario/search", handler.SearchUser)
}

func (h *SearchUserHandler) SearchUser(c *gin.Context) {
	email := c.Query("email")
	usuario, err := h.userService.GetByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Usuario no encontrado: %v", err)
			c.HTML(http.StatusOK, "Usearch.html", gin.H{
				"Error": "Usuario no encontrado. Intente con otro correo electrónico.",
			})
		} else {
			log.Printf("Error searching user: %v", err)
			c.HTML(http.StatusInternalServerError, "Usearch.html", gin.H{
				"Error": "Error al buscar el usuario. Intente nuevamente.",
			})
		}
		return
	}

	c.HTML(http.StatusOK, "Usearch.html", gin.H{
		"Usuario": usuario,
	})
}
