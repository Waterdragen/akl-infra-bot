package handler

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
)

const Trigger = "!akl"

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

	var reply string
	if content == "" {
		reply = "No command provided"
	} else {
		commandArg := strings.SplitN(content, " ", 2)
		command := commandArg[0]
		arg := ""
		if len(commandArg) > 1 {
			arg = commandArg[1]
		}
		id, _ := strconv.ParseUint(m.Author.ID, 10, 64)

		reply, err = Dispatch(command, arg, id)
		if err != nil {
			log.Errorf("Error in command: %s %s => %s", command, arg, err)
		}
	}
	if _, err = s.ChannelMessageSendReply(m.ChannelID, reply, m.Reference()); err != nil {
		log.Error(err)
	}
}
