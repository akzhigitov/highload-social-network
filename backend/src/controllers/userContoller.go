package controllers

import (
	"net/http"
	"strconv"

	"github.com/akzhigitov/highload-social-network/src/models"
	"github.com/akzhigitov/highload-social-network/src/store"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func UpdateUser(c echo.Context) (err error) {

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	u := new(models.User)
	err = c.Bind(&u)
	if err != nil {
		return err
	}

	err = store.UpdateUser(userId, u)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) (err error) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user, err := store.GetUser(userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) (err error) {
	user, err := store.GetUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func Follow(c echo.Context) (err error) {

	userID1 := userIDFromToken(c)

	userId2, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	store.Follow(userID1, userId2)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

func UnFollow(c echo.Context) (err error) {

	userID1 := userIDFromToken(c)

	userId2, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	store.UnFollow(userID1, userId2)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

func userIDFromToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := *user.Claims.(*jwt.MapClaims)
	return int(claims["id"].(float64))
}
