package main

import (
	"Go/student"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func initDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&student.Student{})
	return db
}

func main() {
	db := initDB()
	defer db.Close()

	studentAPI := initStudentAPI(db)
	r := gin.Default()
	r.GET("/students", studentAPI.FindAll)
	r.GET("/students/:id", studentAPI.FindByID)
	r.POST("/students", studentAPI.Create)
	r.PUT("/students/:id", studentAPI.Update)
	r.DELETE("/students/:id", studentAPI.Delete)

	err := r.Run(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}