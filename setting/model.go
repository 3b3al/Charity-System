package setting

type Settings struct {
	Database  Database
	BatchSize int
}

type DatabaseType string

const (
	DatabaseTypeMySQL    DatabaseType = "mysql"
	DatabaseTypePostgres DatabaseType = "postgres"
)

type Database struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Schema   string
	Type     DatabaseType
}

type Queue struct {
	MQAddress    string   `json:"mqAdress"`
	ExchangeName string   `json:"exchangeName"`
	ExchangeType string   `json:"exchangeType"`
	QueueName    string   `json:"queueName"`
	Bindings     []string `json:"routeKeys"`
}
