package main

import (
	"log"

	"github.com/isoment/go-overview/helpers"
)

func main() {
	sum := helpers.Add(2, 2)
	greet, goodbye := helpers.SaySomething("Hi", "See ya")
	log.Println(sum)
	log.Println(greet)
	log.Println(goodbye)
}
