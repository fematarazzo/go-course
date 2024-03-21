package main

import (
	"fmt"
	"log"
	"os"

	"cli-tool/app"
)

func main() {
	fmt.Println("---CLI Application: Search for IPs---")

	application := app.Generate()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
