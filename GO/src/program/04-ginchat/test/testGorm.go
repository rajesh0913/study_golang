package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:qinzhibao@tcp(127.0.0.1:3306)/ginChat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// db.AutoMigrate(&models.UserBasic{})
	// db.AutoMigrate(&models.Message{})
	// db.AutoMigrate(&models.GroupBasic{})
	// db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.Community{})

	// user := &models.UserBasic{
	// 	Name: "qin",
	// }

	// Create
	// db.Create(user)

	// Read
	// var product Product
	// fmt.Println(db.First(user, 1)) // find product with integer primary key-----
	// db.First(user, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	// db.Model(user).Update("Password", "2345")----
	// Update - update multiple fields
	// db.Model(user).Updates(models.UserBasic {Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}
