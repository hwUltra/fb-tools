package test

import (
	"fast-boot/app/rpc/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGormBuilder(t *testing.T) {
	dsn := "root:123456@tcp(192.168.3.88:3306)/boot?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("db = ", db)
	m := model.PmsGoodsModel{}
	db.Preload("Shop", "status = 1").
		Preload("Category", "visible = 1").
		Preload("Brand").
		Preload("Skus").
		Preload("SpecList", "type = 1").
		Preload("AttrList", "type = 2").
		Debug().First(&m, 1)
	fmt.Println("goods = ", m)
	fmt.Println("Shop = ", m.Shop)
	fmt.Println("Category = ", m.Category)
	fmt.Println("Brand = ", m.Brand)
	fmt.Println("Skus = ", m.SkuList)
	fmt.Println("SpecList = ", m.SpecList)
	fmt.Println("AttrList = ", m.AttrList)

	/**
		Shop        PmsShopModel     `gorm:"foreignKey:ShopId"`
	Category    PmsCategoryModel `gorm:"foreignKey:CategoryId"`
	Brand       PmsBrandModel    `gorm:"foreignKey:BrandId"`
	Skus        []PmsSkuModel    `gorm:"foreignKey:GoodsID"`
	*/

	//db.First(&user, 10)

	//roleIds := make([]string, 0)
	//db.Debug().Model(&model.SysRoleUsersModel{}).Where("`user_id` = ?", 1).Pluck("role_id", &roleIds)
	//fmt.Print("roleIds = ", roleIds)

	//roleIds := ""
	//db.Debug().Model(&model.SysRoleUsersModel{}).Where("`user_id` = ?", 1).Limit(1).Pluck("role_id", &roleIds)
	//fmt.Print("roleIds = ", roleIds)
	//
	//roleIds := make([]int64, 0)
	//db.Model(&model.SysRoleUsersModel{}).
	//	Joins("JOIN sys_roles on sys_roles.id = sys_role_users.role_id").
	//	Where("`sys_role_users`.user_id = ?", 1).
	//	Where("`sys_roles`.status = ?", 1).
	//	Pluck("`sys_roles`.id", &roleIds)
	//fmt.Print("roleIds = ", roleIds)
	//
	//menuIds := make([]int64, 0)
	//db.Model(&model.SysRoleMenusModel{}).
	//	Joins("JOIN sys_menu on sys_menu.id = sys_role_menu.menu_id").
	//	Where("`sys_role_menu`.role_id IN ?", roleIds).
	//	Pluck("`sys_role_menu`.menu_id", &menuIds)
	////
	//fmt.Println("menuIds = ", menuIds)
	//
	//items := make([]*model.SysMenusModel, 0)
	//db.Model(&model.SysMenusModel{}).
	//	Where("id IN ?", menuIds).
	//	Order("id asc").
	//	Preload("Roles", "status = ?", 1).
	//	Find(&items)
	//
	//fmt.Println("items = ", items)
	//
	//if err := db.Transaction(func(tx *gorm.DB) error {
	//	tx.Delete(&model.SysUserModel{}, 8)
	//	return nil
	//}); err != nil {
	//	fmt.Println("Transaction err  = ", err)
	//}

	//menuIds := make([]int64, 0)
	//db.Debug().Model(&model.SysRoleMenusModel{}).
	//	Joins("JOIN sys_menu on sys_menu.id = sys_role_menu.menu_id").
	//	Where("`sys_role_menu`.role_id IN ?", roleIds).
	//	Pluck("`sys_role_menu`.menu_id", &menuIds)

	//items := make([]model.SysRoleUsersModel, 0)
	//db.Model(&model.SysRoleUsersModel{}).Where("`user_id` = ?", 1).
	//	Preload("Role").Find(&items)
	//
	//for _, item := range items {
	//
	//	fmt.Println("Role = ", item.Role)
	//}

	//userInfo := model.SysUserModel{
	//	Name:          "testName",
	//	Avatar:        "testAvatar",
	//	Username:      "testName",
	//	RememberToken: "RememberToken",
	//	Password:      "Password",
	//}
	//db.Debug().Create(&userInfo)
}

func TestGormGeo(t *testing.T) {
	dsn := "root:123456@tcp(192.168.3.88:3306)/mall-boot?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	result := make([]GeoModel, 0)
	//
	lon := "116.2939670643206"
	lat := "40.04348155309347"
	withDist := 6000
	db.Model(GeoModel{}).
		Debug().
		Select(fmt.Sprintf("*,(st_distance(POINT(longitude,latitude), POINT( %s, %s ))* 111195) AS distance", lon, lat)).
		Having("distance < ?", withDist).
		Order("distance ASC").
		Find(&result)

	for _, item := range result {
		fmt.Println("item = ", item.Distance, item)
	}

}

type GeoModel struct {
	Id        int64   `gorm:"primarykey" json:"id"`
	Name      string  `gorm:"column:name" json:"name"`
	Address   string  `gorm:"column:address" json:"address"`
	Longitude string  `gorm:"column:longitude" json:"longitude"`
	Latitude  string  `gorm:"column:latitude" json:"latitude"`
	Distance  float64 `json:"distance"`
}

func (*GeoModel) TableName() string {
	return "geo"
}

//SELECT
//*,
//	( st_distance ( POINT ( longitude, latitude ), POINT (116.3424590000, 40.0497810000))* 111195 / 1000 ) AS distance
//FROM
//geo
//HAVING distance<50
//ORDER BY
//distance ASC
