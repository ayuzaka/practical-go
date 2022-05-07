package chapter08

import (
	"encoding/json"
	"log"
	"os"
)

type Response struct {
	Type      string          `json:"type"`
	Timestamp int             `json:"timestamp"`
	Payload   json.RawMessage `json:"payload"`
}

type Message struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

type Sensor struct {
	ID       string `json:"id"`
	DeviceID string `json:"device_id"`
}

func DelayDecode(fileName string) any {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var resp Response
	if err := json.NewDecoder(f).Decode(&resp); err != nil {
		log.Fatal(err)
	}

	switch resp.Type {
	case "message":
		var m Message
		_ = json.Unmarshal(resp.Payload, &m)

		return m
	case "sensor":
		var s Sensor
		_ = json.Unmarshal(resp.Payload, &s)

		return s
	default:
		return nil
	}
}
