package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var courses []Course

type Course struct {
	ID   string
	Name string
}

func generateCourses() {
	course1 := Course{
		ID:   "1",
		Name: "Full Cycle",
	}

	course2 := Course{
		ID:   "2",
		Name: "Bonus Full Cycle",
	}

	courses = append(courses, course1, course2)
}

func main() {
	generateCourses()
	http.HandleFunc("/courses", listCourses)
	http.ListenAndServe(":8081", nil)
}

func listCourses(w http.ResponseWriter, r *http.Request) {
	jsonCourses, _ := json.Marshal(courses)
	// persistCourse()
	w.Write([]byte(jsonCourses))
}

func persistCourse() error {
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("insert into courses values($1, $2)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec("abc", "Benf")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
