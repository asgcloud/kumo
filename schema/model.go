package schema

// Server is a strcuture for a virtual machine information
type Server struct {
	ServerID   string  `json:"server_id"`
	ProjectID  string  `json:"project_id"`
	ServerName string  `json:"server_name"`
	CPU        int64   `json:"cpu"`
	RAM        int64   `json:"ram"`
	Storage    float64 `json:"storage"`
	Status     string  `json:"status"`
	State      string  `json:"state"`
	Tenancy    string  `json:"tenancy"`
	Host       string  `json:"host"`
}
