package services

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type iLogger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
}

type Logger struct {
	debug  *log.Logger
	debugf *log.Logger
	info   *log.Logger
	infof  *log.Logger
}

var _ iLogger = (*Logger)(nil) // requires satisfactory implementation of the interface (unnecessary but nice to know)

func NewLogger() *Logger {
	return &Logger{
		debug:  log.New(os.Stdout, "DEBUG ", log.Ldate|log.Ltime|log.Llongfile),
		debugf: log.New(os.Stdout, "DEBUG ", log.Ldate|log.Ltime|log.Llongfile),
		info:   log.New(os.Stdout, "INFO ", log.LstdFlags),
		infof:  log.New(os.Stdout, "INFO ", log.LstdFlags),
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.debugf.Output(2, fmt.Sprint(args...))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.debugf.Output(2, fmt.Sprintf(format, args...))
}

func (l *Logger) Info(args ...interface{}) {
	l.infof.Print(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.infof.Printf(format, args...)
}

func LogRequestResponse(next http.Handler, l Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := &logResponseWriter{
			w,
			http.StatusOK,
		}

		next.ServeHTTP(lw, r)

		l.Infof("%s %s %d %s", r.Method, r.RequestURI, lw.statusCode, http.StatusText(lw.statusCode))
	})
}

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// customise WriteHeader for next.ServeHTTP http.ResponseWriter
func (lw *logResponseWriter) WriteHeader(statusCode int) {
	lw.statusCode = statusCode
	lw.ResponseWriter.WriteHeader(statusCode)
}
