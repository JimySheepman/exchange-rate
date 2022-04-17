package main

import (
	"log"
	"os"
	"strings"

	"er-rabbit-consumer/receive"
)

func main() {
	switch bodyFrom(os.Args) {
	case "TRY":
		receive.ReceiveTRY()
	case "USD":
		receive.ReceiveUSD()
	case "EUR":
		receive.ReceiveEUR()
	default:
		log.Fatal("Wrong Args")
	}

}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		log.Fatal("give a argument")
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
