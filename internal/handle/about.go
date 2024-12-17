package handle

import (
	"billing/pkg/helpers"

	tele "gopkg.in/telebot.v4"
)

var (
	intro = "Đây là phần mềm bot hỗ trợ tự động lấy các báo cáo,giúp kế toán, Admin tiếp cận thông tin nhanh hơn. Thông tin mang giá trị tham khảo. \n\nHãy liên hệ trực tiếp với @dvhoang123 để biết thêm thông tin chi tiết"

	// photo = &tele.Photo{
	// 	Caption: "Mã QR truy cập bot", // Set caption here
	// 	File:    tele.FromURL("not yet"),
	// }
)

func dispath(c tele.Context) {
	// c.Send(intro)
	// c.Send(photo)
	c.Send(intro, helpers.AskMenu_InlineKeys)
}

func About(c tele.Context) error {
	dispath(c)
	return nil
}
