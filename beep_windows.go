package main

import (
	"syscall"
)

type Beeper struct {
	kernel32 uintptr
	procBeep uintptr
}

func NewBeeper() (*Beeper, error) {
	b := make(Beeper)

	h := syscall.MustLoadLibrary("kernel32.dll")
	b.kernel32 = h

	h := syscall.MustGetProcAddress(b.kernel32, "Beep")
	b.procBeep = h

	return b, nil
}

func (b *Beeper) Beep(freq float32, dur int) {
	syscall.Syscall(b.procBeep, 2,
		uintptr(int(freq)),
		uintptr(dur),
		0)
}
