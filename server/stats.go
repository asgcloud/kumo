package server

// UptimeStats is a structure to hold uptime stats
type UptimeStats struct {
	Status              string `json:"status"`
	StartTime           string `json:"start_time"`
	RequestsReceived    int    `json:"requests_received"`
	ResponsesProvided   int    `json:"responses_provided"`
	OperationsCompleted int    `json:"operations_completed"`
}
