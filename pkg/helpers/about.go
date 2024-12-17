package helpers

import (
	tele "gopkg.in/telebot.v4"
)

var (
	intro = "Phần mềm đang được phát triển"
)

func dispath(c tele.Context) {
	c.Send(intro)
	// c.Send(photo)
	// c.Send("⭕🅿♌⚡🌴⛎♈📈❌✌Ⓩ", helpers.AskMenu_InlineKeys)
}
func About(b *tele.Bot) {
	b.Handle("/about", func(c tele.Context) error {
		dispath(c)
		return nil
	})

	b.Handle("about", func(c tele.Context) error {
		dispath(c)
		return nil
	})

	// b.Handle(&helpers.Intro, func(c tele.Context) error {
	// 	dispath(c)
	// 	return nil
	// })
}
