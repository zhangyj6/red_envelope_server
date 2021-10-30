package sql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var userId int64 = 0

type User struct {
	ID    int64
	Count int64
}

func (User) TableName() string {
	return "user"
}

func GetUser(uid int64) (user User) {
	db, err := gorm.Open("mysql", "root:zyz123456@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.FirstOrCreate(&user, User{ID: uid})
	return
}

func UpdateCount(user *User) {
	db, err := gorm.Open("mysql", "root:zyz123456@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	user.Count++
	db.Model(&user).Update("count", user.Count)
}
