package client

import (
	"fmt"
)

func (c *Client) GetUser(userID string) error {
	url := fmt.Sprintf("/users/%s", userID)
	return c.sendRequest("GET", url, nil)
}
