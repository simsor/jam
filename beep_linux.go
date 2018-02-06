package main

import (
	"os"
	"syscall"
	"time"
)

const (
	KIOCSOUND       = 0x4B2F
	CLOCK_TICK_RATE = 1193180
)

type Beeper struct {
	fd *os.File
}

func NewBeeper() (*Beeper, error) {
	b := &Beeper{}

	fd, err := os.Create("/dev/tty0")
	if err != nil {
		return nil, err
	}

	b.fd = fd
	return b, nil
}

func (b *Beeper) syscallBeep(freq float32) {
	var f uintptr

	if freq == 0.0 {
		f = uintptr(0)
	} else {
		f = uintptr(CLOCK_TICK_RATE / freq)
	}

	syscall.Syscall(syscall.SYS_IOCTL, b.fd.Fd(), uintptr(KIOCSOUND), f)
}

func (b *Beeper) Beep(freq float32, dur int) {
	b.syscallBeep(freq)
	time.Sleep(time.Duration(dur) * time.Millisecond)
	b.syscallBeep(0.0)
}
