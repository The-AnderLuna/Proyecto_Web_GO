package Login

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func NewWelcomeHandler(r *gin.Engine) {
	r.GET("/login/welcome", ShowWelcomePage)
}

func ShowWelcomePage(c *gin.Context) {
	session := sessions.Default(c)
	userNombre := session.Get("user_nombre")

	if userNombre == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"Nombre": userNombre,
	})
}
