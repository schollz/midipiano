package main

import (
	"fmt"
	"log"
	"os"

	"github.com/schollz/midi"
)

func main() {
	f, err := os.Open("midi.mid")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := midi.NewDecoder(f)
	decoder.Debug = true
	if err := decoder.Decode(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("format:", decoder.Format)
	fmt.Println(decoder.TicksPerQuarterNote, "ticks per quarter")
	fmt.Println("Debugger on:", decoder.Debug)
	for _, tr := range decoder.Tracks {
		for _, ev := range tr.Events {
			fmt.Println(ev)
		}
	}

}
