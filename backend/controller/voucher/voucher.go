package controller

import (
	"net/http"

	"github.com/Pcyfer/se-66-stock/entity"
	"github.com/gin-gonic/gin"
)

func CreateVoucher(c *gin.Context) {
	var Voucher entity.Voucher

	// bind เข้าตัวแปร Voucher
	if err := c.ShouldBindJSON(&Voucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//สร้าง Product
	u := entity.Voucher{
		Code:        Voucher.Code,
	}

	//บันทึก
	if err := entity.DB().Create(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": u})

}

func GetVoucher(c *gin.Context) {
	var Vouchers entity.Voucher
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Vouchers WHERE id = ?", id).Find(&Vouchers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Vouchers})
}

func ListVouchers(c *gin.Context) {
	var Vouchers []entity.Voucher
	if err := entity.DB().Raw("SELECT * FROM Vouchers").Find(&Vouchers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Vouchers})
}

func DeleteVoucher(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Vouchers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Voucher not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func UpdateVoucher(c *gin.Context) {
	var Voucher entity.Voucher
	var result entity.Product
	if err := c.ShouldBindJSON(&Voucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//ค้นหา Product ด้วย id
	if tx := entity.DB().Where("id = ?", Voucher.ID).First(&result); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Voucher not found"})
		return
	}
	if err := entity.DB().Save(&Voucher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &Voucher})
}