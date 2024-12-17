package helpers

import (
	tele "gopkg.in/telebot.v4"
)

// ============ VAS ===================
var Vas_1900 = tele.InlineButton{
	Unique: "1900",
	Text: "Chi tiết cước 1900",
}
var Vas_1800 = tele.InlineButton{
	Unique: "1800",
	Text: "Chi tiết cước 1800",
}
// ==============  ===========
// ============ TELCO ===================
var Telco_VIETTEL = tele.InlineButton{
	Unique: "btn_telco_VIETTEL",
	Text: "VIETTEL",
	// Data: "button_VTL",
}

var Telco_MBC = tele.InlineButton{
	Unique: "btn_telco_MBC",
	Text: "MOBICAST",
	// Data: "button_MBC",
}
var Telco_GPC  = tele.InlineButton{
	Unique: "btn_telco_GPC",
	Text: "GPC",
	// Data: "button_GPC",
}
var Telco_ITEL = tele.InlineButton{
	Unique: "btn_telco_ITEL",
	Text: "ITEL",
	// Data: "button_ITEL",
}
var Telco_VNPT_FIXED = tele.InlineButton{
	Unique: "btn_telco_VNPT_FIXED",
	Text: "Cố định VNPT",
	// Data: "button_VNPT_FIXED",
}
var Telco_VMS = tele.InlineButton{
	Unique: "btn_telco_VMS",
	Text: "MOBIFONE",
	// Data: "button_VMS",
}
var Telco_VNM = tele.InlineButton{
	Unique: "btn_telco_VNM",
	Text: "VIETNAMOBILE",
	// Data: "button_VNM",
}

var Telco_CMC = tele.InlineButton{
	Unique: "btn_telco_CMC",
	Text: "Cố định CMC",
	// Data: "button_CMC",
}

var Telco_FPT = tele.InlineButton{
	Unique: "btn_telco_FPT",
	Text: "Cố định FPT",
	// Data: "button_FPT",
}

var Total_TELCO = tele.ReplyMarkup{
	InlineKeyboard: [][]tele.InlineButton{
		{Telco_VIETTEL,Telco_GPC,Telco_ITEL},
		{Telco_MBC,Telco_VNPT_FIXED,Telco_CMC},
		{Telco_VMS,Telco_VNM,Telco_FPT},
		{Back_To_Main_Menu},
	},
}
