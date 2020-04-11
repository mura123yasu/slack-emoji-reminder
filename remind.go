package reminder

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/mura123yasu/slack-emoji-reminder/format"
	"github.com/mura123yasu/slack-emoji-reminder/notice"
	"github.com/mura123yasu/slack-emoji-reminder/search"
)

func Remind(ctx context.Context, msg *pubsub.Message) error {

	var (
		reaction           = os.Getenv("REACTION")
		user               = os.Getenv("USER")
		slackBotToken      = os.Getenv("SLACK_BOT_TOKEN")
		slackRemindChannel = os.Getenv("SLACK_REMIND_CHANNEL")
		slackAPIToken      = os.Getenv("SLACK_API_TOKEN")
		slackSearchChannel = os.Getenv("SLACK_SEARCH_CHANNEL")
	)

	// get data to remind from slack
	query := "has::" + reaction + ": " + "from:@" + user + " " + "in:#" + slackSearchChannel
	fmt.Printf("[USER-INFO]query string: %v\n", query)
	result := search.Search(slackAPIToken, query)

	// format message
	message := format.Format(result)

	// send to slack
	notifier := notice.NewSlackNotifier(slackBotToken, slackRemindChannel)
	ts, err := notifier.Notify(message)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Printf("[USER-INFO]message sent via slack-emoji-reminder to %v. timestamp: %v\n", slackRemindChannel, ts)
	return nil
}
