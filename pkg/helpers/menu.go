package helpers

import (
	tele "gopkg.in/telebot.v4"
)

var Login = tele.Btn{
	Text:   "Đăng nhập",
	Data:   "button_login",
}

var Intro = tele.InlineButton{
	Unique: "btn_callback1_gioithieu",
	Text:   "Giới thiệu",
	Data:   "button_about",
}
var Cdr = tele.InlineButton{
	Unique: "btn_cdr",
	Text: "CTC 1900 1800 DIGITEL",
	Data: "cdr",
}

var Cdr_Vas_Telco = tele.InlineButton{
	Unique: "btn_cdr",
	Text: "CTC 1900 1800 TELCO",
	Data: "cdr_vas_telco",
}

var InfoNumber = tele.InlineButton{
	Unique: "btn_infor_number",
	Text: "Thông tin đầu số",
	Data: "infor_number",
}

var MainMenu_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{Cdr,	Cdr_Vas_Telco},
		{InfoNumber},
		{ Intro},
		// {AskMenu}, // Row 4: Button 4
	},
}
var Quantity = tele.InlineButton{
	Unique: "btn_quantity",
	Text:   "Báo cáo sản lượng",
	Data:   "button_quantity",
}
var Back_To_Main_Menu = tele.InlineButton{
	Unique: "btn_callback1_ask_menu",
	Text:   "Trở về menu chính",
	Data:   "button1_clicked",
}

var AskMenu_InlineKeys = &tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{Back_To_Main_Menu},
	},
}
