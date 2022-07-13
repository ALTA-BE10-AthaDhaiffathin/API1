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

func FindUser(id int) (User, int) {
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			return users[i], i
		}
	}
	return User{}, 0
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

	findUser, _ := FindUser(cnv)

	if findUser.Name != "" {
		res := map[string]interface{}{
			"message": "Get user " + param,
			"data":    findUser,
		}
		return c.JSON(http.StatusOK, res)
	}
	return c.JSON(http.StatusInternalServerError, "Cant find user " + param)
}

func removeUser(slice []User, index int) []User {
	newSlice := []User{}
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
	param := c.Param("id")
	cnv, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Cannot convert to int", err.Error())
		return c.JSON(http.StatusInternalServerError, "cannot convert id")
	}

	findUser, index := FindUser(cnv)

	if findUser.Name != "" {
		users = removeUser(users, index)
		res := map[string]interface{}{
			"message": "Success delete user " + param,
			"data":    findUser,
		}
		return c.JSON(http.StatusOK, res)
	}
	return c.JSON(http.StatusInternalServerError, "Cant find user " + param)
}

// update user by id
func UpdateUserController(c echo.Context) error {
  	// your solution here
	param := c.Param("id")
	cnv, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Cannot convert to int", err.Error())
		return c.JSON(http.StatusInternalServerError, "cannot convert id")
	}

	findUser, index := FindUser(cnv)

	if findUser.Name != "" {
		user := User{}
		user.Id = findUser.Id
		err := c.Bind(&user)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Wrong input")
		}

		// if user.Name == "" {
		// 	user.Name = findUser.Name
		// }
		// if user.Email == "" {
		// 	user.Name = findUser.Email
		// }
		// if user.Password == "" {
		// 	user.Password = findUser.Password
		// }

		users[index] = user
		res := map[string]interface{}{
			"message": "Update user " + param,
			"data":  users[index],
		}
		return c.JSON(http.StatusOK, res)
	}
	return c.JSON(http.StatusInternalServerError, "Cant find user " + param)
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