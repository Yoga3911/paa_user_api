package routes

import (
	"user_api/configs"
	"user_api/controllers"
	"user_api/repository"
	"user_api/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	DB *pgxpool.Pool = configs.DatabaseConnection()

	jwtS services.JWTS = services.NewJWTS()

	userR repository.UserR  = repository.NewUserR(DB)
	userS services.UserS    = services.NewUserS(DB, userR, jwtS)
	userC controllers.UserC = controllers.NewUserC(userS)

	authR repository.AuthR  = repository.NewAuthR(DB)
	authS services.AuthS    = services.NewAuthS(authR, jwtS)
	authC controllers.AuthC = controllers.NewAuthC(authS)

)

func Route(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/auth/login", authC.Login)
	api.Post("/auth/register", authC.Register)

	api.Get("/user", userC.GetByToken)
	api.Put("/user", userC.UpdateUser)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK!")
	})
}