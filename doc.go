/*
Package logentries is a drop in solution for integrating logentries into your app.
Implements the **io.Writer** interface so it will easily forward anything logged from the **log** package will show up in logentries.

Uses channels and buffered goroutines so it's non blocking for most use.

Author: Colin Edwards

Example

	logentriesWriter := logentries.New("XXXX-XXXXXXX-XXXXX-XXX") //Token here
	defer logentriesWriter.Close()

	multiWriter := io.MultiWriter(os.Stderr, logentriesWriter)
	
	log.SetOutput(multiWriter)

	log.Print("test")
*/
package logentries