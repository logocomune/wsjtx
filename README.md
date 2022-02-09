# WSJTX
[![CircleCI](https://circleci.com/gh/logocomune/wsjtx/tree/main.svg?style=svg)](https://circleci.com/gh/logocomune/wsjtx/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/logocomune/wsjtx)](https://goreportcard.com/report/github.com/logocomune/wsjtx)
[![codecov](https://codecov.io/gh/logocomune/wsjtx/branch/main/graph/badge.svg?token=GGN3PHjyZV)](https://codecov.io/gh/logocomune/wsjtx)
[![CodeFactor](https://www.codefactor.io/repository/github/logocomune/wsjtx/badge)](https://www.codefactor.io/repository/github/logocomune/wsjtx)

Golang library for WSJTX-X provides:

- Functions for encoding and decoding of WSJT-X message up to version 2.5.2
- UDP Server for receiving messages from WSJT-X and sending to WSJT-X.

## Installation

```
go get github.com/logocomune/wsjtx
```

## Example

```go
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


```
