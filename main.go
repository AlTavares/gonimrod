package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	message := flag.String("m", "", "Message to send")
	key := flag.String("k", "", "API Key")
	flag.Parse()
	nimrod := Nimrod{*key}
	if len(*message) == 0 {
		log.Fatalln("Message can't be empty")
	}
	if len(*key) == 0 {
		log.Fatalln("API key can't be empty")
	}
	err := nimrod.sendMessage(*message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message sent")
}
