package databases

import (
	"log"
	"gorm.io/gorm"
)

type Action struct {
	Fullname  string `gorm:"column:fullname"`
	Name      string `gorm:"column:name"`
	Username  string `gorm:"column:username"`
	UserTele  string `gorm:"column:user_tele"`
	ChatID    string `gorm:"column:chat_id"`
	Actions   string `gorm:"column:actions"`
}

func WriteLog(db *gorm.DB, act Action) bool {
	// Sử dụng phương thức Create() để chèn bản ghi mới vào bảng "actions"
	if err := db.Create(&act).Error; err != nil {
		log.Printf("Lỗi khi chèn log vào cơ sở dữ liệu: %v", err)
		return false
	}
	return true
}