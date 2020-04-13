package notice

import (
	"fmt"
	"time"

	"github.com/slack-go/slack"
)

// Notifier has target information
type Notifier struct {
	client  *slack.Client
	channel string
}

// New returns notifier object
func New(token, channel string) *Notifier {
	client := slack.New(token)
	return &Notifier{
		client:  client,
		channel: channel,
	}
}

// Notify send message to specified channel
func (n *Notifier) Notify(messages *slack.SearchMessages) (string, error) {
	// post head message
	t := time.Now()
	const layout = "2006-01-02"
	ts, err := n.postInline(":izakaya_lantern: " + t.Format(layout) + "のSlackEmojiRemind :izakaya_lantern:")

	if err != nil {
		fmt.Printf("[USER-ERROR]head post failed. %v\n", err.Error())
		return ts, err
	}

	// post messages to thread
	for index := range messages.Matches {
		fmt.Println(messages.Matches[index].Text)
		message := messages.Matches[index].Text + "\n" + messages.Matches[index].Permalink + "\n" + "```" + messages.Matches[index].Timestamp + "```" + "\n"
		err = n.postThreadInline(message, ts)
		if err != nil {
			fmt.Printf("[USER-ERROR]thread post failed. %v\n", err.Error())
			break
		}
	}
	return ts, err
}

// Return post message's timestamp to post in the thread
func (n *Notifier) postInline(text string) (string, error) {
	_, ts, err := n.client.PostMessage(
		n.channel,
		slack.MsgOptionText(text, false),
	)
	return ts, err
}

// ts is parent message's timestamp to post in the thread
func (n *Notifier) postThreadInline(text, ts string) error {
	_, _, err := n.client.PostMessage(
		n.channel,
		slack.MsgOptionText(text, false),
		slack.MsgOptionTS(ts),
	)
	return err
}
