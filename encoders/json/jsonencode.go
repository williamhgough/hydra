package main

import (
	"encoding/json"
	"log"
	"os"
)

type CrewMember struct {
	ID                int      `json:"id, omitempty"`
	Name              string   `json:"name"`
	SecurityClearance int      `json:"clearance"`
	AccessCodes       []string `json:"access_codes"`
}

type ShipInfo struct {
	ShipID    int
	ShipClass string
	Captain   CrewMember
}

func main() {
	f, err := os.Create("jfile.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cm := CrewMember{1, "William", 10, []string{"FED"}}
	si := ShipInfo{1, "Cruiser", cm}

	err = json.NewEncoder(f).Encode(&si)
	if err != nil {
		log.Fatal(err)
	}
}
