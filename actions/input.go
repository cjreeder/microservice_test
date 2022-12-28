package actions

import (
	"log"

	"context"
)

type Input struct {
	Input string `json:"input,omitempty"`
}

var CurrentInput = "hdmi1"

// GetInput returns the input being shown on the display
func GetInput(address string) (Input, error) {
	var output Input
	log.Printf("This returns the input on a device: %s\n", address)
	//Great information on how you get input from a Hex valued return
	// 3A 3A 30 31 47 3A 30 30 30 0D = Get Input
	/*
		payload := []byte{0x3A, 0x30, 0x31, 0x47, 0x3A, 0x30, 0x30, 0x30, 0x0D}
		log.Println("getting input")
		resp, err := sendCommand(address, payload)
		if err != nil {
			return output, err
		}
		log.Printf("got input. hex string: %x\n", payload)

		switch {
		case bytes.Contains(resp, []byte{0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x30, 0x31, 0x0D}):
			output.Input = "hdmi1"
			CurrentInput = "hdmi1"
		case bytes.Contains(resp, []byte{0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x30, 0x32, 0x0D}):
			output.Input = "hdmi2"
			CurrentInput = "hdmi2"
		case bytes.Contains(resp, []byte{0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x32, 0x31, 0x0D}):
			output.Input = "hdmi3"
			CurrentInput = "hdmi3"
		case bytes.Contains(resp, []byte{0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x32, 0x32, 0x0D}):
			output.Input = "hdmi4"
			CurrentInput = "hdmi4"
		case bytes.Contains(resp, []byte{0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x35, 0x31, 0x0D}):
			output.Input = "tv"
			CurrentInput = "tv"
		case bytes.Contains(resp, []byte{0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x30, 0x37, 0x0D}):
			output.Input = "dp"
			CurrentInput = "dp"
		case bytes.Contains(resp, []byte{0x3A, 0x30, 0x31, 0x72, 0x3A, 0x31, 0x30, 0x31, 0x0D}):
			output.Input = "android"
			CurrentInput = "android"
		case bytes.Contains(resp, []byte{0x3A, 0x30, 0x31, 0x72, 0x3A, 0x31, 0x30, 0x34, 0x0D}):
			output.Input = "type-c1"
			CurrentInput = "type-c1"
		case bytes.Contains(resp, []byte{0x3A, 0x30, 0x31, 0x72, 0x3A, 0x31, 0x30, 0x35, 0x0D}):
			output.Input = "type-c2"
			CurrentInput = "type-c2"
		default:
			return output, errors.New("unknown input")
		}
	*/
	output.Input = "hdmi1"
	return output, nil
}

// SetInput sets the input on the display
func SetInput(ctx context.Context, address string, input string) error {
	//var payload []byte

	log.Printf("Setting input %s on device: %s\n", input, address)
	CurrentInput = input

	/*
		switch input {
		case "tv":
			// 3A 3A 30 31 53 3A 30 35 31 0d = TV
			payload = []byte{0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x35, 0x31, 0x0d}
		case "hdmi1":
			// 3A 3A 30 31 53 3A 30 30 31 0d = HDMI1
			payload = []byte{0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x30, 0x31, 0x0d}
		case "hdmi2":
			// 3A 3A 30 31 53 3A 30 30 32 0d = HDMI2
			payload = []byte{0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x30, 0x32, 0x0d}
		case "hdmi3":
			// 3A 3A 30 31 53 3A 30 32 31 0d = HDMI3
			payload = []byte{0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x32, 0x31, 0x0d}
		case "hdmi4":
			// 3A 3A 30 31 53 3A 30 32 32 0d = HDMI4
			payload = []byte{0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x32, 0x32, 0x0d}
		case "dp":
			// 3A 3A 30 31 53 3A 30 30 37 0d = DP
			payload = []byte{0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x30, 0x37, 0x0d}
		case "android":
			// 3A 3A 30 31 53 3A 31 30 34 0d = Type-C1
			payload = []byte{0x3A, 0x30, 0x31, 0x53, 0x3A, 0x31, 0x30, 0x31, 0x0d}
		case "type-c1":
			// 3A 3A 30 31 53 3A 31 30 34 0d = Type-C1
			payload = []byte{0x3A, 0x30, 0x31, 0x53, 0x3A, 0x31, 0x30, 0x34, 0x0d}
		case "type-c2":
			// 3A 3A 30 31 53 3A 31 30 35 0d = Type-C2
			payload = []byte{0x3A, 0x30, 0x31, 0x53, 0x3A, 0x31, 0x30, 0x35, 0x0d}
		default:
			return errors.New("unknown input")
		}
		_, err := sendCommand(address, payload)
		if err != nil {
			return err
		}
	*/
	return nil
}
