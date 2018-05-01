package main

import (
	"fmt"
	"log"

	natureremo "github.com/70-10/nature-remo-go"
)

func main() {
	c := natureremo.NewClient("****************************")
	devices, err := c.GetDevices()
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range devices {
		fmt.Println("ID  : " + device.ID)
		fmt.Println("Name: " + device.Name)
	}

}
