package main

type Beeper struct {
}

func NewBeeper() (*Beeper, error) {
	panic("Sorry, but beep-jam doesn't work on Windows :(\nUse the Web version!")
}

func (b *Beeper) Beep(freq float32, dur int) {
}
