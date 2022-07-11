package config

import (
	"fmt"
	"gorm.io/gorm"
)

// UpdateData1 保存所有字段
func UpdateData1(db *gorm.DB) {
	var user User
	db.First(&user)

	user.Name = "zs2"
	user.Age = 98
	db.Save(&user)
}

// UpdateData2 更新单个列
func UpdateData2(db *gorm.DB) {
	// 条件更新
	db.Model(&User{}).Where("age = ?", 98).Update("age", "89")

	var user User
	// User 的 ID 是 1
	user.ID = 1
	db.Model(&user).Update("name", "hello")

	// 根据条件和model的值进行更新
	user = User{}
	db.Model(&user).Where("ID = ?", 1).Update("name", "hello2")
}

// UpdateData3 更新多个列
func UpdateData3(db *gorm.DB) {
	// 根据 `struct` 更新属性，只会更新非零值的字段
	var user User
	user.ID = 2
	db.Model(&user).Updates(User{
		Name: "hello2",
		Age:  89,
	})

	// 根据 `map` 更新属性
	user = User{}
	user.ID = 1
	db.Model(&user).Updates(map[string]interface{}{
		"name": "zs",
		"age":  18,
	})
}

// UpdateData4 更新选定字段  Select 和 Map
func UpdateData4(db *gorm.DB) {
	user := User{
		ID: 1,
	}
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 28})

	user = User{
		ID: 2,
	}
	db.Model(&user).Omit("name").Updates(map[string]interface{}{
		"name": "ls",
		"age":  23,
	})

	// Select 和 Struct （可以选中更新零值字段）
	user = User{
		ID: 3,
	}
	db.Model(&user).Select("Name", "Age").Updates(User{
		Name: "word",
		Age:  0,
	})
}

// BeforeUpdate 更新选定字段 对于更新操作，GORM 支持 `BeforeSave`、`BeforeUpdate`、`AfterSave`、`AfterUpdate` 钩子，这些方法将在更新记录时被调用
//func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
//	if u.Name == "zs" {
//		return errors.New("zs unique")
//	}
//
//	return
//}

// UpdateData5 批量更新
func UpdateData5(db *gorm.DB) {
	// 根据 struct 更新
	db.Model(&User{}).Where("id in (?)", []int{10, 11, 12}).Updates(User{
		Name: "test123",
		Age:  99,
	})

	// 根据 map 更新
	db.Table("user").Where("id in ?", []int{1, 2, 3}).Updates(map[string]interface{}{
		"name": "test234",
		"age":  88,
	})
}

// UpdateData6 阻止全局更新
func UpdateData6(db *gorm.DB) {
	db.Model(&User{}).Update("name", "zs")

	db.Model(&User{}).Where("1 = 1").Update("name", "zs")

	db.Exec("UPDATE user set name = ?", "ls")

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("age", 20)
}

// UpdateData7 更新的记录数
func UpdateData7(db *gorm.DB) {
	// 通过 `RowsAffected` 得到更新的记录数
	result := db.Model(&User{}).Where("id in ?", []int{10, 11, 12, 13}).Updates(User{Name: "zs"})
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
}

// UpdateData8 更新的记录数
func UpdateData8(db *gorm.DB) {
	db.Model(&User{}).Where("age = ?", 20).Update("age", gorm.Expr("age * ? + ?", 2, 1))

	db.Model(&User{}).Where("age = 40")
}

// UpdateData9 使用 SQL 表达式更新
func UpdateData9(db *gorm.DB) {
	user := User{
		ID: 3,
	}
	db.Model(&user).Update("age", gorm.Expr("age * ? + ?", 2, 100))

	user = User{
		ID: 2,
	}
	db.Model(&user).Updates(map[string]interface{}{"age": gorm.Expr("age * ? + ?", 2, 10)})

	user = User{ID: 1}
	db.Model(&user).UpdateColumn("age", gorm.Expr("age - ?", 100))

	user = User{}
	db.Model(&user).Where("age > 100").UpdateColumn("age", gorm.Expr("age - ?", 1))
}

// UpdateData10 不使用 Hook 和时间追踪
func UpdateData10(db *gorm.DB) {
	db.Model(&User{ID: 1}).UpdateColumn("name", "test111")

	db.Model(&User{ID: 2}).UpdateColumns(User{Name: "test222", Age: 99})

	db.Model(&User{ID: 3}).Select("name", "age").UpdateColumns(User{
		Name: "test34333",
		Age:  88,
	})

}

// UpdateData11 根据子查询进行更新
func UpdateData11(db *gorm.DB) {
	db.Model(&User{ID: 1}).Update("name", db.Model(&UserInfo{}).Select("name").Where("user.id = user_infos.id"))

	db.Table("user as u").Where("name = ?", "test").Update("name", db.Table("user_infos as uf").Select("name").Where("u.id = uf.id"))

	db.Table("user as u").Where("name = ?", "test222").Updates(map[string]interface{}{"name": db.Table("user_infos as uf").Select("name").Where("u.id = uf.id")})
}

//func (u *User) BeforeUpdate(db *gorm.DB) (err error) {
//	if db.Statement.Changed("MemberNumber") {
//		return errors.New("MemberNumber not allowed to change")
//	}
//
//	if db.Statement.Changed("Name", "Age") {
//		db.Statement.SetColumn("Age", 98)
//	}
//
//	if db.Statement.Changed() {
//		db.Statement.SetColumn("email", "515@qq.com")
//	}
//	return nil
//}

// UpdateData12 检查字段是否有变更？
func UpdateData12(db *gorm.DB) {
	db.Model(&User{ID: 1}).Updates(map[string]interface{}{"name": "zs"})
	// Changed("Name") => true

	db.Model(&User{ID: 1, Name: "zs"}).Updates(map[string]interface{}{"name": "zs"})
	// Changed("Name") => false, 因为 `Name` 没有变更

	db.Model(&User{ID: 1, Name: "zs"}).Select("Age").Updates(map[string]interface{}{"name": "zs2", "age": 66})
	// Changed("Name") => false, 因为 `Name` 没有被 Select 选中并更新

	//测试失败
	//db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(User{Name: "jinzhu2"})
	//// Changed("Name") => true
	//db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(User{Name: "jinzhu"})
	//// Changed("Name") => false, 因为 `Name` 没有变更
	//db.Model(&User{ID: 1, Name: "jinzhu"}).Select("Admin").Updates(User{Name: "jinzhu2"})
	//// Changed("Name") => false, 因为 `Name` 没有被 Select 选中并更新

}

func (user *User) BeforeSave(db *gorm.DB) (err error) {
	if db.Statement.Changed("name") {
		user.Age += 20
		db.Statement.SetColumn("Age", user.Age)
	}
	return nil
}

// UpdateData13 在更新时修改值？
func UpdateData13(db *gorm.DB) {
	db.Model(&User{ID: 1}).Updates(map[string]interface{}{"name": "zs"})
	// Changed("Name") => true

}
