package handle

import (
	"billing/pkg/helpers"

	tele "gopkg.in/telebot.v4"
)


func Start(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		intro := "Chào mừng bạn đến với bot Billing. Vui lòng đăng nhập tài khoản billing của bạn"
		return c.Send(intro)
	})
	b.Handle(&helpers.Back_To_Main_Menu, func(c tele.Context) error {
		intro := "Đã trở lại menu chính. Mọi thắc mắc vui lòng liên hệ @hoangdv123"
		return c.Send(intro, helpers.MainMenu_InlineKeys)
	})
}

