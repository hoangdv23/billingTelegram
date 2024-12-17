package helpers

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"strconv"
	tele "gopkg.in/telebot.v4"

)

// ============ Lấy thông tin ngày tháng năm ================
func Get_this_month_with_year() string {
	now := time.Now()
	currentMonth := now.Month()
	currentYear := now.Year()
	return fmt.Sprintf("%02d/%d", currentMonth, currentYear)
}

func GetLastMonth_with_year() string {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)
	return fmt.Sprintf("%02d/%d", lastMonth.Month(), lastMonth.Year())
}

func GetLast2Month_with_year() string {
	now := time.Now()
	lastMonth := now.AddDate(0, -2, 0)
	return fmt.Sprintf("%02d/%d", lastMonth.Month(), lastMonth.Year())
}

// ========== trimspace để loại bỏ khoảng trắng,.... =============
func TrimSpace(space string) string {
	removeSpace := strings.TrimSpace(space)
	return removeSpace
}
// ========= Xử lý ký tự giống nhau =================
func RegexString(sub string, data string) bool{
	// substring := fmt.Sprintf(".*\\|%s$", sub) 
	substring := fmt.Sprintf("^%s.*", sub) 
	match,_ := regexp.MatchString(substring, data)

	return match
}
// ========= Xử lý ký tự giống nhau, yêu cầu phải có ký tự trước =================
func RegexPreString(sub string, data string) bool{
	substring := fmt.Sprintf(`^[^\|]+\|%s$`, sub)
	// substring := fmt.Sprintf("^%s.*", sub) 
	match,_ := regexp.MatchString(substring, data)

	return match
}
// ========= Xử lý ký tự giống nhau, IN OUT, yêu cầu phải có cdr| =================
func RegexCdrVasStringCallType(data string) bool {
	// Regex để kiểm tra chuỗi có định dạng cdr|...|IN hoặc cdr|...|OUT
	substring := `^cdr\|[A-Za-z0-9_]+\|(IN|OUT)$`
	match, _ := regexp.MatchString(substring, data)
	return match
}

// ========= Xử lý ký tự giống nhau, IN OUT, yêu cầu phải có ký tự trước =================
func RegexCdrVasStringMonth(data string) bool {
	substring := `^cdr\|[A-Za-z0-9_]+\|(IN|OUT)\|\d{2}/\d{4}$`
	match, _ := regexp.MatchString(substring, data)
	return match
}
//======== Xử lý lấy nhà mạng ============
func Extract_telco(data string) (string, bool) {
	// Tách phần sau "btn_telco_"
	re := regexp.MustCompile(`^btn_telco_(\w+)$`)
	matches := re.FindStringSubmatch(data)

	if len(matches) > 1 {
		return matches[1], true // Trả về MBC nếu có khớp
	}
	return "", false // Trả về "" nếu không tìm thấy
}
// ====== Xử lý lấy telo vas ============
func Extract_telco_vas_type(data string) (string, string, string, string, string, bool) {
    // Biểu thức chính quy mới để tách cdr, vas, call_type, month, và year
    re := regexp.MustCompile(`^(\w+)\|(\w+)\|(\w+)\|(\d{2})/(\d{4})$`)
    matches := re.FindStringSubmatch(data)

    // Kiểm tra xem có khớp không
    if len(matches) == 6 {
        // Trả về các giá trị đã tách
        return matches[1], matches[2], matches[3], matches[4], matches[5], true
    }

    // Nếu không khớp với bất kỳ định dạng nào
    return "", "", "", "", "", false
}


// ========= Tạo Button Call Type tự động =========================
func DynamicCallType(text string,data string) tele.InlineButton {
	return tele.InlineButton{
		Unique: data,
		Text:   text,
	}
}

//===== lấy tháng động_theo call_type IN hoặc OUT
func DynamicMonth(text string,data string) tele.InlineButton {
	return tele.InlineButton{
		Unique: fmt.Sprintf("btn_%s", data),
		Text:   fmt.Sprintf("Button từ %s", text),
	}
}
// ======== tách month year =========
func ParseMonthYear(monthYear string) (int, int) {
	parts := strings.Split(monthYear, "/")
	if len(parts) == 2 {
		month, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0 // Nếu không thể chuyển đổi tháng thành int, trả về 0
		}
		year, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, 0 // Nếu không thể chuyển đổi năm thành int, trả về 0
		}
		return month, year // Trả về tháng và năm dưới dạng int
	}
	return 0, 0 // Nếu không có dữ liệu hợp lệ, trả về 0
}
