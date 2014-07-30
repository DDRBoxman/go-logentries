package main

import (
	"io"
	"log"
	"os"

	"github.com/DDRBoxman/go-logentries"
)

func main() {
	logentriesWriter := logentries.New("XXXX-XXXXXXX-XXXXX-XXX") //Token here
	defer logentriesWriter.Close()

	multiWriter := io.MultiWriter(os.Stderr, logentriesWriter)

	log.SetOutput(multiWriter)

	log.Print("test")
	log.Print("test2")
	log.Print("test3")
}
