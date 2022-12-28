package actions

import (
	"context"
	"fmt"
	"log"
)

func SetBlank(ctx context.Context, address string, status bool) error {
	if status {
		log.Printf("blanking %s\n", address)
		//Blank ON = set to android input
		//status = false
	} else {
		log.Printf("un-blanking %s, restoring previous input\n", address)
		//Blank OFF = set to previous input
		//status = true
	}
	return nil
}

func GetBlanked(address string) (string, error) {
	return fmt.Sprintf("Getting Blank Status: %s", address), nil
}
