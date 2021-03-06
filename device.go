package natureremo

import (
	"encoding/json"
	"time"
)

type SensorValue struct {
	Value     float64   `json:"val"`
	CreatedAt time.Time `json:"created_at"`
}

type NewestEvents struct {
	Temperature  SensorValue `json:"te"`
	Humidity     SensorValue `json:"hu"`
	Illumination SensorValue `json:"il"`
}
type Device struct {
	ID                string       `jons:"id"`
	Name              string       `josn:"name"`
	TemperatureOffset int          `json:"temperature_offset"`
	HumidityOffset    int          `json:"humidity_offset"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	FirmwareVersion   string       `json:"firmware_version"`
	MacAddress        string       `json:"mac_address"`
	SerialNumber      string       `json:"serial_number"`
	NewestEvents      NewestEvents `json:"newest_events"`
}

func (c *Client) GetDevices() ([]Device, error) {
	resp, err := c.Get("/devices")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	devices := make([]Device, 0)
	err = json.NewDecoder(resp.Body).Decode(&devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}
