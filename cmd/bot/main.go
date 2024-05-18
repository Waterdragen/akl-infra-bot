package main

import (
	"os"
	"os/signal"

	"github.com/akl-infra/bot/internal/handler"

	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

var BotToken string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	BotToken = os.Getenv("DISCORD_TOKEN")
}

func main() {
	log.Info("Booting up...")
	s, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	s.AddHandler(handler.OnReady)
	s.AddHandler(handler.OnMessage)

	s.Identify.Intents = discordgo.IntentDirectMessages | discordgo.IntentGuildMessages

	if err := s.Open(); err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Info("Graceful shutdown")
}
