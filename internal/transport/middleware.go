package transport

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type (
	ResponseData struct {
		Status int
		Size   int
	}

	LoggingResponseWriter struct {
		http.ResponseWriter
		ResponseData *ResponseData
	}
)

func (r *LoggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.ResponseData.Size += size
	return size, err
}

func (r *LoggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.ResponseData.Status = statusCode
}

func logFields(handler string) log.Fields {
	return log.Fields{
		"handler": handler,
	}
}

func logError(handler string, err error) {
	log.WithFields(logFields(handler)).Error(err)
}

// withLogging выполняет функцию middleware с логированием.
// Содержит сведения о URI, методе запроса и времени, затраченного на его выполнение.
// Сведения об ответах должны содержать код статуса и размер содержимого ответа.
func withLogging(next http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &ResponseData{
			Status: 0,
			Size:   0,
		}
		lw := LoggingResponseWriter{
			ResponseWriter: w,
			ResponseData:   responseData,
		}

		next.ServeHTTP(&lw, r)

		duration := time.Since(start)

		log.WithFields(log.Fields{
			"uri":      r.RequestURI,
			"method":   r.Method,
			"duration": duration,
			"status":   responseData.Status,
			"size":     responseData.Size,
		}).Info("request details: ")
	}
	return http.HandlerFunc(logFn)
}
