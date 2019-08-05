package schema

// Server is a strcuture for a virtual machine information
type Server struct {
	ServerID   string  `sql:"server_id,pk" json:"server_id"`
	ProjectID  string  `sql:"project_id,notnull" pg:"fk:project_id" json:"project_id"`
	ServerName string  `sql:"server_name,notnull,unique" json:"server_name"`
	CPU        int64   `sql:"cpu" json:"cpu"`
	RAM        int64   `sql:"ram" json:"ram"`
	Storage    float64 `sql:"storage" json:"storage"`
	Status     string  `sql:"server_status" json:"server_status"`
	State      string  `sql:"server_state" json:"server_state"`
	Tenancy    string  `sql:"tenancy" json:"tenancy"`
	Host       string  `sql:"host" json:"host"`
}
