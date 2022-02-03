package handler

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	modMessage = `Your message will self-destroy in 30 seconds, contact a Prefect or read the Community Rules in <#809043464964014100> to know why.
I'm a bot, bleep, bloop. 🤖`
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
		bm, _ := s.ChannelMessageSendReply(m.ChannelID, modMessage, m.Reference())
		go DelayedMessageRemover(s, bm.ChannelID, bm.ID)
	}
}

func DelayedMessageRemover(s *discordgo.Session, channelID string, messageID string) {
	time.Sleep(10 * time.Second)
	_ = s.ChannelMessageDelete(channelID, messageID)
}
