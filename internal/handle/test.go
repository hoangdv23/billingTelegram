package handle

import (
	"log"
	telebot "gopkg.in/telebot.v4" // Alias 'telebot' cho thư viện
)

func Tester(bot *telebot.Bot) {
	// Xử lý lệnh /test
	bot.Handle("/1", func(c telebot.Context) error {
        // Define the file path
		file := &telebot.Document{
            File: telebot.FromDisk("/root/billingTeleBot/assets/excel/Cdr_1900_call_IN_Digitel_10_2024.xlsx"),
            FileName: "Cdr_1900_call_IN_Digitel_10_2024_CustomName.xlsx",
        }

        // Send the file to the user who triggered the command
        err := c.Reply("Chờ tí")
        if err != nil {
            log.Println("Error sending file:", err)
            return err
        }
        return c.Send(file)
    })
	
}
