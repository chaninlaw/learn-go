package main

import (
	"fmt"

	"github.com/chaninlaw/go-example/chaninlaw"
	// "github.com/chaninlaw/go-example/chaninlaw/internal/ninja" // internal lib
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	chaninlaw.SayHello()
	chaninlaw.SayTest()
	fmt.Printf("UUID: %s", id)
}
