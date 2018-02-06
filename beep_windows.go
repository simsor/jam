package main

import (
	"syscall"
)

type Beeper struct {
	kernel32 *syscall.DLL
	procBeep *syscall.Proc
}

func NewBeeper() (*Beeper, error) {
	b := &Beeper{}

	h := syscall.MustLoadDLL("kernel32.dll")
	b.kernel32 = h

	proc := b.kernel32.MustFindProc("Beep")
	b.procBeep = proc

	return b, nil
}

func (b *Beeper) Beep(freq float32, dur int) {
	r1, _, lastErr := b.procBeep.Call(uintptr(freq), uintptr(dur))
	if r1 != 0 {
		panic(lastErr)
	}
}
