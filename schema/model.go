package schema

// Server is a strcuture for a virtual machine information
type Server struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	CPU     int64   `json:"cpu"`
	RAM     int64   `json:"ram"`
	Storage float64 `json:"storage"`
	Status  string  `json:"status"`
	State   string  `json:"state"`
	Tenancy string  `json:"tenancy"`
	Host    string  `json:"host"`
}
