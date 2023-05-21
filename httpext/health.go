package httpext

import (
	"log"
	"net/http"
	"time"
)

type HealthChecker interface {
	Start()
}

type healthChecker struct {
	port   int
	status int
}

func NewHealthChecker(port int) HealthChecker {
	return &healthChecker{
		port: port,
	}
}

func (h *healthChecker) Start() {
	http.HandleFunc("/", h.HealthCheck)
	log.Fatalln(http.ListenAndServe(Port(h.port).Address(), nil))
}

func (h *healthChecker) HealthCheck(writer http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		log.Printf("%s %s%s %s\n", r.Method, r.Host, r.RequestURI, time.Since(start).String())
	}()

	_, _ = writer.Write([]byte("OK"))
}
