package actions

import "fmt"

func GetHardwareInfo(address string) (string, error) {
	return fmt.Sprintf("Getting General Hardware Info: %s\n", address), nil
}
