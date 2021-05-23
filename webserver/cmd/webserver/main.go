package main

//

import (
	"log"
	"webserver/webserver"
)

func main() {

	log.Printf("START")
	webserver.Start(8080)
	log.Printf("END")

}
