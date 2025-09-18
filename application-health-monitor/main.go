package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

var fileLogger *log.Logger

type Status int

const (
	OK Status = iota
	ERROR
	WARNING
)

var StatusName = [...]string{"OK", "ERROR", "WARNING"}

type HealthStatus struct {
	Service   string
	Status    Status
	Timestamp time.Time
}

func checkHealth(serviceName string) HealthStatus {
	return HealthStatus{
		Service:   serviceName,
		Status:    Status(rand.Intn(2)),
		Timestamp: time.Now(),
	}
}

func main() {
	file, err := os.OpenFile("health.log", os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalf("%d", err)
	}
	defer file.Close()

	fileLogger = log.New(file, "HEALTH: ", log.LstdFlags)
}
