package trello

import "encoding/json"

func (b *Board) Labels() (labels []Label, err error) {
	body, err := b.client.Get("/boards/" + b.Id + "/labels")
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &labels)

	return
}
