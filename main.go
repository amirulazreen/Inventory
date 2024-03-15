package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inv/internal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("Inventory.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&internal.Product{})
    
    router := gin.Default()

    router.Use(internal.CorsMiddleware())

    router.GET("/inventory/:id", internal.GetOne(db))
    router.GET("/inventory", internal.GetAll(db))
    router.POST("/add-inventory", internal.Add(db))
    router.DELETE("/delete-inventory/:id", internal.Deletee(db))
    router.PUT("/update-inventory/:id", internal.UpdateItem(db))

    router.Run(":8080")
}