package Usuario

import (
	"Proyecto_Web_GO/internal/entity"
	"Proyecto_Web_GO/internal/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateUserHandler struct {
	userService *services.UserService
}

func NewCreateUserHandler(r *gin.Engine, userService *services.UserService) {
	handler := &CreateUserHandler{
		userService: userService,
	}

	r.GET("/usuario/create", handler.ShowCreateUserPage)
	r.POST("/usuario/create", handler.CreateUser)
}

func (h *CreateUserHandler) ShowCreateUserPage(c *gin.Context) {
	message := c.Query("message")
	errorMessage := c.Query("error")
	c.HTML(http.StatusOK, "Ucreate.html", gin.H{
		"message": message,
		"error":   errorMessage,
	})
}

func (h *CreateUserHandler) CreateUser(c *gin.Context) {
	var request struct {
		Nombre    string `form:"nombre" binding:"required"`
		Apellidos string `form:"apellidos" binding:"required"`
		Email     string `form:"email" binding:"required,email"`
		Telefono  string `form:"telefono" binding:"required"`
		Password  string `form:"password" binding:"required"`
		Rol       string `form:"rol" binding:"required"`
		Estado    string `form:"estado" binding:"required"`
	}

	if err := c.ShouldBind(&request); err != nil {
		log.Printf("Error binding create user request: %v", err)
		c.Redirect(http.StatusFound, "/usuario/create?error=Por favor, complete todos los campos correctamente.")
		return
	}

	usuario := &entity.Usuario{
		Nombre:        request.Nombre,
		Apellidos:     request.Apellidos,
		Email:         request.Email,
		Telefono:      request.Telefono,
		Password:      request.Password,
		Rol:           request.Rol,
		Estado:        request.Estado,
		FechaRegistro: time.Now(),
	}

	if err := h.userService.CreateUser(usuario); err != nil {
		log.Printf("Error creating user: %v", err)
		c.Redirect(http.StatusFound, "/usuario/create?error=Error al crear el usuario. Intente nuevamente.")
		return
	}

	log.Printf("User created successfully: %v", usuario)
	c.Redirect(http.StatusFound, "/usuario/create?message=Usuario creado exitosamente.")
}
