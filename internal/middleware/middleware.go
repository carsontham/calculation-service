package middleware

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// simple logging that outputs to app.logs
// to use this logging -> simply use log.Println()
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the current date
		currentTime := time.Now()

		// Format the date as a string
		dateString := currentTime.Format("20060102")

		// Use the date string in the log file name
		logFileName := "./logs/server_" + dateString + ".logs"
		logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()

		// print in terminal and write to logs
		multi := io.MultiWriter(logFile, os.Stdout)
		log.SetOutput(multi)

		// write only in logs, does not appear terminal
		// log.SetOutput(logFile)
		log.SetPrefix("TRACE: ")
		log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
