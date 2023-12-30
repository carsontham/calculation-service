package middleware

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger() *zap.Logger {
	// Ensure the directory for the log file exists
	currentDate := time.Now().Format("20060102")

	logFilePath := "./logs/server_" + currentDate + ".logs"

	if err := os.MkdirAll(filepath.Dir(logFilePath), 0755); err != nil {
		panic(err)
	}

	// Zap configuration
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", logFilePath}
	config.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		TimeKey:      "ts",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := NewZapLogger()
		defer logger.Sync()

		// Log incoming request details
		logger.Info("Incoming request received",
			zap.String("method", r.Method),
			zap.String("uri", r.RequestURI),
			zap.String("remote_addr", r.RemoteAddr),
		)

		// Call the next handler, which can be another middleware in the chain or the final handler.
		next.ServeHTTP(w, r)
	})
}

// // // simple logging that outputs to app.logs
// // // to use this logging -> simply use log.Println()
// // func LoggingMiddleware(next http.Handler) http.Handler {
// // 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// // 		// Get the current date
// // 		currentTime := time.Now()

// // 		// Format the date as a string
// // 		dateString := currentTime.Format("20060102")

// // 		// Use the date string in the log file name
// // 		logFileName := "./logs/server_" + dateString + ".logs"
// // 		logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// // 		if err != nil {
// // 			log.Fatal(err)
// // 		}
// // 		defer logFile.Close()

// // 		// print in terminal and write to logs
// // 		multi := io.MultiWriter(logFile, os.Stdout)
// // 		log.SetOutput(multi)

// // 		// write only in logs, does not appear terminal
// // 		// log.SetOutput(logFile)
// // 		log.SetPrefix("TRACE: ")
// // 		log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
// // 		log.Println("Client IP: ", r.RemoteAddr)

// // 		log.Println(r.RequestURI, "hi")

// // 		// Call the next handler, which can be another middleware in the chain, or the final handler.
// // 		next.ServeHTTP(w, r)
// // 	})
// // }
