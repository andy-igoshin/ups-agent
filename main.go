package main

import (
//	"fmt"

	"git.vsu.ru/ai/ups-agent/config"
	"git.vsu.ru/ai/ups-agent/status"
)

func main() {

	config := config.New()

	if config.Status {
		status.Print(config)
		return
	}


}
