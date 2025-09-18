package main

import (
	"encoding/json"
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
	Status    string
	Timestamp time.Time
}

func checkHealth(serviceName string) HealthStatus {
	return HealthStatus{
		Service:   serviceName,
		Status:    StatusName[Status(rand.Intn(2))],
		Timestamp: time.Now(),
	}
}

func main() {
	file, err := os.OpenFile("health.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	fileLogger = log.New(file, "HEALTH: ", log.LstdFlags)
	services := []string{"Database", "Cache", "API", "Storage"}

	results := []HealthStatus{}
	for _, v := range services {
		results = append(results, checkHealth(v))
	}

	for _, v := range results {
		log.Printf("Checked %s: %s\n", v.Service, v.Status)
		b, err := json.Marshal(v)
		if err != nil {
			log.Printf("Error while trying to marshal: %s\n", err)
		} else {
			fileLogger.Println(string(b))
		}
	}

	b, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatalf("Error while trying to marshal results%s\n", err)
	}
	log.Println(string(b))
}
