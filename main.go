package main

import (
	"log"

	"github.com/GoProgramming/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
	}
}
