package controller

import (
	"net/http"

	"github.com/Pcyfer/se-66-stock/entity"
	"github.com/gin-gonic/gin"
)


func CreateCartItem(c *gin.Context) {
	var CartItem entity.CartItem
	var Cart entity.Cart
	var Softwarekey entity.Softwarekey

	// bind เข้าตัวแปร Product
	if err := c.ShouldBindJSON(&CartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", CartItem.CartID).First(&Cart); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", CartItem.SoftwarekeyId).First(&Softwarekey); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	//สร้าง Product
	u := entity.CartItem{
		Cart:        	 Cart,
		Softwarekey:     Softwarekey,
		
	}

	//บันทึก
	if err := entity.DB().Create(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": u})

}

func GetCartItem(c *gin.Context) {
	var CartItems entity.Product
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM CartItems WHERE id = ?", id).Find(&CartItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": CartItems})
}

func ListCartItems(c *gin.Context) {
	var CartItems []entity.CartItem
	if err := entity.DB().Raw("SELECT * FROM CartItems").Find(&CartItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": CartItems})
}

func DeleteCartItem(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM CartItems WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CartItems not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func UpdateCartItem(c *gin.Context) {
	var CartItem entity.CartItem
	var result entity.Product
	if err := c.ShouldBindJSON(&CartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//ค้นหา Product ด้วย id
	if tx := entity.DB().Where("id = ?", CartItem.ID).First(&result); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CartItem not found"})
		return
	}
	if err := entity.DB().Save(&CartItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &CartItem})
}