package main

import (
	"fmt"

	"github.com/bandrosh/auth-api/internal/config"
)

func runApplication() {
	fmt.Println("Executing application")

	config.LoadConfig()
}
