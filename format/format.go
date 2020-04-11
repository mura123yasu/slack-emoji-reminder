package format

import (
	"fmt"

	"github.com/slack-go/slack"
)

func Format(messages *slack.SearchMessages) string {
	message := ""
	fmt.Printf("[USER-INFO]message hit: %v\n", messages.TotalCount)

	for index := range messages.Matches {
		fmt.Println(messages.Matches[index].Text)
		message = message + messages.Matches[index].Text + "\n" + "```" + messages.Matches[index].Permalink + "```" + "\n"
	}

	return message
}
