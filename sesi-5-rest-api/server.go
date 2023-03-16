package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)
type User struct {
	Id      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
}

var users []User

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) < 1 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, v := range users {
		if v.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success user found",
				"user":     v,
			})
		}
	}
	return c.JSON(echo.ErrBadRequest.Code, map[string]interface{}{
		"messages": "User Not Found",
		"user":     nil,
	})
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	input := User{}
	c.Bind(&input)
	for i, v := range users {
		if v.Id == id {
			v.Name = input.Name
			v.Email = input.Email
			v.Address = input.Address
			users[i] = v
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success updated user",
				"user":     v,
			})
		}
	}
	return c.JSON(echo.ErrBadRequest.Code, map[string]interface{}{
		"messages": "User Not Found",
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, v := range users {
		if v.Id == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "delete success",
			})
		}
	}
	return c.JSON(echo.ErrBadRequest.Code, map[string]interface{}{
		"messages": "User Not Found",
	})
}
func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)          //get all users
	e.GET("/users/:id", GetUserController)       //get user by id
	e.POST("/users", CreateUserController)       //Create new user
	e.PUT("/users/:id", UpdateUserController)    //update user
	e.DELETE("/users/:id", DeleteUserController) //delete user
	e.Logger.Fatal(e.Start(":8000"))
}