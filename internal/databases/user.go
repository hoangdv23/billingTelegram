package databases

import (
	"log"
	"gorm.io/gorm"
)

// Kiểm tra người dùng billing với user_name
func CheckUserBilling(db *gorm.DB, user_name string) string {
	// Sử dụng phương thức First hoặc Find của GORM để tìm người dùng trong bảng 'users'
	var found_user_name string

	// Thực hiện truy vấn tìm người dùng
	err := db.Table("users").
		Where("company_name = ? AND user_code = ?", "digitel", user_name).
		Pluck("user_name", &found_user_name).Error

	// Xử lý lỗi nếu có
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Nếu không tìm thấy người dùng
			log.Printf("Không tìm thấy người dùng với tên: %s", user_name)
			return "" // Trả về chuỗi rỗng nếu không tìm thấy người dùng
		}
		// Nếu có lỗi khác, log lỗi để dễ dàng debug
		log.Printf("Lỗi khi truy vấn dữ liệu cho người dùng %s: %v", user_name, err)
		return "" // Trả về chuỗi rỗng nếu có lỗi trong quá trình truy vấn
	}

	// Nếu tìm thấy người dùng, trả về user_name
	log.Printf("Đã tìm thấy người dùng: %s", found_user_name)
	return found_user_name
}
