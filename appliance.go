package natureremo

import (
	"encoding/json"
)

type Appliance struct {
	ID       string         `jons:"id"`
	Device   DeviceCore     `json:"device"`
	Model    ApplianceModel `json:"model"`
	Nickname string         `json:"nickname"`
	Image    string         `json:"image"`
	Type     string         `json:"type"`
	Settings AirConParams   `json:"settings"`
	AirCon   AirCon         `json:"aircon"`
	Signals  []Signal       `json:"signals"`
}

type ApplianceModel struct {
	ID           string `jons:"id"`
	Manufacturer string `json:"manufacturer"`
	RemoteName   string `json:"remote_name"`
	Name         string `json:"name"`
	Image        string `json:"image"`
}

type AirConParams struct {
	Temp   string `json:"temp"`
	Mode   string `json:"mode"`
	Vol    string `json:"vol"`
	Dir    string `json:"dir"`
	Button string `json:"button"`
}

type AirCon struct {
	Range    Range  `json:"range"`
	TempUnit string `json:"tempUnit"`
}

type Range struct {
	Modes        Modes    `json:"modes"`
	FixedButtons []string `json:"fixedButtons"`
}

type Modes struct {
	Cool AirConRangeMode `json:"cool"`
	Warm AirConRangeMode `json:"warm"`
	Dry  AirConRangeMode `json:"dry"`
	Blow AirConRangeMode `json:"blow"`
	Auto AirConRangeMode `json:"auto"`
}
type AirConRangeMode struct {
	Temp []string `json:"temp"`
	Vol  []string `json:"vol"`
	Dir  []string `json:"dir"`
}

func (c *Client) GetAppliances() ([]Appliance, error) {
	resp, err := c.Get("/appliances")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	appliances := make([]Appliance, 0)
	err = json.NewDecoder(resp.Body).Decode(&appliances)
	if err != nil {
		return nil, err
	}
	return appliances, nil
}
