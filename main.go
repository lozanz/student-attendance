package main

import (
	"database/sql"
	"fmt"
	"os"
	"student-attendance/controllers"
	"student-attendance/database"
	"student-attendance/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("Success read file environment")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}
	database.DbMigrate(DB)
	defer DB.Close()

	r := gin.Default()
	// Router User
	r.GET("/user", controllers.GetAllUsers)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	// Router Class
	r.GET("/class", middleware.Auth([]string{"admin", "siswa"}), controllers.GetAllClass)
	r.POST("/class", middleware.Auth([]string{"admin"}), controllers.InsertClass)
	r.PUT("/class/:id", middleware.Auth([]string{"admin"}), controllers.UpdateClass)
	r.DELETE("/class/:id", middleware.Auth([]string{"admin"}), controllers.DeleteClass)
	// Router student
	r.GET("/student", middleware.Auth([]string{"admin", "siswa"}), controllers.GetAllStudent)
	r.POST("/student", middleware.Auth([]string{"admin"}), controllers.InsertStudent)
	r.PUT("/student/:id", middleware.Auth([]string{"admin"}), controllers.UpdateStudent)
	r.DELETE("/student/:id", middleware.Auth([]string{"admin"}), controllers.DeleteStudent)
	// Router attendance
	r.GET("/attendance", middleware.Auth([]string{"admin", "siswa"}), controllers.GetAllAttendance)
	r.POST("/attendance", middleware.Auth([]string{"admin"}), controllers.InsertAttendance)
	r.PUT("/attendance/:id", middleware.Auth([]string{"admin"}), controllers.UpdateAttendance)
	r.DELETE("/attendance/:id", middleware.Auth([]string{"admin"}), controllers.DeleteAttendance)

	r.Run(":" + os.Getenv("PORT"))
	// r.Run("localhost:8080")
}
