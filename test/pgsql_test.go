package test

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	// 使用 gorm 标签来指定字段的属性
	Id       int64  `gorm:"primaryKey"` // 主键
	Username string `gorm:"unique"`     // 唯一键
	Email    string `gorm:"unique"`     // 唯一键
}

func TestPGSql(t *testing.T) {
	dsn := "host=192.168.3.88 user=kyle password=kyle@2023 dbname=mall-boot port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//fmt.Print("TestPGSql", db, err)

	//err = db.AutoMigrate(&User{})
	//// 检查是否有错误
	//if err != nil {
	//	fmt.Println("创建表失败：", err)
	//	return
	//}
	//// 打印成功信息
	//fmt.Println("创建表成功")

	//user := User{
	//	Username: "david",
	//	Email:    "david@example.com",
	//}
	//// 使用 Create 方法来插入数据
	//err = db.Create(&user).Error
	//// 检查是否有错误
	//if err != nil {
	//	fmt.Println("插入数据失败：", err)
	//	return
	//}
	//// 打印成功信息
	//fmt.Println("插入数据成功")

	var user *User
	// 使用 First 方法来查询第一条数据，传入一个条件
	err = db.First(&user, "username = ?", "david").Error
	// 检查是否有错误
	if err != nil {
		fmt.Println("查询数据失败：", err)
		return
	}
	// 打印查询结果
	fmt.Println("查询数据成功：", user)

}
