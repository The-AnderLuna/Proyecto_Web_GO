package Login

import (
	"Proyecto_Web_GO/internal/entity"
	"Proyecto_Web_GO/internal/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	userService *services.UserService
}

func NewRegisterHandler(r *gin.Engine, userService *services.UserService) {
	handler := &RegisterHandler{
		userService: userService,
	}

	r.GET("/login/register", handler.ShowRegisterPage)
	r.POST("/register", handler.Register)
}

func (h *RegisterHandler) ShowRegisterPage(c *gin.Context) {
	// Cargar los archivos HTML individuales
	c.HTML(http.StatusOK, "register.html", nil)
}

func (h *RegisterHandler) Register(c *gin.Context) {
	var registerRequest struct {
		Nombre    string `form:"nombre" binding:"required"`
		Apellidos string `form:"apellidos" binding:"required"`
		Rol       string `form:"rol" binding:"required"`
		Email     string `form:"email" binding:"required"`
		Password  string `form:"password" binding:"required"`
		Telefono  string `form:"telefono" binding:"required"`
	}

	if err := c.ShouldBind(&registerRequest); err != nil {
		log.Printf("Error binding register request: %v", err)
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"Error": "Datos de registro inválidos",
		})
		return
	}

	usuario := &entity.Usuario{
		Nombre:        registerRequest.Nombre,
		Apellidos:     registerRequest.Apellidos,
		Rol:           registerRequest.Rol,
		Email:         registerRequest.Email,
		Password:      registerRequest.Password,
		Telefono:      registerRequest.Telefono,
		Estado:        "Activo",
		FechaRegistro: time.Now(),
	}

	if err := h.userService.CreateUser(usuario); err != nil {
		log.Printf("Error creating user: %v", err)
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"Error": "No se pudo completar el registro",
		})
		return
	}

	log.Printf("User registered successfully: %v", usuario)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Message": "Registro exitoso, ahora puedes iniciar sesión",
	})
}
