package search

import (
	"fmt"

	"github.com/slack-go/slack"
)

// Search return result of query
func Search(token, query string) *slack.SearchMessages {
	c := slack.New(token)
	searchParams := slack.SearchParameters{
		Sort:          "timestamp",
		SortDirection: "desc",
		Count:         100,
		Page:          1,
	}
	res, _ := c.SearchMessages(query, searchParams)
	fmt.Printf("[USER-INFO]message hits: %v\n", res.TotalCount)
	return res
}
