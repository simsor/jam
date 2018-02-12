package main

import "fmt"

type Beeper struct {
}

func NewBeeper() (*Beeper, error) {
	return nil, fmt.Errorf("Sorry, beeping the motherboard only works on Linux for now! Use the Web version :)")
}

func (b *Beeper) Beep(freq float32, dur int) {
}
