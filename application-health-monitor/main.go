package main

import "time"

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
