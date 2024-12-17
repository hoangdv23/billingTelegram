package handle

import (
	"billing/internal/repository"
	"billing/internal/usecase"
	"billing/internal/databases"
	"billing/pkg/helpers"
	"fmt"

	tele "gopkg.in/telebot.v4"
)

// Khi click cdr
func Cdr_Vas(b *tele.Bot){
	b.Handle(&helpers.Cdr, func(ctx tele.Context) error {
		cdr := helpers.TrimSpace(ctx.Callback().Data)
		btn1900 := fmt.Sprintf("%s|1900",cdr)
		btn1800 := fmt.Sprintf("%s|1800",cdr)
		btnVAS:= fmt.Sprintf("%s|VAS",cdr)

		btn_1900 := helpers.DynamicCallType("1900",btn1900)
		btn_1800 := helpers.DynamicCallType("1800",btn1800)
		btn_VAS := helpers.DynamicCallType("Cả 1900 và 1800",btnVAS)

		replyMarkup_1900_1800 := tele.ReplyMarkup{
			InlineKeyboard: [][]tele.InlineButton{
				{btn_1900, btn_1800,btn_VAS},
				{helpers.Back_To_Main_Menu},
			},
		}
		return ctx.Send("Bạn muốn lấy chi tiết cước theo GTGT nào?", &tele.SendOptions{
			ParseMode: tele.ModeHTML,  // Chỉ định sử dụng HTML
			ReplyMarkup: &replyMarkup_1900_1800,
		})
	})
	b.Handle(&helpers.Cdr_Vas_Telco, func(c tele.Context) error {
		cdr := helpers.TrimSpace(c.Callback().Data)
		btn1900 := fmt.Sprintf("%s|1900",cdr)
		btn1800 := fmt.Sprintf("%s|1800",cdr)
		btnVAS:= fmt.Sprintf("%s|VAS",cdr)

		btn_1900 := helpers.DynamicCallType("1900",btn1900)
		btn_1800 := helpers.DynamicCallType("1800",btn1800)
		btn_VAS := helpers.DynamicCallType("Cả 1900 và 1800",btnVAS)

		replyMarkup_1900_1800 := tele.ReplyMarkup{
			InlineKeyboard: [][]tele.InlineButton{
				{btn_1900, btn_1800,btn_VAS},
				{helpers.Back_To_Main_Menu},
			},
		}
		return c.Send("Bạn muốn lấy CTC Digitel gọi dịch vụ 1800 hay 1900", &tele.SendOptions{
			ParseMode: tele.ModeHTML,  // Chỉ định sử dụng HTML
			ReplyMarkup: &replyMarkup_1900_1800,
		})
	})
	// Xử lý dịch vụ cần lấy (1900 1800 VAS)
	b.Handle(tele.OnCallback, func(c tele.Context) error {
		callback := helpers.TrimSpace(c.Callback().Data)
		if callback == "cdr|1900" || callback == "cdr|1800" || callback == "cdr|VAS"{
			_, vas,_,_,_,_ := helpers.Extract_telco_vas_type(callback)
			datacallIn := fmt.Sprintf("%s|IN",callback)
			datacallOut := fmt.Sprintf("%s|OUT",callback)
			Call_in := helpers.DynamicCallType("Call IN Digitel",datacallIn)
			Call_out := helpers.DynamicCallType("Call OUT TELCO",datacallOut)
			replyMarkup_In_Out := tele.ReplyMarkup{
				InlineKeyboard: [][]tele.InlineButton{
					{Call_in, Call_out},
					{helpers.Back_To_Main_Menu},
				},
			}
			text_reponds := fmt.Sprintf("<STRONG>Tuyệt vời!!! </STRONG> Bot sẽ lấy Chi tiết cước theo <b>%s</b>. Giờ hãy chọn 1 trong 2: \n- <b>TELCO gọi tới %s Digitel</b>. \n- <b>Digitel gọi tới %s TELCO</b>",vas,vas,vas)
			return c.Send(text_reponds, &tele.SendOptions{
				ParseMode: tele.ModeHTML,  // Chỉ định sử dụng HTML
				ReplyMarkup: &replyMarkup_In_Out,
			})

		}
		// xử lý lấy giá trị cdr|...|call_type
		if helpers.RegexCdrVasStringCallType(callback) {
			_,vas,call_type,_,_,_ := helpers.Extract_telco_vas_type(callback)
			this_month := fmt.Sprintf("Tháng %s",helpers.Get_this_month_with_year()) 
			last_month := fmt.Sprintf("Tháng %s",helpers.GetLastMonth_with_year()) 
			last_2_month := fmt.Sprintf("Tháng %s",helpers.GetLast2Month_with_year()) 

			unique_this_month := helpers.Get_this_month_with_year()
			unique_last_month := helpers.GetLastMonth_with_year()
			unique_last_2_month := helpers.GetLast2Month_with_year()

			data_thisMonth := fmt.Sprintf("%s|%s",callback,unique_this_month)
			data_lastMonth := fmt.Sprintf("%s|%s",callback,unique_last_month)
			data_last2Month := fmt.Sprintf("%s|%s",callback,unique_last_2_month)

			text_thisMonth := helpers.DynamicCallType(this_month,data_thisMonth)
			text_lastMonth := helpers.DynamicCallType(last_month,data_lastMonth)
			text_last2Month := helpers.DynamicCallType(last_2_month,data_last2Month)

			replyMarkup_month := tele.ReplyMarkup{
				InlineKeyboard: [][]tele.InlineButton{
					{text_thisMonth, text_lastMonth, text_last2Month},
					{helpers.Back_To_Main_Menu},
				},
			}
			text_reponds := fmt.Sprintf("Bot xử lý lấy chi tiết cước <b>%s</b> - Call <b>%s</b> . \nHãy chọn tháng cần lấy",vas,call_type)
			return c.Reply(text_reponds, &tele.SendOptions{
				ParseMode: tele.ModeHTML,  // Chỉ định sử dụng HTML
				ReplyMarkup: &replyMarkup_month,
			})
		}
		if helpers.RegexCdrVasStringMonth(callback) {
			_,vas,call_type,month,year,_ := helpers.Extract_telco_vas_type(callback)
			db136 := databases.GetDB136()
			repo := repository.NewCDRRepository(db136) // Đây là repository cần truyền vào
			cdrUsecaseInstance := usecase.NewCdrUsecase(repo)
			filepath,filename := cdrUsecaseInstance.Fetch_Cdr_vas(year,month,vas,call_type)

			file := &tele.Document{
				File: tele.FromDisk(filepath),
				FileName: filename,
			}
			text := fmt.Sprintf("Bot gửi file cdr %s - Call %s tháng %s năm %s",vas,call_type,month,year)
			return c.Send(text,file)
		}
		return c.Send(callback)
	})

}
