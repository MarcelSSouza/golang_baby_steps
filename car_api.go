package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
type Car struct {
	gorm.Model
	Make  string
	Modelo string
	Year  int
}
func main() {
	// Initialize DB connection
	dsn := "username:password@tcp(127.0.0.1:3306)/car_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db

	// Migrate the schema
	DB.AutoMigrate(&Car{})

	// Set up Gin
	r := gin.Default()

	// Define routes
	r.POST("/cars", CreateCar)
	r.GET("/cars/:id", GetCar)
	r.GET("/cars", GetCars)
	r.PUT("/cars/:id", UpdateCar)
	r.DELETE("/cars/:id", DeleteCar)

	// Run the server
	r.Run()
}


func CreateCar(c *gin.Context) {
	var input Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	car := Car{Make: input.Make, Model: input.Model, Year: input.Year}
	DB.Create(&car)
	c.JSON(200, car)
}
func GetCar(c *gin.Context) {
	id := c.Param("id")
	var car Car
	if result := DB.First(&car, id); result.Error != nil {
		c.JSON(404, gin.H{"error": "Car not found"})
		return
	}
	c.JSON(200, car)
}
func GetCars(c *gin.Context) {
	var cars []Car
	DB.Find(&cars)
	c.JSON(200, cars)
}
func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var car Car
	if result := DB.First(&car, id); result.Error != nil {
		c.JSON(404, gin.H{"error": "Car not found"})
		return
	}
	var input Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	DB.Model(&car).Updates(input)
	c.JSON(200, car)
}
func DeleteCar(c *gin.Context) {
	id := c.Param("id")
	var car Car
	if result := DB.First(&car, id); result.Error != nil {
		c.JSON(404, gin.H{"error": "Car not found"})
		return
	}
	DB.Delete(&car)
	c.JSON(200, gin.H{"message": "Car deleted"})
}

