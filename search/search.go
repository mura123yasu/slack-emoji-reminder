package search

import (
	"github.com/slack-go/slack"
)

func Search(slackAPIToken, query string) *slack.SearchMessages {
	c := slack.New(slackAPIToken)
	searchParams := slack.SearchParameters{
		Sort:          "timestamp",
		SortDirection: "desc",
		Count:         100,
		Page:          1,
	}
	res, _ := c.SearchMessages(query, searchParams)
	return res
}
