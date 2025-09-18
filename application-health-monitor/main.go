package main

import (
	"math/rand"
	"time"
)

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
