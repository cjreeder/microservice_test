package actions

import "fmt"

func GetHardwareInfo(address string) (string, error) {
	return fmt.Printf("Getting General Hardware Info: %s\n", address), nil
}
