package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

const URL = "https://api.akl.gg/"

var BotToken string
var Client *http.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	BotToken = os.Getenv("DISCORD_TOKEN")
	Client = &http.Client{}
}

func API(path string) (*http.Response, error) {
	return Client.Get(URL + path)
}

type Layouts []string

func main() {
	s, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info("Bot is ready")
	})

	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot {
			return
		}
		log.Infof("[%s] %s", m.Author, m.Content)

		Reply := func(msg string) {
			if _, err := s.ChannelMessageSendReply(
				m.ChannelID,
				msg,
				m.Reference(),
			); err != nil {
				log.Error(err)
			}
		}

		if command, ok := strings.CutPrefix(m.Content, "!"); ok {
			switch strings.TrimSpace(command) {
			case "ping":
				Reply("pong")
			case "list":
				if res, err := API("layouts"); err != nil {
					Reply(fmt.Sprintf("Error contacting API: %s", err))
					return
				} else {
					defer res.Body.Close()
					var body []byte
					var layouts Layouts
					body, err := io.ReadAll(res.Body)
					if err != nil {
						Reply(fmt.Sprintf("Error reading response: %s", err))
						return
					}
					if err := json.Unmarshal(body, &layouts); err != nil {
						Reply(fmt.Sprintf("Error parsing response: %s", err))
						return
					}

					trimmed :=  strings.Join(layouts[:10], "\n")
					if len(layouts) > MaxLayouts {
						msg := fmt.Sprintf(
							"Found %d layouts"
						)
					} else {
						msg := fmt.Sprintf(
							"Found %d layouts"
						)
					}
					Reply(msg)
				}
			}
		}
	})

	s.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	if err := s.Open(); err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Info("Graceful shutdown")
}
