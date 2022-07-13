package main

import (
	"strconv"
	"log"
	"net/http"
  	"github.com/labstack/echo"
)

type User struct {
  	Id    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
  	// your solution here
	param := c.Param("id")
	cnv, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Cannot convert to int", err.Error())
		return c.JSON(http.StatusInternalServerError, "cannot convert id")
	}

	if cnv == 0 || cnv > len(users) {
		log.Println("Index out of range")
		return c.JSON(http.StatusInternalServerError, "Index out of range")
	} 
  
	res := map[string]interface{}{
		"message": "Get user " + param,
		"data":    users[cnv-1],
	}
	return c.JSON(http.StatusOK, res)
}

func removeUser(slice []User, index int) []User {
	newSlice := []User{}
	i := 0
	for i < len(users) {
		if users[i].Id == index {
			index = i
			break
		}
		if i == len(users) - 1 && users[i].Id != index {
			return nil
		}
		i++
	}
	if index == 0 {
		newSlice = slice[1:]
	} else if index == len(slice)-1 {
		newSlice = slice[:len(slice)-1]
	} else {
		newSlice = append(newSlice, slice[:index]...)
		newSlice = append(newSlice, slice[index+1:]...)
	}
	return newSlice
}

// delete user by id
func DeleteUserController(c echo.Context) error {
  	// your solution here
	param := c.Param("id")
	cnv, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Cannot convert to int", err.Error())
		return c.JSON(http.StatusInternalServerError, "cannot convert id")
	}
  
	if cnv == 0 || cnv > len(users) {
		log.Println("Index out of range")
		return c.JSON(http.StatusInternalServerError, "Index out of range")
	} 
	
	res := map[string]interface{}{
		"message": "Delete user " + param,
		"data":    users[cnv-1],
	}

	if removeUser(users, cnv-1) == nil {
		return c.JSON(http.StatusInternalServerError, "Cant find user " + param)
	}

	users = removeUser(users, cnv-1)
	return c.JSON(http.StatusOK, res)
}
// update user by id
func UpdateUserController(c echo.Context) error {
  	// your solution here
	
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
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
// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}