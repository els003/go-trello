package trello

import (
	"encoding/json"
	"net/url"
	"strings"
)

func (l *List) AddBottomCard(opts Card) (*Card, error) {
	opts.IdList = l.Id

	payload := url.Values{}
	payload.Set("name", opts.Name)
	payload.Set("desc", opts.Desc)
	payload.Set("pos", "bottom")
	payload.Set("due", opts.Due)
	payload.Set("idList", opts.IdList)
	payload.Set("idMembers", strings.Join(opts.IdMembers, ","))

	body, err := l.client.Post("/cards", payload)
	if err != nil {
		return nil, err
	}

	var card Card
	if err = json.Unmarshal(body, &card); err != nil {
		return nil, err
	}
	card.client = l.client
	return &card, nil
}
