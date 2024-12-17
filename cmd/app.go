package cmd

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Không thể nạp tệp .env: %v", err)
	}
	// billingDB := databases.GetBillingDB()
	// db136 := databases.GetDB136()
	// dbLog := databases.GetDBLog()

	//================= Bot ================
	token := os.Getenv("TOKEN")
	Pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(Pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Hi, your bot is running as expected!")
	Handle(b)
	b.Start()
}
