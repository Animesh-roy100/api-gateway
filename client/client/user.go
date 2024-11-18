package client

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetUser(id string) error {
	url := fmt.Sprintf("%s/users/%s", GatewayURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+AuthToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	fmt.Printf("User Response: %s\n", string(body))
	return nil
}
