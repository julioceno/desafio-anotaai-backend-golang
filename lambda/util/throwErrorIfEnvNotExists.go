package util

import (
	"fmt"
	"log"
)

func ThrowErrorIfEnvNotExists(key string, value string) {
	if value == "" {
		log.Fatal(fmt.Sprintf("%s não existe", key))
	}
}
