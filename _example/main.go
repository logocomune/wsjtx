package main

import (
	"context"
	"encoding/hex"
	"github.com/logocomune/wsjtx/message"
	"github.com/logocomune/wsjtx/udpserver"
	"log"
)

func main() {

	server, err := udpserver.NewServer(context.Background(), udpserver.Multicast, udpserver.DefaultPort, log.Default())
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	for r := range server.Read() {
		parse, err := message.Parse(r)
		if err != nil {
			log.Println("Error:", err)
		}
		log.Printf("Message Type: %s\n", parse.ResponseType)
		log.Printf("Raw message (hex): %s\n", hex.EncodeToString(r))
		log.Printf("Decoded message: %+v\n", parse.Message)
		log.Println("-----------------------------------------------------")
	}
}
