package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	/*notes = map[string]float32{*/
	//"C0":  130.81,
	//"C0X": 138.59,
	//"D0":  146.83,
	//"D0X": 155.56,
	//"E0":  164.81,
	//"F0":  174.61,
	//"F0X": 185.00,
	//"G0":  196.00,
	//"G0X": 207.65,
	//"A0":  220.00,
	//"A0X": 233.08,
	//"B0":  246.94,
	//"C1":  261.63,
	//"C1X": 277.18,
	//"D1":  293.66,
	//"D1X": 311.13,
	//"E1":  329.63,
	//"F1":  349.23,
	//"F1X": 369.99,
	//"G1":  391.00,
	//"G1X": 415.30,
	//"A1":  440.00,
	//"A1X": 466.16,
	//"B1":  493.88,
	//"C2":  523.25,
	//"C2X": 554.37,
	//"D2":  587.33,
	//"D2X": 622.25,
	//"E2":  659.26,
	//"F2":  698.46,
	//"F2X": 739.99,
	//"G2":  783.99,
	//"G2X": 830.61,
	//"A2":  880.00,
	//"A2X": 923.33,
	//"B2":  987.77,
	//"C3":  1046.50,
	/*}*/
	notes = map[string]int{
		"B0":  31,
		"C1":  33,
		"CS1": 35,
		"D1":  37,
		"DS1": 39,
		"E1":  41,
		"F1":  44,
		"FS1": 46,
		"G1":  49,
		"GS1": 52,
		"A1":  55,
		"AS1": 58,
		"B1":  62,
		"C2":  65,
		"CS2": 69,
		"D2":  73,
		"DS2": 78,
		"E2":  82,
		"F2":  87,
		"FS2": 93,
		"G2":  98,
		"GS2": 104,
		"A2":  110,
		"AS2": 117,
		"B2":  123,
		"C3":  131,
		"CS3": 139,
		"D3":  147,
		"DS3": 156,
		"E3":  165,
		"F3":  175,
		"FS3": 185,
		"G3":  196,
		"GS3": 208,
		"A3":  220,
		"AS3": 233,
		"B3":  247,
		"C4":  262,
		"CS4": 277,
		"D4":  294,
		"DS4": 311,
		"E4":  330,
		"F4":  349,
		"FS4": 370,
		"G4":  392,
		"GS4": 415,
		"A4":  440,
		"AS4": 466,
		"B4":  494,
		"C5":  523,
		"CS5": 554,
		"D5":  587,
		"DS5": 622,
		"E5":  659,
		"F5":  698,
		"FS5": 740,
		"G5":  784,
		"GS5": 831,
		"A5":  880,
		"AS5": 932,
		"B5":  988,
		"C6":  1047,
		"CS6": 1109,
		"D6":  1175,
		"DS6": 1245,
		"E6":  1319,
		"F6":  1397,
		"FS6": 1480,
		"G6":  1568,
		"GS6": 1661,
		"A6":  1760,
		"AS6": 1865,
		"B6":  1976,
		"C7":  2093,
		"CS7": 2217,
		"D7":  2349,
		"DS7": 2489,
		"E7":  2637,
		"F7":  2794,
		"FS7": 2960,
		"G7":  3136,
		"GS7": 3322,
		"A7":  3520,
		"AS7": 3729,
		"B7":  3951,
		"C8":  4186,
		"CS8": 4435,
		"D8":  4699,
		"DS8": 4978,
	}
	music_sheet []byte
)

func beep(freq int, dur int) {
	cmd := exec.Command("./beep", "-f", strconv.Itoa(freq), "-l", strconv.Itoa(dur))
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func init() {
	var filename *string
	var err error

	filename = flag.String("music_sheet", "", "the music sheet to load and play")
	flag.Parse()
	music_sheet, err = ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("beep-jam 0.1")

	s := string(music_sheet[:])
	lines := strings.Split(s, "\n")
	for _, l := range lines {
		l := strings.TrimSpace(l)
		if l == "" {
			continue
		}
		if strings.HasPrefix(l, "PAUSE") {
			parts := strings.Split(l, " ")
			delay, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			time.Sleep(time.Duration(delay) * time.Millisecond)
		} else if strings.HasPrefix(l, ";") {
			// RIEN COMMENTAIRE
		} else {
			parts := strings.Split(l, " ")
			note := parts[0]
			duration, err := strconv.Atoi(parts[1])

			if err != nil {
				panic(err)
			}

			fmt.Println("Playing", note, "for", duration, "ms")
			beep(notes[note], duration)
		}
	}
}
