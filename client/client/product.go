package client

import (
	"fmt"
)

func (c *Client) GetProduct(productID string) error {
	url := fmt.Sprintf("/products/%s", productID)
	return c.sendRequest("GET", url, nil)
}
