package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ThomasK33/UpstartGriefBot/pkg/underlords"
	"github.com/gempir/go-twitch-irc/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/reiver/go-telnet"
)

// Version and build vars linked at compile time
var (
	Version string
	Build   string
)

// Command - Twitch Command structure
type Command struct {
	requiredVotes int
	currentVotes  int
	handler       func(string)
}

func main() {
	log.Println("Starting " + Version + " (" + Build + ")")

	channelName, botName, oauthToken := os.Getenv("TWITCH_CHANNEL_NAME"), os.Getenv("TWITCH_BOT_NAME"), os.Getenv("TWITCH_OAUTH_TOKEN")
	telnetServerAddress, isTelnetServerAddressSet := os.LookupEnv("TELNET_ADDRESS")

	caller := &underlords.TelnetCaller{}
	gc := &underlords.GameCommands{Caller: caller}

	commands := map[string]*Command{
		"!disconnect":  {3, 0, gc.Disconnect},
		"!quit":        {3, 0, gc.Quit},
		"!buy 1":       {3, 0, gc.BuySlot},
		"!buy 2":       {3, 0, gc.BuySlot},
		"!buy 3":       {3, 0, gc.BuySlot},
		"!buy 4":       {3, 0, gc.BuySlot},
		"!buy 5":       {3, 0, gc.BuySlot},
		"!sell":        {3, 0, gc.SellUnit},
		"!spray":       {3, 0, gc.BoardSpray},
		"!benchUnit":   {3, 0, gc.BenchUnit},
		"!fakeDown":    {3, 0, gc.FakeGCDown},
		"!levelup":     {3, 0, gc.Levelup},
		"!lock":        {3, 0, gc.Lock},
		"!reroll":      {3, 0, gc.Reroll},
		"!toggle":      {1, 0, gc.Toggle},
		"!cameraDown":  {3, 0, gc.CameraDown},
		"!cameraUp":    {3, 0, gc.CameraUp},
		"!enemiesDown": {3, 0, gc.EnemiesDown},
		"!enemiesUp":   {3, 0, gc.EnemiesUp},
		"!away":        {3, 0, gc.Away},
		"!home":        {3, 0, gc.Home},
		"!opponent":    {3, 0, gc.Opponent},
		"!dps":         {3, 0, gc.Dps},
		// "!code":        {3, 0, gc.Sharecode},
	}

	// TODO: Refactor this into person specific counter
	// counter := map[string]map[string]time.Time{}
	// counter["greycodes"] = map[string]time.Time{}
	// counter["greycodes"]["!buy"] = time.Now()

	go func() {
		client := twitch.NewClient(botName, oauthToken)

		client.OnPrivateMessage(func(message twitch.PrivateMessage) {
			log.Println(message.User.DisplayName, message.Message)

			for key, value := range commands {
				if strings.HasPrefix(message.Message, key) {
					value.currentVotes++

					if value.currentVotes >= value.requiredVotes {
						argString := strings.TrimSpace(message.Message)
						argStrings := strings.Fields(argString)
						argString = strings.Join(argStrings[1:], " ")

						value.handler(argString)
						value.currentVotes = 0

						client.Say(channelName, "Executed "+key)
					} else {
						client.Say(channelName, "Need "+fmt.Sprint(value.requiredVotes-value.currentVotes)+" more votes to execute "+key)
					}
				}
			}

			if strings.HasPrefix(message.Message, "!help") {
				var builder strings.Builder

				builder.WriteString("!help")
				for key := range commands {
					builder.WriteString(", " + key)
				}

				builder.WriteString(", " + "!keep")

				client.Say(channelName, "Following commands are available: "+builder.String())
			}

			if strings.HasPrefix(message.Message, "!keep") {
				if commands["!sell"].currentVotes > 0 {
					commands["!sell"].currentVotes--
				}
				client.Say(channelName, "Need "+fmt.Sprint(commands["!sell"].requiredVotes-commands["!sell"].currentVotes)+" more votes to execute !sell")
			}

			isBroadcaster := message.User.Badges["broadcaster"] != 0
			if isBroadcaster {
				if strings.HasPrefix(message.Message, "!vote") {
					commands["!sell"].currentVotes = 0
					client.Say(channelName, "Reset sell vote count to 0")
				}
			}

		})

		client.OnConnect(func() {
			log.Println("Connected to Twitch")
			client.Say(channelName, "Bring it on!")
		})

		client.Join(channelName)

		twitchErr := client.Connect()
		if twitchErr != nil {
			log.Println("An error occurred while connecting to Twitch")
			panic(twitchErr)
		}
	}()

	if !isTelnetServerAddressSet {
		telnetServerAddress = "localhost:27015"
	}

	log.Println("Connecting to " + telnetServerAddress)

	err := telnet.DialToAndCall(telnetServerAddress, caller)

	if err != nil {
		log.Println("An error occurred while connecting to the telnet server")
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
}
