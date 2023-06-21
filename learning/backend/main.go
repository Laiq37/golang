package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course --file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"priice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullName"`
	Webssite string `json:"website"`
}

// fake DB(slic/list of courses)
var courses []Course

// middleware/helper --file
func (course *Course) IsEmpty() bool {
	return course.CourseName == ""
}

func main() {
	fmt.Printf("Hello World!")
	courses = append(courses, Course{CourseId: "1", CourseName: "English", CoursePrice: "2000", Author: &Author{
		Fullname: "Syed Laiq Afzal",
		Webssite: "https://flutter.dev/",
	}})

	//creating router
	r := mux.NewRouter()

	//routing

	//GET
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")

	//POST
	r.HandleFunc("/course", createOneCourse).Methods("POST")

	//DELETE
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("POST")

	//UPDATE
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//Controllers -- controllerfolder/file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get on coures")
	w.Header().Set("Content-Type", "application/json")

	//gradb id from params
	params := mux.Vars(r)

	//loop through course, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	//what about -{}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
		}
	}
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Data required for updation")
		return
	}

	var userCourse Course
	_ = json.NewDecoder(r.Body).Decode(&userCourse)
	if userCourse.IsEmpty() {
		json.NewEncoder(w).Encode("Data required for updation")
		return
	}
	for index, course := range courses {
		if course.CourseId == params["id"] {
			userCourse.CourseId = params["id"]
			uCourses := make([]Course, 0)
			uCourses = append(uCourses, userCourse)
			uCourses = append(uCourses, courses[index+1:]...)
			courses = append(courses[:index], uCourses...)
			fmt.Printf("%v", courses)
			json.NewEncoder(w).Encode(userCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("Course doesn't EXIST")
}
