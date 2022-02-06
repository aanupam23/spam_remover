package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aanupam23/spam_remover/internal/config"
	"github.com/aanupam23/spam_remover/internal/handler"
	"github.com/bwmarrin/discordgo"
)

// Token variables will be used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Get configuration from config and env
	conf := config.ReadConfig()
	fmt.Printf("Spam Removal Duration is %d seconds\n", conf.Duration)

	if conf.BOT_TOKEN != "" {
		Token = conf.BOT_TOKEN
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Let's add the SpamHandler with the above created discord Server
	dg.AddHandler(handler.SpamHandler)

	// In this case, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("Goodbye!")

	// Cleanly close down the Discord session.
	dg.Close()
}
