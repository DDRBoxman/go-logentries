package main

import (
	"log"
	"io"
	"os"

	"github.com/DDRBoxman/go-logentries"
)

func main() {
	logentriesWriter := logentries.New("XXXX-XXXXXXX-XXXXX-XXX") //Token here

	multiWriter := io.MultiWriter(os.Stderr, logentriesWriter)

	log.SetOutput(multiWriter)

	log.Print("test")
}