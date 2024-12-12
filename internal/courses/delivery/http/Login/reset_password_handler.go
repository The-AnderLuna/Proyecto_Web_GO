package Login

import (
	"Proyecto_Web_GO/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResetPasswordHandler struct {
	userService *services.UserService
}

func NewResetPasswordHandler(r *gin.Engine, userService *services.UserService) {
	handler := &ResetPasswordHandler{
		userService: userService,
	}

	r.GET("/login/reset-password", handler.ShowResetPasswordPage)
	r.POST("/reset-password", handler.ResetPassword)
	r.GET("/login/change-password", handler.ShowChangePasswordPage)
	r.POST("/change-password", handler.ChangePassword)
}

func (h *ResetPasswordHandler) ShowResetPasswordPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login/reset_password.html", nil)
}

func (h *ResetPasswordHandler) ResetPassword(c *gin.Context) {
	var resetRequest struct {
		Email string `form:"email" binding:"required"`
	}

	if err := c.ShouldBind(&resetRequest); err != nil {
		c.HTML(http.StatusBadRequest, "login/reset_password.html", gin.H{
			"Error": "Correo electrónico inválido",
		})
		return
	}

	user, err := h.userService.GetByEmail(resetRequest.Email)
	if err != nil {
		c.HTML(http.StatusNotFound, "login/reset_password.html", gin.H{
			"Error": "Correo electrónico no encontrado",
		})
		return
	}

	c.HTML(http.StatusOK, "login/change_password.html", gin.H{
		"Email": user.Email,
	})
}

func (h *ResetPasswordHandler) ShowChangePasswordPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login/change_password.html", nil)
}

func (h *ResetPasswordHandler) ChangePassword(c *gin.Context) {
	var changePasswordRequest struct {
		Email       string `form:"email" binding:"required"`
		NewPassword string `form:"new_password" binding:"required"`
	}

	if err := c.ShouldBind(&changePasswordRequest); err != nil {
		c.HTML(http.StatusBadRequest, "login/change_password.html", gin.H{
			"Error": "Datos inválidos",
		})
		return
	}

	user, err := h.userService.GetByEmail(changePasswordRequest.Email)
	if err != nil {
		c.HTML(http.StatusNotFound, "login/change_password.html", gin.H{
			"Error": "Correo electrónico no encontrado",
		})
		return
	}

	user.Password = changePasswordRequest.NewPassword
	if err := h.userService.UpdateUser(user); err != nil {
		c.HTML(http.StatusInternalServerError, "login/change_password.html", gin.H{
			"Error": "No se pudo cambiar la contraseña",
		})
		return
	}

	c.HTML(http.StatusOK, "login/index.html", gin.H{
		"Message": "Contraseña cambiada exitosamente, ahora puedes iniciar sesión",
	})
}
