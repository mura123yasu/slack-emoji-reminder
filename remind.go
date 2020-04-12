package reminder

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/mura123yasu/slack-emoji-reminder/notice"
	"github.com/mura123yasu/slack-emoji-reminder/search"
)

// Remind post message to target channel
func Remind(ctx context.Context, msg *pubsub.Message) error {

	var (
		reaction      = os.Getenv("REACTION")
		user          = os.Getenv("USER")
		botToken      = os.Getenv("SLACK_BOT_TOKEN")
		remindChannel = os.Getenv("SLACK_REMIND_CHANNEL")
		apiToken      = os.Getenv("SLACK_API_TOKEN")
		searchChannel = os.Getenv("SLACK_SEARCH_CHANNEL")
	)

	// get data to remind from slack
	query := "has::" + reaction + ": " + "from:@" + user + " " + "in:#" + searchChannel
	fmt.Printf("[USER-INFO]query string: %v\n", query)
	result := search.Search(apiToken, query)

	// send to slack
	notifier := notice.New(botToken, remindChannel)
	ts, err := notifier.Notify(result)

	if err != nil {
		panic(fmt.Sprintf("[USER-ERROR]notify failed. %v\n", err.Error()))
	}

	fmt.Printf("[USER-INFO]message sent via slack-emoji-reminder to %v. timestamp: %v\n", remindChannel, ts)
	return nil
}
