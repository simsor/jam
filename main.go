package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
)

var (
	music_sheet []byte
)

func init() {
	var filename *string
	var err error

	filename = flag.String("music_sheet", "", "the music sheet to load and play")
	flag.Parse()

	if *filename == "" {
		webServer, _ := NewWebServer("website/")
		fmt.Println("Running the web server on port 8080")
		fmt.Println("Run 'jam -h' if you don't want that")
		webServer.OpenBrowser()
		webServer.Run()
	}

	music_sheet, err = ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	fmt.Println("beep-jam 0.1")

	s := string(music_sheet[:])
	jammer, err := NewJammer(s)

	if err != nil {
		panic(err)
	}

	// Create the handler for SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("\n\nGot Ctrl-C'd")
		jammer.beeper.Beep(0.0, 1)
		os.Exit(1)
	}()

	// Play the song!
	numLines := len(jammer.Lines)
	for i := 0; i < numLines; i++ {
		progress := int(float64(i) / float64(numLines) * 100)

		fmt.Printf("\r[")
		for j := 0; j < progress-1; j++ {
			fmt.Printf("=")
		}
		fmt.Printf(">")

		for j := progress; j < 100; j++ {
			fmt.Printf(" ")
		}

		fmt.Printf("] %d%% ", progress)

		jammer.PlayNext()
		jammer.CurrentLine++
	}
	fmt.Print("\n")
}
