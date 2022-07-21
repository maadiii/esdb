package repo

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

// nolint:gochecknoglobals
var (
	cfg  Config
	once sync.Once
)

func GetConfig() Config {
	testv := flag.Lookup("test.v")
	e := os.Getenv("REALTEST")

	if testv != nil && e == "" {
		cfg = initConfig()
	} else if testv != nil && e != "" {
		once.Do(func() {
			cfg = initConfig()
		})
	}

	return cfg
}

func initConfig() Config {
	host := os.Getenv(HostEnv)
	if host == "" {
		panic("host name can not be empty")
	}

	port := os.Getenv(PortEnv)
	if port == "" {
		panic("port can not be empty")
	}

	username := os.Getenv(UsernameEnv)
	if username == "" {
		panic("username can not be empty")
	}

	password := os.Getenv(PasswordEnv)
	if password == "" {
		panic("password can not be empty")
	}

	dbname := os.Getenv(DBNameEnv)
	if dbname == "" {
		panic("dbname can not be empty")
	}

	schema := os.Getenv(SchemaEnv)
	if schema == "" {
		panic("schema can not be empty")
	}

	sslMode := os.Getenv(SSLModeEnv)
	if sslMode == "" {
		panic("ssl mode can not be empty")
	}

	timezone := os.Getenv(TimezoneEnv)
	if timezone == "" {
		panic("timezone can not be empty")
	}

	cfg := Config{
		Driver:     DriverName,
		Host:       host,
		Port:       port,
		Username:   username,
		Password:   password,
		DBName:     dbname,
		Schema:     schema,
		TestDBName: "esdb_test",
		SslMode:    sslMode,
		TimeZone:   timezone,
	}

	return cfg
}

type Config struct {
	Driver     string
	Host       string
	Port       string
	Username   string
	Password   string
	DBName     string
	Schema     string
	TestDBName string
	SslMode    string
	TimeZone   string
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.Password, c.DBName, c.SslMode,
	)
}

func (c *Config) DSNWithSchema() string {
	dsn := fmt.Sprintf("%s search_path=%s", c.DSN(), c.Schema)

	return dsn
}
