package Login

import (
	"Proyecto_Web_GO/internal/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginService *services.LoginService
}

func NewLoginHandler(r *gin.Engine, loginService *services.LoginService) {
	handler := &LoginHandler{
		loginService: loginService,
	}

	r.GET("/login", handler.ShowLoginPage)
	r.POST("/login", handler.Login)
}

func (h *LoginHandler) ShowLoginPage(c *gin.Context) {
	// Cargar los archivos HTML individuales
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Error": "", // Explicitamente pasar un error vacío
	})
}

func (h *LoginHandler) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&loginRequest); err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"Error": "Datos de inicio de sesión inválidos",
		})
		return
	}

	user, err := h.loginService.Authenticate(loginRequest.Email, loginRequest.Password)
	if err != nil {
		// Redirigir a la página de inicio de sesión con un mensaje de error
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Error": "Correo o contraseña incorrectos",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("user_nombre", user.Nombre)
	session.Save()

	c.Redirect(http.StatusFound, "/login/welcome")
}
