package db

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
