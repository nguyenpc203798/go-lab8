package main

import (
	"golang-book-management/controllers"
	"golang-book-management/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "host=dpg-cse55um8ii6s738t0heg-a.singapore-postgres.render.com user=nguyen_postgresql_user password=CWQ70J8X8YcXztEiHaULXJKSdwd5pYu0 dbname=nguyen_postgresql port=5432 sslmode=require TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	db.AutoMigrate(&models.Book{})
	return db
}

func main() {
	db := initDB()
	router := gin.Default()

	// Set up the routes for Book management
	bookController := controllers.NewBookController(db)
	router.GET("/books", bookController.GetBooks)
	router.POST("/books", bookController.CreateBook)
	router.GET("/books/:id", bookController.GetBookByID)
	router.PUT("/books/:id", bookController.UpdateBook)
	router.DELETE("/books/:id", bookController.DeleteBook)

	router.Run(":8080")
}
