package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var cruds []Crud

type Crud struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"Email"`
	Password string `json:"Password"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/crud", listUser)
	e.POST("/crud", createUser)
	e.GET("/crud/:id", getUser)
	e.POST("/crud/login", crudLogin)
	e.PUT("/crud/:id", putUser)
	e.DELETE("/crud/:id", deleteUser)
	e.Logger.Fatal(e.Start(":8081"))
}

func listUser(c echo.Context) error {
	return c.JSON(http.StatusOK, cruds)
}

func createUser(c echo.Context) error {
	crud := Crud{}
	err := c.Bind(&crud)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	cruds = append(cruds, crud)
	return c.JSON(http.StatusOK, cruds)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, crud := range cruds {
		if crud.ID == id {
			return c.JSON(http.StatusOK, crud)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func putUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	crud := Crud{}
	err := c.Bind(&crud)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	for i, _ := range cruds {
		if cruds[i].ID == id {
			cruds[i].Name = crud.Name
			return c.JSON(http.StatusOK, cruds)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range cruds {
		if cruds[i].ID == id {
			cruds = remove(cruds, i)
			return c.JSON(http.StatusOK, cruds)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func remove(s []Crud, i int) []Crud {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func crudLogin(c echo.Context) error {
	crud := Crud{}
	err := c.Bind(&crud)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	for _, crudData := range cruds {
		if crudData.Name == crud.Name {
			if crudData.Password == crud.Password {
				return c.JSON(http.StatusOK, true)
			}
		}
	}

	return c.JSON(http.StatusOK, false)
}