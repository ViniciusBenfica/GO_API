package main

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var courses []Course

type Course struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func generateCourses() {
	course1 := Course{
		ID:   "1",
		Name: "Vinicius",
	}

	course2 := Course{
		ID:   "2",
		Name: "Benfica",
	}

	courses = append(courses, course1, course2)
}

func main() {
	generateCourses()
	e := echo.New()
	e.GET("/course", listCourses)
	e.POST("/course", createCourse)
	e.Logger.Fatal(e.Start(":8081"))
}

func listCourses(c echo.Context) error {
	return c.JSON(http.StatusOK, courses)
}

func createCourse(c echo.Context) error {
	course := Course{}
	c.Bind(&course)
	courses = append(courses, course)
	return c.JSON(http.StatusOK, courses)
}
