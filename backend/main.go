package main

import (
	"github.com/gin-gonic/gin"

	admin_controller "github.com/Pcyfer/se-66-stock/controller/admin"
	category_controller "github.com/Pcyfer/se-66-stock/controller/category"
	login_controller "github.com/Pcyfer/se-66-stock/controller/login"
	Product_controller "github.com/Pcyfer/se-66-stock/controller/product"
	Softwarekey_controller "github.com/Pcyfer/se-66-stock/controller/softwarekey"
	user_controller "github.com/Pcyfer/se-66-stock/controller/user"
	cart_controller "github.com/Pcyfer/se-66-stock/controller/cart"
	cartitem_controller "github.com/Pcyfer/se-66-stock/controller/cartitem"
	voucher_controller "github.com/Pcyfer/se-66-stock/controller/voucher"
	"github.com/Pcyfer/se-66-stock/entity"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	// login User Route
	r.POST("/login/user", login_controller.LoginUser)
	r.POST("/users", user_controller.CreateUser)

	// login Admin Route
	r.POST("/login/admin", login_controller.LoginAdmin)
	r.POST("/admin", admin_controller.CreateAdmin)

	// Admin Routes
	// r.GET("/productkey", Product_controller.ListProductsJoinKey)
	//Products
	r.GET("/Products", Product_controller.ListProducts)
	r.GET("/search", Product_controller.SearchProducts)
	r.GET("/Products/:id", Product_controller.GetProduct)
	r.POST("/Products", Product_controller.CreateProduct)
	r.PATCH("/Products", Product_controller.UpdateProduct)
	r.DELETE("/Products/:id", Product_controller.DeleteProduct)
	//Softwarekey
	r.GET("/key", Softwarekey_controller.ListSoftwarekeys)
	r.GET("/key/:id", Softwarekey_controller.GetSoftwarekey)
	r.POST("/key", Softwarekey_controller.CreateSoftwarekey)
	r.PATCH("/key", Softwarekey_controller.UpdateSoftwarekey)
	r.DELETE("/key/:id", Softwarekey_controller.DeleteSoftwarekey)

	//Category
	r.POST("/category", category_controller.CreateCategory)
	r.GET("/category", category_controller.ListCategory)
	r.GET("/category/:id", category_controller.GetCategory)
	//Manufacturer
	r.POST("/manufacturer", Product_controller.CreateManufacturer)
	r.GET("/manufacturer", Product_controller.ListManufacturer)
	r.GET("/manufacturer/:id", Product_controller.GetManufacturer)

	//Cart
	r.POST("/cart", cart_controller.CreateCart)
	r.DELETE("/cart/:id", cart_controller.DeleteCart)
	r.GET("/cart/:id", cart_controller.GetCart)
	r.PATCH("cart", cart_controller.UpdateCart)

	//CartItem
	r.POST("/cartitem", cartitem_controller.CreateCartItem)
	r.GET("/cartitem/:id", cartitem_controller.GetCartItem)
	r.GET("/cartitem", cartitem_controller.ListCartItems)
	r.DELETE("/cartitem/:id", cartitem_controller.DeleteCartItem)
	r.PATCH("/cartitem/:id", cartitem_controller.UpdateCartItem)

	//Voucher
	r.POST("/voucher", voucher_controller.CreateVoucher)
	r.GET("/voucher/:id", voucher_controller.GetVoucher)
	r.GET("/voucher", voucher_controller.ListVouchers)
	r.DELETE("/voucher/:id", voucher_controller.DeleteVoucher)
	r.PATCH("/voucher/:id", voucher_controller.UpdateVoucher)
	r.Run()

	


}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,PATCH,DELETE")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
