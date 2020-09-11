package main

//go:generate gotext -srclang=en update -out=catalog/catalog.go -lang=en,el

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.Greek)
	p.Printf("Hello world!")
	p.Println()

	p.Printf("Hello", "world!")
	p.Println()

	person := "Alex"
	place := "Utah"

	p.Printf("Hello ", person, " in ", place, "!")
	p.Println()

	// Greet everyone.
	p.Printf("Hello world!")
	p.Println()

	city := "Munich"
	p.Printf("Hello %s!", city)
	p.Println()

	// Person visiting a place.
	p.Printf("%s is visiting %s!",
		person,
		place)
	p.Println()

	// Double arguments.
	miles := 1.2345
	p.Printf("%.2[1]f miles traveled (%[1]f)", miles)
}
