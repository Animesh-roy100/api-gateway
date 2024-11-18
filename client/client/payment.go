package client

func (c *Client) MakePayment(payment map[string]interface{}) error {
	return c.sendRequest("POST", "/payments", payment)
}
