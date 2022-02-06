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

		// If IgnoreGroup is specified, then check if user is part of ignored group
		if config.C.IgnoreGroup != "" {
			// First check if member is in IgnoreGroup
			allRolesID := m.Member.Roles
			for _, roleID := range allRolesID {
				rolesStruct, err := s.State.Role(m.GuildID, roleID)
				if err != nil {
					panic(err)
				}
				if rolesStruct.Name == config.C.IgnoreGroup {
					return
				}
			}
		}

		// If IgnoreDuration is specified, then check users time duration in group
		if config.C.IgnoreDuration > 0 {
			// Now Check if User is in Ignore Duration
			joinedAt := m.Member.JoinedAt
			messageTime := m.Timestamp

			// Time Duration for which user has been part of community
			spendDuration := messageTime.Sub(joinedAt)

			// Converting the Member Group Duration in Weeks
			spendDurationinWeek := spendDuration.Hours() / 24 / 7

			// If user has spent more time then required duration then message is okay to be posted
			if spendDurationinWeek > config.C.IgnoreDuration {
				return
			}
		}

		// Now only spammers are present and we need to take action!
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
