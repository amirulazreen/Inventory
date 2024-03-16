package internal

import (
	"math/rand"
	"net/http"
	"strconv"

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
        if err := c.BindJSON(&newitem); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        
        if newitem.Item == "" || newitem.Price <= 0 || newitem.Quantity <= 0 || newitem.Supplier.Name == "" || newitem.Supplier.Address == "" || newitem.Supplier.Tel == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "All fields must be filled"})
            return
        }

        var allItems []Product
        if err := db.Find(&allItems).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to fetch items from the database"})
            return
        }

        if len(newitem.Item) > 50 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Item name should have at most 50 characters"})
            return
        }
        for _, existingItem := range allItems {
            if existingItem.Item == newitem.Item {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Item name should be unique"})
                return
            }
        }

        if newitem.Price > 100000 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Price should not be more than 6 digits"})
            return
        }

        if newitem.Quantity > 1000000 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "item quantity should not be more than 1,000,000"})
            return
        }

        if err := db.Create(&newitem).Error; err != nil {
            c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to add item"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Item added"})
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

        if item.Item == "" || item.Price <= 0 || item.Quantity <= 0 || item.Supplier.Name == "" || item.Supplier.Address == "" || item.Supplier.Tel == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "All fields must be filled"})
            return
        }

        var allItems []Product
        if err := db.Find(&allItems).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to fetch items from the database"})
            return
        }

        if len(item.Item) > 50 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Item name should have at most 50 characters"})
            return
        }
        for _, existingItem := range allItems {
            if existingItem.Item == item.Item {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Item name should be unique"})
                return
            }
        }

        if item.Price > 100000 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Price should not be more than 6 digits"})
            return
        }

        if item.Quantity > 1000000 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "item quantity should not be more than 1,000,000"})
            return
        }

        if err := db.Model(&item).Updates(&item).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update item"})
            return
        }
        c.JSON(http.StatusOK, item)
    }
}

func Addmany(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        for i := 0; i < 1000; i++ {
            newItem := Random(i)

            if err := db.Create(&newItem).Error; err != nil {
                c.JSON(http.StatusBadGateway, gin.H{
                    "error": "Failed to add item",
                })
                return
            }
        }

        c.JSON(http.StatusOK, gin.H{"message": "Items added"})
    }
}

func Random(i int) Product {
    return Product{
        Item:     "Item" + strconv.Itoa(i),
        Price:    float32(rand.Intn(10000)) / 100, 
        Quantity: rand.Intn(100),       
        Supplier: Supplier{
            Name:    "Supplier" + strconv.Itoa(i),
            Address: "23 Taman Bakti, Taman Perkasa Sakti, 2300 Kuala Lumpur",
            Tel:     "020 020 2020",
        },
    }
}




