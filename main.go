package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	. "github.com/AlTavares/gonimrod/nimrod"
)

func main() {
	flag.Usage = usage
	keyPointer := flag.String("k", "", "API Key")
	flag.Parse()
	message := getMessage()
	key := validateKey(keyPointer)
	nimrod := Nimrod{key}
	err := nimrod.SendMessage(message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message sent")
}

func getMessage() string {
	message := strings.Join(flag.Args(), " ")
	pipedMessage := getPipedMessage()
	if len(pipedMessage) != 0 {
		message += " " + pipedMessage
	}
	if len(message) == 0 {
		log.Fatalln("Message can't be empty")
	}
	return message
}

func validateKey(keyPointer *string) string {
	key := *keyPointer
	if len(key) == 0 {
		key = os.Getenv("NIMROD_KEY")
		if len(key) == 0 {
			log.Fatalln("API key can't be empty")
		}
	}
	return key
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: gonimrod -k [api_key] [message]
	
	Sends a message to the Nimrod api https://www.nimrod-messenger.io/
	You can set the environment variable NIMROD_KEY 
	with your api key and omit it when executing the program.

	If both are set, the key passed on execution takes precedence.
	`)
}

func getPipedMessage() string {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		// no pipe
		return ""
	}
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return ""
	}
	return string(bytes)
}
