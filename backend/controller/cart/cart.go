package controller

import (
	"net/http"

	"github.com/Pcyfer/se-66-stock/entity"
	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	var Cart entity.Cart
	var User entity.User
	var Voucher entity.Voucher

	// bind เข้าตัวแปร Cart
	if err := c.ShouldBindJSON(&Cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Cart.UserID).First(&User); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", Cart.VoucherID).First(&Voucher); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Voucher not found"})
		return
	}

	//สร้าง Cart
	u := entity.Cart{
		User:    	User,
		Voucher: 	Voucher,
		Total:      Cart.Total,
	}

	//บันทึก
	if err := entity.DB().Create(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": u})

}

func GetCart(c *gin.Context) {
	var Carts entity.Product
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Carts WHERE id = ?", id).Find(&Carts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Carts})
}

func DeleteCart(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Carts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func UpdateCart(c *gin.Context) {
	var Cart entity.Cart
	var result entity.Cart
	if err := c.ShouldBindJSON(&Cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//ค้นหา Cart ด้วย id
	if tx := entity.DB().Where("id = ?", Cart.ID).First(&result); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
		return
	}
	if err := entity.DB().Save(&Cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &Cart})
}