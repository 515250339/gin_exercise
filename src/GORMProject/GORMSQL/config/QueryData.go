package config

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// QueryTest1 ## 检索单个对象 `First`、`Take`、`Last`
func QueryTest1(db *gorm.DB) {
	// 获取第一条记录（主键升序）
	user1 := User{}
	db.First(&user1)
	fmt.Println(user1)

	// 获取一条记录，没有指定排序字段
	user2 := User{}
	db.Take(&user2)
	fmt.Println(user2)

	// 获取最后一条记录（主键降序）
	user3 := User{}
	db.Last(&user3)
	fmt.Println(user3)

	result := db.First(&user1)
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(result.Error)        // returns error
	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// QueryTest2 ## 检索单个对象 `First`、`Take`、`Last`
func QueryTest2(db *gorm.DB) {
	// 可以
	var user User
	db.First(&user)
	fmt.Println(user)
	fmt.Println()

	// 可以
	result := map[string]interface{}{}
	db.Model(&User{}).First(result)
	fmt.Println(result)
	fmt.Println()

	// 不行
	result2 := map[string]interface{}{}
	db.Table("user").First(result2)
	fmt.Println(result2)
	fmt.Println()

	// 但可以配合 Take 使用
	result3 := map[string]interface{}{}
	db.Table("user").Take(result3)
	fmt.Println(result3)
	fmt.Println()

	// 根据第一个字段排序
	result4 := db.First(&User{})
	fmt.Println(result4.RowsAffected)
}

// QueryTest3 根据主键检索
func QueryTest3(db *gorm.DB) {
	user := User{}
	db.First(&user, 10)
	fmt.Println(user)
	fmt.Println()

	user2 := User{}
	db.First(&user2, "10")
	fmt.Println(user2)
	fmt.Println()

	users := []User{{}}
	db.Find(&users, []int{1, 2, 3})
	for _, user := range users {
		fmt.Println(user)
	}
	fmt.Println()

}

// QueryALlTest ## 检索全部对象
func QueryALlTest(db *gorm.DB) {
	users := []User{{}}
	result := db.Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)

}

// FilterTest String 条件
func FilterTest(db *gorm.DB) {
	// 获取第一条匹配的记录
	var user User
	db.Where("name = ?", "ee").First(&user)
	fmt.Println(1, user)

	// 获取全部匹配的记录
	users := []User{{}}
	result := db.Where("name = ?", "zs").Find(&users)
	fmt.Println(2, result.RowsAffected)

	// IN
	users = []User{{}}
	result2 := db.Where("name IN ?", []string{"zs", "aa"}).Find(&users)
	fmt.Println(3, result2.RowsAffected)

	// LIKE
	users = []User{{}}
	result3 := db.Where("name like ?", "%z%").Find(&users)
	fmt.Println(4, result3.RowsAffected)

	// AND
	users = []User{{}}
	result4 := db.Where("name = ? AND age >= ?", "zs", 18).Find(&users)
	fmt.Println(5, result4.RowsAffected)

	// Time
	users = []User{{}}
	result5 := db.Where("updated_at > ?", "2022-06-21 00:00:00").Find(&users)
	fmt.Println(6, result5.RowsAffected)

	// BETWEEN
	users = []User{{}}
	result6 := db.Where("created_at BETWEEN ? AND ?", "2022-06-20 00:00:00", "2022-06-21 00:00:00").Find(&users)
	fmt.Println(7, result6.RowsAffected)
}

// FilterMapAndStruct Struct & Map 条件
func FilterMapAndStruct(db *gorm.DB) {
	//Struct
	var user User
	db.Where(&User{Name: "zs", Age: 18}).First(&user)
	fmt.Println(user)

	// Map
	users := []User{{}}
	db.Where(map[string]interface{}{"name": "ls", "age": 22}).Find(&users)
	fmt.Println(users)

	// 主键切片条件
	users = []User{{}}
	db.Where([]int{1, 3, 5}).Find(&users)
	fmt.Println("主键切片条件", users)

}

// FilterTest2 内联条件
func FilterTest2(db *gorm.DB) {
	// 根据主键获取记录，如果是非整型主键
	var user User
	db.First(&user, "id = ?", "5")
	fmt.Println(1, user)

	// Plain SQL
	user = User{}
	db.Find(&user, "name = ?", "zs")
	fmt.Println(2, user)

	users := []User{{}}
	db.Find(&users, "name <> ? AND age > ?", "zs", 20)
	fmt.Println(3, users)

	// Struct
	users = []User{{}}
	db.Find(&users, User{Age: 22})
	fmt.Println(4, users)

	// Map
	users = []User{{}}
	db.Find(&users, map[string]interface{}{"age": 20})
	fmt.Println(5, users)

}

// FilterNot 内联条件
func FilterNot(db *gorm.DB) {
	var user User
	db.Not("name = ?", "zs").First(&user)
	fmt.Println(1, user)

	// Not In
	users := []User{{}}
	db.Not(map[string]interface{}{"name": []string{"zs", "ls"}}).Find(&users)
	fmt.Println(2, users)

	// Struct
	user = User{}
	db.Not(User{Name: "zs", Age: 18}).First(&user)
	fmt.Println(3, user)

	// 不在主键切片中的记录
	user = User{}
	db.Debug().Not([]int{1, 2, 3}).First(&user)
	fmt.Println(4, user)

}

// FilterOr Or 条件
func FilterOr(db *gorm.DB) {
	users := []User{{}}
	db.Where("name = ?", "zs").Or("name = ?", "ls").Find(&users)
	fmt.Println(1, users)

	// Struct
	users = []User{{}}
	db.Where("name = 'zs'").Or(User{Name: "ls", Age: 22}).Find(&users)
	fmt.Println(2, users)

	// Map
	users = []User{{}}
	db.Where("name = ?", "zs").Or(map[string]interface{}{"name": "ls", "age": 22}).Find(&users)
	fmt.Println(3, users)

}

// SelectField ## 选择特定字段
func SelectField(db *gorm.DB) {
	users := []User{{}}
	db.Select("name", "age").Find(&users)

	users = []User{{}}
	db.Select([]string{"name", "age"}).Find(&users)

	users = []User{{}}
	db.Table("user").Select("COALESCE(age,?)", 21).Rows()
}

// OrderData 指定从数据库检索记录时的排序方式
func OrderData(db *gorm.DB) {
	users := []User{{}}
	db.Order("age desc, name").Find(&users)

	// 多个 order
	users = []User{{}}
	db.Order("age desc").Order("name").Find(&users)

	users = []User{{}}
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}},
			WithoutParentheses: true},
	}).Find(&users)
}
