package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	url   = "https://api.nature.global/1/devices"
	token = ""
)

type SensorValue struct {
	Value     int       `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

type NewestEvents struct {
	Te SensorValue `json:"te"`
	Hu SensorValue `json:"hu"`
}
type Device struct {
	ID                string       `jons:"id"`
	Name              string       `josn:"name"`
	TemperatureOffset int          `json:"temperature_offset"`
	HumidityOffset    int          `json:"humidity_offset"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	FirmwareVersion   string       `json:"firmware_version"`
	NewestEvents      NewestEvents `json:"newest_events"`
}

func main() {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var d []Device
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range d {
		fmt.Println(device)
	}
}
