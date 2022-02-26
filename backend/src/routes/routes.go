package routes

import (
	"github.com/akzhigitov/highload-social-network/src/controllers"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(e *echo.Echo) {

	r := e.Group("/restricted")

	config := middleware.JWTConfig{
		Claims:     &jwt.MapClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	e.POST("/auth/sign-in", controllers.Login)
	e.POST("/auth/sign-up", controllers.Register)

	e.GET("/users/", controllers.GetUsers)

	e.GET("/users/:id", controllers.GetUser)
	e.PUT("/users/:id", controllers.UpdateUser)

	r.POST("/friends/:id", controllers.Follow)
	r.DELETE("/friends/:id", controllers.UnFollow)
}
