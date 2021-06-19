package main

import (
	"flag"
	"fmt"

	"github.com/ifvictr/clubhouse-go"
)

func main() {
	token := flag.String("token", "", "Clubhouse account auth token")
	flag.Parse()

	client := clubhouse.NewFromToken(*token)
	res, _, err := client.GetChannels()
	if err != nil {
		fmt.Printf("Could not get channels: %s\n", err)
		return
	}

	// Print out channels seen in the hallway screen
	for i, channel := range res.Channels {
		var topic string
		if channel.Topic != nil {
			topic = *channel.Topic
		}

		fmt.Printf("%2d | %s | %s | %s\n", i+1, getVisibilityEmoji(channel.IsPrivate, channel.IsSocialMode), channel.Channel, topic)
	}
}

func getVisibilityEmoji(private bool, social bool) string {
	switch true {
	case private:
		return "🔒"
	case social:
		return "👥"
	default:
		return "🌎"
	}
}
