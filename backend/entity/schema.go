package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string `gorm:"uniqueIndex" valid:"มีสิ้นค้านี้อยู่แล้ว,required~กรุณากรอกชื่อใหม่อีกครั้ง"`
	Image          string `gorm:"type:longtext"`
	Price          float64
	Desciption     string
	Softwarekey   []Softwarekey `gorm:"foreignKey:ProductID"`
	CategoryID     *uint
	Category       Category `gorm:"references:id"`
	AdminID        *uint
	Admin          Admin `gorm:"references:id"`
	ManufacturerID *uint
	Manufacturer   Manufacturer `gorm:"references:id"`
}
type Softwarekey struct {
	gorm.Model
	Key       	string 		`gorm:"uniqueIndex" valid:"มีKeyนี้อยู่แล้ว,required~กรุณากรอกใหม่อีกครั้ง"`
	Status    	bool
	ProductID 	*uint
	Product   	Product 	`gorm:"references:id"`
	CartItem	[]CartItem  `gorm:"foreignKey:SoftwareKeyID"`
}

type User struct {
	gorm.Model
	Email           string `gorm:"uniqueIndex" valid:"email~รูปแบบ email ไม่ถูกต้อง,required~กรุณากรอก email"`
	Password        string
	Profile_Picture string `gorm:"type:longtext"`
}
type Admin struct {
	gorm.Model
	Email           string 		 `gorm:"uniqueIndex" valid:"email~รูปแบบ email ไม่ถูกต้อง,required~กรุณากรอก email"`
	Password        string
	Name            string   	 `gorm:"uniqueIndex" valid:"มีUserนี้อยู่แล้ว,required~กรุณากรอกชื่อใหม่อีกครั้ง"`
	Profile_Picture string   	 `gorm:"type:longtext"`
	Product         []Product 	 `gorm:"foreignKey:AdminID"`
}
type Category struct {
	gorm.Model
	Name    string    `gorm:"uniqueIndex" valid:"มีชื่อนี้อยู่แล้ว,required~กรุณากรอกชื่อใหม่อีกครั้ง"`
	Product []Product `gorm:"foreignKey:CategoryID"`
}

type Manufacturer struct {
	gorm.Model
	Name    string    `gorm:"uniqueIndex" valid:"มีชื่อนี้อยู่แล้ว,required~กรุณากรอกชื่อใหม่อีกครั้ง"`
	Product []Product `gorm:"foreignKey:ManufacturerID"`
}

type Cart struct {
	gorm.Model
	Total	float64
	CartItem	[]CartItem `gorm:"foreignKey:CartID"`
	VoucherID	*uint
	Voucher		Voucher		`gorm:"references:id"`
	UserID		*uint
	User		User		`gorm:"references:id"`

}

type CartItem struct {
	gorm.Model
	CartID			*uint
	Cart			Cart			`gorm:"references:id"`
	SoftwarekeyId	*uint
	Softwarekey		Softwarekey		`gorm:"references:id"`
}

type Voucher struct {
	gorm.Model
	Code	string
	Cart	[]Cart		`gorm:"foreignKey:VoucherID"`
}