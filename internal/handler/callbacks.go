package handler

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
)

const Trigger = "!akl"

func parseArgs(str string) []string {
	re := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`)
	matches := re.FindAllString(str, -1)
	return matches
}

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Info("Bot is ready")
}

func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	ch, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Errorf("Error getting channel: %s", err)
		return
	}
	isDM := ch.Type == discordgo.ChannelTypeDM
	var isMention bool
	if me, err := s.User("@me"); err != nil {
		log.Errorf("Error looking up our User: %s", err)
		return
	} else {
		for _, mention := range m.Mentions {
			if mention.ID == me.ID {
				isMention = true
			}
		}
	}

	// Match user/nick/role/channel mentions
	mentionRegex := regexp.MustCompile(`<@!?\d+>|<@&\d+>|<#\d+>`)

	// Remove mentions from the message content
	content := mentionRegex.ReplaceAllString(m.Content, "")

	if isDM {
		log.Infof("(DM) [%s] %s", m.Author, content)
	} else if isMention {
		log.Infof("(@) [%s] %s", m.Author, content)
	} else {
		log.Infof("[%s] %s", m.Author, content)
	}

	// Check for trigger
	if !(isDM || isMention) {
		var ok bool
		if strings.HasPrefix(m.Content, "!") {
			content, ok = strings.CutPrefix(content, Trigger)
			if !ok {
				log.Info("Didn't find trigger")
				return
			}
		} else {
			// Ignore unprefixed channel messages
			return
		}
	}

	args := parseArgs(content)
	command := args[0]
	if len(args) > 1 {
		args = args[1:]
	} else {
		args = []string{}
	}
	log.Infof("Cmd: %s, Args: %s", command, args)

	reply, err := Dispatch(command, args)
	if err != nil {
		log.Errorf("Error in command: %s %s => %s", command, args, err)
	}
	if _, err := s.ChannelMessageSendReply(
		m.ChannelID,
		reply,
		m.Reference(),
	); err != nil {
		log.Error(err)
	}
}
