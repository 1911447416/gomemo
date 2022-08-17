package models

import (
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Memo struct {
	Id     int
	Title  string
	Status int
}

func (Memo) TableName() string {
	return "memos"
}

func init() {
	AppConfig := LoadConfig()
	// 连接数据库 dsn="root:123456@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", AppConfig.MysqlConfig.User, AppConfig.MysqlConfig.Password, AppConfig.MysqlConfig.Ip, AppConfig.MysqlConfig.Port, AppConfig.MysqlConfig.Database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("open db err!err:", err)
	}

}

// 获取首页memo
func GetAllMemo() []Memo {
	memolist := []Memo{}
	DB.Find(&memolist)
	return memolist
}

// 创建memo
func CreateMemo(content string) bool {
	newmemo := &Memo{
		Title: content,
	}
	result := DB.Create(&newmemo)
	fmt.Println(result.RowsAffected)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

// 删除memo
func DeleteMemo(id string) bool {
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("string to int err! err:", err)
	}
	dmemo := &Memo{
		Id: num,
	}
	DB.Delete(&dmemo)
	return true
}

func DoneMemo(id string) bool {
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("string to int err! err:", err)
	}
	dmemo := Memo{Id: num}
	DB.Find(&dmemo)
	dmemo.Status = 1
	DB.Save(&dmemo)
	return true
}
func NoDoneMemo(id string) bool {
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("string to int err! err:", err)
	}
	dmemo := Memo{Id: num}
	DB.Find(&dmemo)
	dmemo.Status = 0
	DB.Save(&dmemo)
	return true
}
