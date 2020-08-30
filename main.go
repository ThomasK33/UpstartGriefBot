package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/reiver/go-telnet"
)

// Command - Twitch Command structure
type Command struct {
	requiredVotes int
	currentVotes  int
	handler       func()
}

func main() {
	caller := &telnetCaller{}

	gc := &GameCommands{
		tc: caller,
	}

	commands := map[string]*Command{
		"!disconnect":  &Command{3, 0, gc.disconnect},
		"!quit":        &Command{3, 0, gc.quit},
		"!sell":        &Command{3, 0, gc.sellUnit},
		"!spray":       &Command{3, 0, gc.boardSpray},
		"!benchUnit":   &Command{3, 0, gc.benchUnit},
		"!fakeDown":    &Command{3, 0, gc.fakeGCDown},
		"!levelup":     &Command{3, 0, gc.levelup},
		"!lock":        &Command{3, 0, gc.lock},
		"!reroll":      &Command{3, 0, gc.reroll},
		"!toggle":      &Command{3, 0, gc.toggle},
		"!cameraDown":  &Command{3, 0, gc.cameraDown},
		"!cameraUp":    &Command{3, 0, gc.cameraUp},
		"!enemiesDown": &Command{3, 0, gc.enemiesDown},
		"!enemiesUp":   &Command{3, 0, gc.enemiesUp},
		"!away":        &Command{3, 0, gc.away},
		"!home":        &Command{3, 0, gc.home},
		"!opponent":    &Command{3, 0, gc.opponent},
		"!dps":         &Command{3, 0, gc.dps},
		"!sharecode":   &Command{3, 0, gc.sharecode},
	}

	go func(c *telnetCaller) {
		// // or client := twitch.NewAnonymousClient() for an anonymous user (no write capabilities)
		client := twitch.NewClient(os.Getenv("TWITCH_BOT_NAME"), os.Getenv("TWITCH_OAUTH_TOKEN"))

		client.OnPrivateMessage(func(message twitch.PrivateMessage) {
			fmt.Println(message.Message)

			for key, value := range commands {
				if strings.HasPrefix(message.Message, key) {
					if value.currentVotes >= value.requiredVotes {
						value.handler()
						value.currentVotes = 0
					} else {
						value.currentVotes++
					}
				}
			}

			if strings.HasPrefix(message.Message, "!buy") {
				slotString := strings.TrimSpace(strings.TrimPrefix(message.Message, "!buy"))

				if slot, err := strconv.Atoi(slotString); err == nil {
					if slot >= 1 && slot <= 5 {
						gc.buySlot(slot)
					}
				}
			}
		})

		client.OnConnect(func() {
			fmt.Println("Connected to Twitch")
			client.Say(os.Getenv("TWITCH_CHANNEL_NAME"), "Bring it on!")
		})

		client.Join(os.Getenv("TWITCH_CHANNEL_NAME"))

		twitchErr := client.Connect()
		if twitchErr != nil {
			panic(twitchErr)
		}
	}(caller)

	err := telnet.DialToAndCall("localhost:27015", caller)

	if err != nil {
		log.Fatal(err)
	}
}
