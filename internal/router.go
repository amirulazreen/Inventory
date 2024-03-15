package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOne(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var item Product
        if err := db.Preload("Supplier").First(&item, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"Error": "Record not found"})
            return
        }
        c.JSON(http.StatusOK, item)
    }
}

func GetAll(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
		var items []Product
		if err := db.Preload("Supplier").Find(&items).Error; err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"Error" : "Cannot Fetch Data"})
		}
		c.JSON(http.StatusOK,items)
	}
}

func Add(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var newitem Product
        if err := c.BindJSON(&newitem);
        err != nil{
            c.JSON(http.StatusBadRequest, err)
            return
        }
        if err := db.Create(&newitem).Error;
        err != nil {
            c.JSON(http.StatusBadGateway, gin.H{
                "Error" : "Failed to add item",
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{"Message" : "Item added"})
    }
}

func Deletee(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var item Product

        if err := db.Preload("Supplier").Find(&item,id).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "Error": "Item does not exist",
            })
            return
        }

        if err := db.Delete(&item).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "Error": "Failed to delete item",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Item deleted",
        })
    }
}

func UpdateItem(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")

        var item Product
        if err := db.Preload("Supplier").First(&item, id).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"Error": "Item not found"})
            return
        }

        if err := c.BindJSON(&item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid data format"})
            return
        }

        if err := db.Model(&item).Updates(&item).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update item"})
            return
        }

        c.JSON(http.StatusOK, item)
    }
}


