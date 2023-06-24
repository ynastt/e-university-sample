package configs

type postgresConfig struct {
	User     string
	Password string
	DBname   string
	Host 	 string
	Port     int
}

var PostgresConfig postgresConfig

func init() {
	PostgresConfig = postgresConfig {
		User:     "postgres",
		Password: "ub7u3nAntu",      // "postgres"
		DBname:   "Euniversity",
		Host: 	  "localhost",
		Port:     5432,
	}
}