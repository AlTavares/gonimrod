package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Usage = usage
	keyFlag := flag.String("k", "", "API Key")
	flag.Parse()
	message := strings.Join(flag.Args(), " ")
	key := *keyFlag
	if len(message) == 0 {
		log.Fatalln("Message can't be empty")
	}
	if len(key) == 0 {
		key = os.Getenv("NIMROD_KEY")
		if len(key) == 0 {
			log.Fatalln("API key can't be empty")
		}
	}
	nimrod := Nimrod{key}
	err := nimrod.sendMessage(message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message sent")
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: gonimrod -k [api_key] [message]
	
	Sends a message to the Nimrod api https://www.nimrod-messenger.io/
	You can set the environment variable NIMROD_KEY 
	with your api key and omit it when executing the program.

	If both are set, the key passed on execution takes precedence.
	`)
}
