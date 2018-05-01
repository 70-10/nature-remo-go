package natureremo

import "encoding/json"

type User struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

func (c *Client) GetUser() (*User, error) {
	resp, err := c.Get("/users/me")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
