package setting

var Conf = Settings{
	Database: Database{
		Username: "postgres",
		Password: "password",
		Host:     "127.0.0.1",
		Port:     "5432",
		Database: "charity_system",
		Schema:   "charity_system",
		Type:     DatabaseTypePostgres,
	},
	BatchSize: 100,
}
