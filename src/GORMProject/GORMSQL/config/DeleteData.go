package config

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// DeleteData1 删除一条记录
func DeleteData1(db *gorm.DB) {
	// 删除ID = 16
	db.Delete(&User{ID: 16})
	// 带额外条件的删除
	db.Where("name = ?", "zs").Delete(&User{})
}

// DeleteData2 根据主键删除
func DeleteData2(db *gorm.DB) {
	db.Delete(&User{}, 15)
	db.Delete(&User{}, "14")
	db.Delete(&User{}, []int{2, 3})
}

// BeforeDelete Delete Hook
func (u *User) BeforeDelete(db *gorm.DB) error {
	if u.Name == "zs" {
		return errors.New("user name is not zs")
	}
	return nil
}

// DeleteData3 批量删除
func DeleteData3(db *gorm.DB) {
	db.Where("name like ?", "%l%").Delete(&User{})
	db.Delete(&User{}, "name is null")
}

// DeleteData4 阻止全局删除
func DeleteData4(db *gorm.DB) {
	err := db.Delete(&User{}).Error
	fmt.Println(err)

	db.Where("1 = 1").Delete(&User{})

	db.Exec("DELETE FROM user")

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
}

// DeleteData5 ## 软删除
func DeleteData5(db *gorm.DB) {
	db.Delete(&UserInfo{
		Model: gorm.Model{ID: 1},
	})

	// 批量删除
	db.Where("id in (?)", []int{2, 3, 4}).Delete(&UserInfo{})

	// 在查询时会忽略被软删除的记录
	var userInfo []UserInfo
	db.Where("name = ?", "test").Find(&userInfo)
	fmt.Println(userInfo)
}

// DeleteData6 ## 软删除
func DeleteData6(db *gorm.DB) {
	db.Where("id in (?)", []int{1, 2, 3, 4}).Delete(&TestModel{})
}

// GetDeleteData1 ### 查找被软删除的记录
func GetDeleteData1(db *gorm.DB) {
	var userInfo []UserInfo
	db.Unscoped().Where("id > ?", 1).Find(&userInfo)
	fmt.Println(userInfo)
}

// DeleteDataUnscoped ### 永久删除
func DeleteDataUnscoped(db *gorm.DB) {
	db.Unscoped().Delete(&UserInfo{Model: gorm.Model{ID: 1}})
}
