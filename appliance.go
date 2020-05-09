package natureremo

import (
	"encoding/json"
	"time"
)

type Appliance struct {
	ID         string         `jons:"id"`
	Device     DeviceCore     `json:"device"`
	Model      ApplianceModel `json:"model"`
	Nickname   string         `json:"nickname"`
	Image      string         `json:"image"`
	Type       string         `json:"type"`
	Settings   AirConParams   `json:"settings"`
	AirCon     AirCon         `json:"aircon"`
	Signals    []Signal       `json:"signals"`
	Tv         Tv             `json:"tv"`
	Light      Light          `json:"light"`
	SmartMeter SmartMeter     `json:"smart_meter"`
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

type Tv struct {
	State   TvState  `json:"state"`
	Buttons []Button `json:"buttons"`
}

type TvState struct {
	Input string `json:"input"`
}

type Button struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Label string `json:"label"`
}

type Light struct {
	State   LightState `json:"state"`
	Buttons []Button   `json:"buttons"`
}

type LightState struct {
	Brightness string `json:"brightness"`
	Power      string `json:"power"`
	Lastbutton string `json:"last_button"`
}

type SmartMeter struct {
	EchonetliteProperties []EchonetliteProperty `json:"echonetlite_properties"`
}

type EchonetliteProperty struct {
	Name      string    `json:"name"`
	Epc       int64     `json:"epc"`
	Val       string    `json:"val"`
	UpdatedAt time.Time `json:"updated_at"`
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
