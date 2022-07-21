package repo_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/maadiii/esdb/internal/repo"
	"github.com/stretchr/testify/assert"
)

// nolint:paralleltest,funlen
func TestGetConfig(t *testing.T) {
	panicFunc := func() {
		_ = repo.GetConfig()
	}

	t.Run("when host name is empty", func(t *testing.T) {
		assert.PanicsWithValue(t, "host name can not be empty", panicFunc)
	})
	t.Setenv(repo.HostEnv, "127.0.0.1")

	t.Run("when port is empty", func(t *testing.T) {
		assert.PanicsWithValue(t, "port can not be empty", panicFunc)
	})
	t.Setenv(repo.PortEnv, "5432")

	t.Run("when username is empty", func(t *testing.T) {
		assert.PanicsWithValue(t, "username can not be empty", panicFunc)
	})
	t.Setenv(repo.UsernameEnv, "username")

	t.Run("when passowrd is empty", func(t *testing.T) {
		assert.PanicsWithValue(t, "password can not be empty", panicFunc)
	})
	t.Setenv(repo.PasswordEnv, "password")

	t.Run("when dbname is empty", func(t *testing.T) {
		assert.PanicsWithValue(t, "dbname can not be empty", panicFunc)
	})
	t.Setenv(repo.DBNameEnv, "esdb")

	t.Run("when schema is empty", func(t *testing.T) {
		assert.PanicsWithValue(t, "schema can not be empty", panicFunc)
	})
	t.Setenv(repo.SchemaEnv, "schema")

	t.Run("when ssl mode is empty", func(t *testing.T) {
		assert.PanicsWithValue(t, "ssl mode can not be empty", panicFunc)
	})
	t.Setenv(repo.SSLModeEnv, "sslmode")

	t.Run("when timezone is empty", func(t *testing.T) {
		assert.PanicsWithValue(t, "timezone can not be empty", panicFunc)
	})
	t.Setenv(repo.TimezoneEnv, "timezone")

	// nolint:tenv
	os.Setenv("REALTEST", "REALTEST")
	t.Run("when config init OK", func(t *testing.T) {
		t.Parallel()
		cfg := repo.Config{
			Driver:     "postgres",
			Host:       "127.0.0.1",
			Port:       "5432",
			Username:   "username",
			Password:   "password",
			DBName:     "esdb",
			Schema:     "schema",
			TestDBName: "esdb_test",
			SslMode:    "sslmode",
			TimeZone:   "timezone",
		}

		gotCfg := repo.GetConfig()

		assert.Equal(t, cfg, gotCfg)
	})

	t.Run("when get config in parallel", func(t *testing.T) {
		t.Parallel()
		cfg := repo.GetConfig()

		assert.Equal(t, "127.0.0.1", cfg.Host)
	})

	t.Run("when dsn of config is ok", func(t *testing.T) {
		cfg := repo.GetConfig()
		// nolint
		dsn := fmt.Sprint(
			"host=127.0.0.1 ",
			"port=5432 ",
			"user=username ",
			"password=password ",
			"dbname=esdb ",
			"sslmode=sslmode",
		)

		assert.Equal(t, dsn, cfg.DSN())
	})

	t.Run("when dsn with search path of config is ok", func(t *testing.T) {
		cfg := repo.GetConfig()
		// nolint
		dsn := fmt.Sprint(
			"host=127.0.0.1 ",
			"port=5432 ",
			"user=username ",
			"password=password ",
			"dbname=esdb ",
			"sslmode=sslmode ",
			"search_path=schema",
		)

		assert.Equal(t, dsn, cfg.DSNWithSchema())
	})
}
