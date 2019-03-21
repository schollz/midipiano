package main

import (
	"os"

	"github.com/schollz/midi"
)

func main() {

	f, err := os.Create("midi.mid")
	if err != nil {
		panic(err)
	}
	defer func() {
		f.Close()
	}()
	e := midi.NewEncoder(f, midi.SingleTrack, 960)
	tr := e.NewTrack()

	vel := 90
	//C3 to B3
	for i := 60; i < 72; i++ {
		tr.Add(0, midi.NoteOn(0, i, vel))
		tr.Add(0.5, midi.NoteOff(0, i)) // duration
		tr.Add(0.5, midi.NoteOff(0, 0)) // rest until next note
	}
	tr.Add(1, midi.EndOfTrack())

	if err := e.Write(); err != nil {
		panic(err)
	}

}
