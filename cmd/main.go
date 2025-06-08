package main

import (
	"fmt"

	"github.com/Saidurbu/go-qatarat-dashboard/internal/config/env"
)

func main() {

	e := env.NewEnv()
	fmt.Println("Running in environment:", e.Environment)

}
