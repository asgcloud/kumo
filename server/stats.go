package server

import "time"

// UptimeStats is a structure to hold uptime stats
type UptimeStats struct {
	Status              string `json:"status"`
	StartTime           string `json:"start_time"`
	RequestsReceived    int    `json:"requests_received"`
	ResponsesProvided   int    `json:"responses_provided"`
	OperationsCompleted int    `json:"operations_completed"`
}

// NewUptimeStats returns a new set of fresh stats
func NewUptimeStats() *UptimeStats {
	stats := UptimeStats{
		Status:              "OK",
		StartTime:           time.Now().String(),
		RequestsReceived:    0,
		ResponsesProvided:   0,
		OperationsCompleted: 0,
	}
	return &stats
}

// Update increments relevant fields in the stats
func (s *UptimeStats) Update(fields ...string) {
	for _, field := range fields {
		switch field {
		case "requests":
			s.RequestsReceived++
			break
		case "responses":
			s.ResponsesProvided++
			break
		case "operations":
			s.OperationsCompleted++
			break
		default:
			break
		}
	}
}
