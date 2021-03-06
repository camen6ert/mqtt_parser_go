package mqttparser

import (
	"testing"
)

func TestHandleConnect(t *testing.T) {

	b := [...]byte{0x10, 0x45, 0x00, 0x04, 0x4d, 0x51, 0x54, 0x54, 0x05, 0x02, 0x00, 0x3c, 0x00, 0x00, 0x38, 0x6d, 0x71, 0x74, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x4d, 0x51, 0x54, 0x54, 0x5f, 0x35, 0x5f, 0x30, 0x2d, 0x37, 0x38, 0x32, 0x62, 0x36, 0x33, 0x35, 0x63, 0x2d, 0x62, 0x63, 0x37, 0x37, 0x2d, 0x34, 0x30, 0x62, 0x32, 0x2d, 0x39, 0x32, 0x61, 0x36, 0x2d, 0x38, 0x32, 0x39, 0x31, 0x39, 0x63, 0x65, 0x64, 0x64, 0x36, 0x32, 0x65, 0x00}
	c, err := HandleConnect(b[2:])

	if err != nil {
		t.Error("Faiure")
	}

	if c.ClientID != "mqttClient-MQTT_5_0-782b635c-bc77-40b2-92a6-82919cedd62e" {
		t.Error("ClientID is false")
	}
}
