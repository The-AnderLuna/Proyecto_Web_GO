package Usuario

import (
	"Proyecto_Web_GO/internal/entity"
	"Proyecto_Web_GO/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserHandler struct {
	userService *services.UserService
}

func NewUpdateUserHandler(r *gin.Engine, userService *services.UserService) {
	handler := &UpdateUserHandler{
		userService: userService,
	}

	r.GET("/usuario/find", handler.FindUser)
	r.POST("/usuario/update", handler.UpdateUser)
}

func (h *UpdateUserHandler) FindUser(c *gin.Context) {
	email := c.Query("email")
	usuario, err := h.userService.GetByEmail(email)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		c.HTML(http.StatusOK, "Uupdate.html", gin.H{
			"Error": "Usuario no encontrado. Intente nuevamente.",
		})
		return
	}

	c.HTML(http.StatusOK, "Uupdate.html", gin.H{
		"Usuario": usuario,
	})
}

func (h *UpdateUserHandler) UpdateUser(c *gin.Context) {
	var request struct {
		ID        int    `form:"id" binding:"required"`
		Nombre    string `form:"nombre" binding:"required"`
		Apellidos string `form:"apellidos" binding:"required"`
		Email     string `form:"email" binding:"required,email"`
		Telefono  string `form:"telefono" binding:"required"`
		Rol       string `form:"rol" binding:"required"`
		Estado    string `form:"estado" binding:"required"`
	}

	if err := c.ShouldBind(&request); err != nil {
		log.Printf("Error binding update user request: %v", err)
		c.HTML(http.StatusBadRequest, "Uupdate.html", gin.H{
			"Error": "Por favor, complete todos los campos correctamente.",
		})
		return
	}

	usuario := &entity.Usuario{

		Nombre:    request.Nombre,
		Apellidos: request.Apellidos,
		Email:     request.Email,
		Telefono:  request.Telefono,
		Rol:       request.Rol,
		Estado:    request.Estado,
	}

	if err := h.userService.UpdateUser(usuario); err != nil {
		log.Printf("Error updating user: %v", err)
		c.HTML(http.StatusInternalServerError, "Uupdate.html", gin.H{
			"Error": "Error al actualizar el usuario. Intente nuevamente.",
		})
		return
	}

	log.Printf("User updated successfully: %v", usuario)
	c.Redirect(http.StatusFound, "/usuario/list")
}
