package trello

import (
	"fmt"
	"net/url"
	"strconv"
)

func (ts *Card) DestroySelfFromTrello() error {
	_, err := ts.client.Delete("/cards/" + ts.Id)
	if err != nil {
		return err
	}
	return nil
}
func (ts *Card) SyncFromTrello() error {
	c, err := ts.client.Card(ts.Id)
	if err != nil {
		return err
	}
	*ts = *c
	return nil
}

func (c *Card) SetLabels(ids []string) error {

	values := make(map[string][]string)
	values["value"] = ids

	_, err := c.client.Put("/cards/"+c.Id+"/idLabels", values)
	if err != nil {
		return err
	}
	return nil
}

/*
func (c *Card) AddLabel(LID string) error {
	values := make(map[string][]string)

	values["value"] = append(values["value"], LID)
	_, err := c.client.Put("/cards/"+c.Id+"/idLabels", values)
	if err != nil {
		return err
	}
	return nil
}
*/

func (c *Card) DeleteMember(mid string) error {
	_, err := c.client.Delete("/cards/" + c.Id + "/idMembers/" + mid)
	return err
}

func (c *Card) SetMembers(ids []string) error {

	values := make(map[string][]string)
	values["value"] = ids

	_, err := c.client.Put("/cards/"+c.Id+"/idMembers", values)
	if err != nil {
		return err
	}
	return nil

}
func (c *Card) AddMember(m *Member) error {

	var ids []string
	ms, err := c.Members()
	if err != nil {
		return err
	}
	for _, cm := range ms {
		ids = append(ids, cm.Id)
	}
	ids = append(ids, m.Id)
	return c.SetMembers(ids)
}

func (c *Card) SetClosed(closed bool) error {
	values := url.Values{}
	values.Add("value", strconv.FormatBool(closed))

	_, err := c.client.Put("/cards/"+c.Id+"/closed", values)
	if err != nil {
		fmt.Println("setCardclosed error", err)
		return err
	}
	return nil
}
func (c *Card) Rename(newName string) error {
	values := make(map[string][]string)
	values["value"] = append(values["value"], newName)
	_, err := c.client.Put("/cards/"+c.Id+"/name", values)
	if err != nil {
		return err
	}
	c.Name = newName
	return nil
}

// DeleteComment :
func (c *Card) DeleteComment(actionID string) error {
	_, err := c.client.Delete("/cards/" + c.Id + "/actions/" + actionID + "/comments")
	return err
}

// DeleteLabel :
func (c *Card) DeleteLabel(LID string) error {
	_, err := c.client.Delete("/cards/" + c.Id + "/idLabels/" + LID)
	return err
}

func (c *Card) MoveCard(pos float64) error {
	values := url.Values{}
	values.Add("value", strconv.FormatFloat(pos, 'g', -1, 64))

	_, err := c.client.Put("/cards/"+c.Id+"/pos", values)
	return err
}

func (c *Card) MoveToList(idList string) error {
	values := url.Values{}
	values.Add("value", idList)
	_, err := c.client.Put("/cards/"+c.Id+"/idList", values)
	return err
}

//move card to another list
// PUT /1/cards/[card id or shortlink]/idList
