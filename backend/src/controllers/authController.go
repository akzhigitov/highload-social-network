package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/akzhigitov/highload-social-network/src/models"
	"github.com/akzhigitov/highload-social-network/src/store"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type RegisterRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Sex       byte   `json:"sex"`
	City      string `json:"city"`
	Email     string `json:"email"`
	Interests string `json:"interests"`
	Password  string `json:"password,omitempty"`
}

func Register(c echo.Context) (err error) {

	request := RegisterRequest{}

	if err = c.Bind(&request); err != nil {
		return err
	}

	if request.Email == "" || request.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid email or password"}
	}

	user := models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Sex:       request.Sex,
		City:      request.City,
		Birthday:  request.Birthday,
		Email:     request.Email,
		Interests: request.Interests,
	}
	user.SetPassword(request.Password)

	if err = store.CreateUser(&user); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c echo.Context) (err error) {
	request := new(LoginRequest)

	if err = c.Bind(request); err != nil {
		log.Error(err)
		return err
	}

	user, err := store.GetUserInfoByEmail(request.Email)
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	if err != nil {
		log.Panic(err)
	}
	log.Info(user)
	if err = user.ComparePassword(request.Password); err != nil {
		log.Info(err)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	user, err = store.GetUser(user.Id)
	if err != nil {
		log.Info(err)
		return err
	}
	Key := "secret"
	user.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		log.Info(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}
