Logentries logger for Go
========================

![GOPHER](http://ddrboxman.github.io/go-logentries/logo.png)

[![MIT](http://img.shields.io/badge/license-MIT-green.svg)](LICENSE) [![GODOC](http://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/DDRBoxman/go-logentries)

Drop in solution for integrating logentries into your app.
Implements the **io.Writer** interface so it will easily forward anything logged from the **log** package will show up in logentries.

Uses channels and buffered goroutines so it's non blocking for most use.

###Example

```
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
```
