package db_config

type DBConfig struct {
	user     string
	password string
	host     string
	port     int
	database string
}

func CreateNewDBConfig(user string, pass string, host string, port int, dbName string) *DBConfig {
	return &DBConfig{
		user:     user,
		password: pass,
		host:     host,
		port:     port,
		database: dbName,
	}
}

func (c *DBConfig) GetUser() string {
	return c.user
}

func (c *DBConfig) GetPassword() string {
	return c.password
}

func (c *DBConfig) GetHost() string {
	return c.host
}

func (c *DBConfig) GetPort() int {
	return c.port
}

func (c *DBConfig) GetDatabaseName() string {
	return c.database
}
