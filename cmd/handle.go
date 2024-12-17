package cmd

import (
	// "billing/cmd/middleware"
	"billing/internal/handle"

	tele "gopkg.in/telebot.v4"
)

func Handle(b *tele.Bot) {

	handle.Start(b)
	handle.Tester(b)
	handle.Login(b)
	handle.Cdr_Vas(b)


}
