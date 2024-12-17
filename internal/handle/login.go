package handle

import(
	"billing/internal/databases"
	"billing/pkg/helpers"
	"fmt"
	"gorm.io/gorm"
	tele "gopkg.in/telebot.v4"
)

func loginHandle(db *gorm.DB,c tele.Context) error {
	// ===== gửi thông báo lấy username ==========
	first_input := helpers.TrimSpace(c.Text())
	isUser := databases.CheckUserBilling(db,first_input)
	if isUser != "" {
		output := fmt.Sprintf("Xin chào %s, chào mừng bạn đến với bot Billing Digitel", isUser)
		return c.Reply(output,helpers.MainMenu_InlineKeys)
	}

	return c.Reply("Tài khoản của bạn chưa đúng, vui lòng nhập lại")
}

func Login(b *tele.Bot) {
    b.Handle(tele.OnText, func(c tele.Context) error {
		billingDB := databases.GetBillingDB()
        return loginHandle(billingDB, c) // Gọi loginHandle với db và context
    })
}