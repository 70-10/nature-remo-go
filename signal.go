package natureremo

import (
	"encoding/json"
	"fmt"
)

type Signal struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (c *Client) GetSignals(applianceID string) ([]Signal, error) {
	resp, err := c.Get(fmt.Sprintf("/appliances/%s/signals", applianceID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	signals := make([]Signal, 0)
	err = json.NewDecoder(resp.Body).Decode(&signals)
	if err != nil {
		return nil, err
	}
	return signals, nil
}

func (c *Client) SendSignal(signalID string) error {
	_, err := c.Post(fmt.Sprintf("/signals/%s/send", signalID), nil)
	if err != nil {
		return err
	}
	return nil
}
