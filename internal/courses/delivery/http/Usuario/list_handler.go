package Usuario

import (
	"Proyecto_Web_GO/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListUserHandler struct {
	userService *services.UserService
}

func NewListUserHandler(r *gin.Engine, userService *services.UserService) {
	handler := &ListUserHandler{
		userService: userService,
	}

	r.GET("/usuario/list", handler.ListUsers)
}

func (h *ListUserHandler) ListUsers(c *gin.Context) {
	usuarios, err := h.userService.GetAllUsers()
	if err != nil {
		log.Printf("Error listing users: %v", err)
		c.HTML(http.StatusInternalServerError, "Ulist.html", gin.H{
			"Error": "Error al listar los usuarios. Intente nuevamente.",
		})
		return
	}

	c.HTML(http.StatusOK, "Ulist.html", gin.H{
		"Usuarios": usuarios,
	})
}
