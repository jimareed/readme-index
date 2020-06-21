package main

import (
	"flag"
	"log"

	"github.com/jimareed/readme-index/indexer"
	"github.com/jimareed/slides"
)

var mainDeck = slides.SlideDeck{}

func main() {

	input := flag.String("input", "", "path to source")
	output := flag.String("output", "", "path to source")
	help := flag.Bool("help", false, "help")

	flag.Parse()

	if *help {
		log.Fatal("usage: slide-generator [-input <path>][-output <path>][-server][-help]")
	}

	if *input == "" {
		*input = "./test/input"
	}

	log.Print("reading deck from ", *input)

	mainDeck, err := indexer.Read(*input)
	if err == nil {
		log.Print(mainDeck.Title, " read successful.")
	} else {
		log.Fatal(err)
	}

	if *output != "" {
		log.Print("writing ", mainDeck.Title, " to ", *output)
		err = indexer.Write(mainDeck, *output)
		if err == nil {
			log.Print(mainDeck.Title, " write successful.")
		} else {
			log.Fatal(err)
		}
	}
}
