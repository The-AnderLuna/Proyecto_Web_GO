package main

import (
	"Proyecto_Web_GO/internal/config"
	loginHttp "Proyecto_Web_GO/internal/courses/delivery/http/Login"
	welcomeHttp "Proyecto_Web_GO/internal/courses/delivery/http/Login"
	usuarioHttp "Proyecto_Web_GO/internal/courses/delivery/http/Usuario"
	cursoHttp "Proyecto_Web_GO/internal/courses/delivery/http/curso"
	cursoRepo "Proyecto_Web_GO/internal/courses/repository/mysql"
	loginRepo "Proyecto_Web_GO/internal/courses/repository/mysql"
	userRepo "Proyecto_Web_GO/internal/courses/repository/mysql"
	cursoService "Proyecto_Web_GO/internal/services"
	loginService "Proyecto_Web_GO/internal/services"
	userService "Proyecto_Web_GO/internal/services"
	"Proyecto_Web_GO/pkg/database"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := database.NewMySQLDB(config.GetDBConfig())
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Repositories and Services
	cuRepo := cursoRepo.NewCursoRepository(db)
	cuService := cursoService.NewCursoService(cuRepo)
	uRepo := userRepo.NewUserRepository(db)
	uService := userService.NewUserService(uRepo)
	lRepo := loginRepo.NewLoginRepository(db)
	lService := loginService.NewLoginService(lRepo)

	r := gin.Default()

	// Configurar sesiones
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Cargar los archivos HTML individuales
	r.LoadHTMLFiles(
		"templates/login/index.html",
		"templates/login/register.html",
		"templates/login/welcome.html",
		"templates/usuario/Usuario.html",
		"templates/usuario/Ucreate.html",
		"templates/usuario/Ulist.html",
		"templates/usuario/Usearch.html",
		"templates/usuario/Uupdate.html",
		"templates/curso/Curso.html",
		"templates/curso/CCurso.html",
		"templates/curso/CCreateCurso.html",
		"templates/curso/CListCurso.html",
		"templates/curso/CUpdateCurso.html",
		"templates/curso/CDeleteCurso.html",
	)

	r.Static("/css", "./static/css")
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Ruta raíz para servir directamente el HTML
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Handlers de Cursos
	cursoHttp.NewCreateCursoHandler(r, cuService)
	cursoHttp.NewListCursoHandler(r, cuService)
	cursoHttp.NewUpdateCursoHandler(r, cuService)
	cursoHttp.NewDeleteCursoHandler(r, cuService)
	cursoHttp.NewSearchCursoHandler(r, cuService)
	cursoHttp.NewCursoHandler(r)

	// Handlers de Login
	loginHttp.NewLoginHandler(r, lService)
	loginHttp.NewLogoutHandler(r)
	loginHttp.NewRegisterHandler(r, uService)
	loginHttp.NewResetPasswordHandler(r, uService)

	// Handlers de Usuarios
	usuarioHttp.NewCreateUserHandler(r, uService)
	usuarioHttp.NewListUserHandler(r, uService)
	usuarioHttp.NewUpdateUserHandler(r, uService)
	usuarioHttp.NewSearchUserHandler(r, uService)
	usuarioHttp.NewUsuarioHandler(r)

	// Handler de Bienvenida
	welcomeHttp.NewWelcomeHandler(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
