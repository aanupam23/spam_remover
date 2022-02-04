package handler

import (
	"time"

	"github.com/aanupam23/spam_remover/internal/config"
	"github.com/bwmarrin/discordgo"
)

var (
	modMessage = `Your post will be removed!`
	C          *config.MainConfig
)

// SpamHandler checks for all message that are spam and takes action! PS: It is strict!
func SpamHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// We will check for the Moderation Comment by bot
	// Apart from that we will ignore all messages that created by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is a spam, then we start working on it!"
	if m.Content == "Hello" || m.Content == "Hello)" || m.Content == "Hi" || m.Content == "Hi)" || m.Content == "Hey" || m.Content == "Hey)" {
		go DelayedMessageRemover(s, m.ChannelID, m.ID)
		if config.C.ModMessage != "" {
			modMessage = config.C.ModMessage
		}
		bm, _ := s.ChannelMessageSendReply(m.ChannelID, modMessage, m.Reference())
		go DelayedMessageRemover(s, bm.ChannelID, bm.ID)
	}
}

func DelayedMessageRemover(s *discordgo.Session, channelID string, messageID string) {
	time.Sleep(time.Duration(config.C.Duration) * time.Second)
	_ = s.ChannelMessageDelete(channelID, messageID)
}
